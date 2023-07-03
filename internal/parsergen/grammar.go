package parsergen

import (
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/lox/internal/util/logger"
	"github.com/dcaiafa/lox/internal/util/set"
)

type Grammar struct {
	Terminals []*Terminal
	Rules     []*Rule

	logger      *logger.Logger
	eof         *Terminal
	syms        map[string]Symbol
	prods       []*Prod
	sp          *Rule
	states      *stateSet
	transitions *transitions
	actions     *actionMap
	errs        Errors
}

func (g *Grammar) SetLogWriter(w io.Writer) {
	g.logger = logger.New(w)
}

func (g *Grammar) Analyze() error {
	g.preAnalysis()
	if g.failed() {
		return &g.errs
	}
	g.constructParsingTable()
	if g.failed() {
		return &g.errs
	}
	return nil
}

func (g *Grammar) preAnalysis() {
	if g.logger == nil {
		g.logger = logger.New(io.Discard)
	}

	g.syms = make(map[string]Symbol)
	g.Terminals = append(g.Terminals, epsilon)

	g.eof = &Terminal{Name: "$"}
	g.Terminals = append(g.Terminals, g.eof)

	g.sp = &Rule{
		Name: "S'",
		Prods: []*Prod{
			newProd(newTerm(g.Rules[0])),
		},
	}
	g.Rules = append(g.Rules, g.sp)

	g.createNames()
	if g.failed() {
		return
	}
	g.resolveRefs()
	if g.failed() {
		return
	}

	g.normalize()
	g.assignIndexes()
}

func (g *Grammar) fail(err error) {
	g.errs = append(g.errs, err)
}

func (g *Grammar) failed() bool {
	return len(g.errs) != 0
}

func (g *Grammar) createNames() {
	for _, terminal := range g.Terminals {
		if other := g.syms[terminal.SymName()]; other != nil {
			g.fail(&RedeclaredError{Sym: terminal, Other: other})
			continue
		}
		g.syms[terminal.SymName()] = terminal
	}

	for _, rule := range g.Rules {
		if other := g.syms[rule.SymName()]; other != nil {
			g.fail(&RedeclaredError{Sym: rule, Other: other})
			continue
		}
		g.syms[rule.SymName()] = rule
	}
}

func (g *Grammar) resolveRefs() {
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				sym := g.syms[term.Name]
				if sym == nil {
					g.fail(&UndefinedError{Term: term, Prod: prod, Rule: rule})
					continue
				}
				term.sym = sym
			}
		}
	}
}

func (g *Grammar) normalize() {
	changed := true
	for changed {
		changed = false
		for _, rule := range g.Rules {
			for _, prod := range rule.Prods {
				for i, term := range prod.Terms {
					switch term.Cardinality {
					case One:
					case ZeroOrMore:
						// a = b c*
						//  =>
						// a = b a_0
						// a_0 = c+ | e
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							newProd(newTerm(term.sym, OneOrMore)),
							newProd(),
						}
						prod.Terms[i] = newTerm(srule)
						changed = true
					case OneOrMore:
						// a = b c+
						//  =>
						// a = b a_0
						// a_0 = a_0 c
						//     | c
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							newProd(newTerm(srule), newTerm(term.sym)),
							newProd(newTerm(term.sym)),
						}
						prod.Terms[i] = newTerm(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a_0
						// a_0 = c | e
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							newProd(newTerm(term.sym)),
							newProd(),
						}
						prod.Terms[i] = newTerm(srule)
						changed = true
					default:
						panic("not reached")
					}
				}
			}
		}
	}
}

func (g *Grammar) assignIndexes() {
	for i, terminal := range g.Terminals {
		terminal.index = i
	}
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			prod.index = len(g.prods)
			prod.rule = rule
			g.prods = append(g.prods, prod)
		}
	}
}

func (g *Grammar) synthesizeRule(namePrefix string) *Rule {
	r := &Rule{
		Name: fmt.Sprintf("%s__%d", namePrefix, len(g.syms)),
	}
	g.syms[r.Name] = r
	g.Rules = append(g.Rules, r)
	return r
}

func (g *Grammar) first1(s Symbol) *set.Set[*Terminal] {
	switch s := s.(type) {
	case *Terminal:
		terminalSet := new(set.Set[*Terminal])
		terminalSet.Add(s)
		return terminalSet
	case *Rule:
		if s.firstSet != nil {
			return s.firstSet
		}
		s.firstSet = new(set.Set[*Terminal])
		for _, prod := range s.Prods {
			if len(prod.Terms) == 0 {
				s.firstSet.Add(epsilon)
			} else {
				s.firstSet.AddSet(g.first1(prod.Terms[0].sym))
			}
		}
		return s.firstSet
	default:
		panic("not-reached")
	}
}

func (g *Grammar) first(syms []Symbol) *set.Set[*Terminal] {
	var fullSet *set.Set[*Terminal]
	for i, sym := range syms {
		symSet := g.first1(sym)
		if i == 0 {
			fullSet = symSet
		} else {
			if i == 1 {
				fullSet = fullSet.Clone()
			}
			fullSet.AddSet(symSet)
		}
		if !symSet.Has(epsilon) {
			fullSet.Remove(epsilon)
			break
		}
	}
	if fullSet == nil {
		fullSet = new(set.Set[*Terminal])
	}
	return fullSet
}

func (g *Grammar) constructParsingTable() {
	initialState := newStateBuilder()
	initialState.Add(newItem(g.sp.Prods[0].index, 0, g.eof.index))
	g.closure(initialState)

	g.states = newStateSet()
	g.transitions = newTransitions()
	g.actions = newActionMap()
	g.states.Add(initialState.Build())

	for g.states.Changed() {
		g.states.ResetChanged()

		g.states.ForEach(func(fromState *state) {
			for _, sym := range g.transitionSymbols(fromState) {
				toState := g.gotoState(fromState, sym)
				g.transitions.Add(fromState, toState, sym)
			}
		})
	}

	g.logger.Logf("STATES")
	g.logger.Logf("======")
	g.logger.Logf("")

	conflict := false
	g.states.ForEach(func(s *state) {
		logger := g.logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		for _, item := range s.Items {
			prod := g.prods[item.Prod]
			if item.Dot == len(prod.Terms) {
				act := action{Type: actionReduce, Reduce: prod.rule}
				if prod.rule == g.sp {
					act = action{Type: actionAccept}
				}
				terminal := g.Terminals[item.Terminal]
				if !g.actions.Add(s, terminal, act, logger) {
					conflict = true
				}
				continue
			}
			terminal, ok := prod.Terms[item.Dot].sym.(*Terminal)
			if !ok {
				continue
			}
			shiftState := g.transitions.Get(s, terminal)
			shiftAction := action{Type: actionShift, Shift: shiftState}
			if !g.actions.Add(s, terminal, shiftAction, logger) {
				conflict = true
			}
		}
	})
	if conflict {
		g.fail(ErrConflict)
	}
}

func (g *Grammar) closure(i *stateBuilder) {
	changed := true
	for changed {
		changed = false
		// For each item [A -> α.Bβ, a]:
		for _, item := range i.items {
			prod := g.prods[item.Prod]
			if item.Dot == len(prod.Terms) {
				continue
			}
			B, ok := prod.Terms[item.Dot].sym.(*Rule)
			if !ok {
				continue
			}
			beta := termSymbols(prod.Terms[item.Dot+1:])
			a := g.Terminals[item.Terminal]
			firstSet := g.first(append(beta, a))
			for _, prodB := range B.Prods {
				firstSet.ForEach(func(t *Terminal) {
					changed = i.Add(newItem(prodB.index, 0, t.index)) || changed
				})
			}
		}
	}
}

func (g *Grammar) gotoState(i *state, x Symbol) *state {
	j := newStateBuilder()
	for _, item := range i.Items {
		prod := g.prods[item.Prod]
		if item.Dot == len(prod.Terms) {
			continue
		}
		term := prod.Terms[item.Dot].sym
		if term != x {
			continue
		}
		j.Add(newItem(item.Prod, item.Dot+1, item.Terminal))
	}
	g.closure(j)
	return g.states.Add(j.Build())
}

func (g *Grammar) transitionSymbols(s *state) []Symbol {
	symSet := new(set.Set[Symbol])
	for _, item := range s.Items {
		prod := g.prods[item.Prod]
		if item.Dot >= len(prod.Terms) {
			continue
		}
		symSet.Add(prod.Terms[item.Dot].sym)
	}
	syms := symSet.Elements()

	// Symbol order determines state creation order.
	// Make the analysis deterministic by sorting.
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})

	return syms
}

func (g *Grammar) PrintStateGraph(w io.Writer) {
	fmt.Fprintf(w, "digraph G {\n")
	g.states.ForEach(func(s *state) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			s.Index,
			fmt.Sprintf("I%d\n%v", s.Index, s.ToString(g)),
		)
	})
	g.transitions.ForEach(func(from, to *state, sym Symbol) {
		fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
			from.Index,
			to.Index,
			sym.SymName())
	})
	fmt.Fprintf(w, "}\n")
}

func (g *Grammar) Print(w io.Writer) {
	fmt.Fprintf(w, "Terminals:\n")
	for _, terminal := range g.Terminals {
		fmt.Fprintf(w, "  %v\n", terminal.Name)
	}
	fmt.Fprintf(w, "Rules:\n")
	for _, rule := range g.Rules {
		rule.Print(w)
	}
}
