package parsergen

import (
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/lox/internal/util/set"
)

func (g *Grammar) Analyze() error {
	ctx := newContext()

	g.preAnalysis(ctx)
	if err := ctx.Err(); err != nil {
		return err
	}

	return nil
}

func (g *Grammar) preAnalysis(ctx *context) {
	g.syms = make(map[string]Symbol)
	g.Terminals = append(g.Terminals, epsilon)

	g.eof = &Terminal{Name: "$"}
	g.Terminals = append(g.Terminals, g.eof)

	g.sp = &Rule{
		Name: "S'",
		Prods: []*Prod{
			makeProd(makeTerm(g.Rules[0])),
		},
	}
	g.Rules = append(g.Rules, g.sp)

	g.createNames(ctx)
	if ctx.Err() != nil {
		return
	}
	g.resolveRefs(ctx)
	if ctx.Err() != nil {
		return
	}

	g.normalize()
	g.assignIndexes()
}

func (g *Grammar) createNames(ctx *context) {
	for _, terminal := range g.Terminals {
		if other := g.syms[terminal.SymName()]; other != nil {
			ctx.Fail(&RedeclaredError{Sym: terminal, Other: other})
			continue
		}
		g.syms[terminal.SymName()] = terminal
	}

	for _, rule := range g.Rules {
		if other := g.syms[rule.SymName()]; other != nil {
			ctx.Fail(&RedeclaredError{Sym: rule, Other: other})
			continue
		}
		g.syms[rule.SymName()] = rule
	}
}

func (g *Grammar) resolveRefs(ctx *context) {
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				sym := g.syms[term.Name]
				if sym == nil {
					ctx.Fail(&UndefinedError{Term: term, Prod: prod, Rule: rule})
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
					switch term.Qualifier {
					case NoQualifier:
					case ZeroOrMore:
						// a = b c*
						//  =>
						// a = b a_0
						// a_0 = c+ | e
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(term.sym, OneOrMore)),
							makeProd(),
						}
						prod.Terms[i] = makeTerm(srule)
						changed = true
					case OneOrMore:
						// a = b c+
						//  =>
						// a = b a_0
						// a_0 = a_0 c
						//     | c
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(srule), makeTerm(term.sym)),
							makeProd(makeTerm(term.sym)),
						}
						prod.Terms[i] = makeTerm(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a_0
						// a_0 = c | e
						srule := g.synthesizeRule(rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(term.sym)),
							makeProd(),
						}
						prod.Terms[i] = makeTerm(srule)
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

func (g *Grammar) firstOne(s Symbol) *set.Set[*Terminal] {
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
				s.firstSet.AddSet(g.firstOne(prod.Terms[0].sym))
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
		symSet := g.firstOne(sym)
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

func (g *Grammar) construct() {
	initialState := newItemSetBuilder()
	initialState.Add(newItem(g.sp.Prods[0].index, 0, g.eof.index))
	g.closure(initialState)

	g.states = newStateSet()
	g.transitions = newTransitions()
	g.states.Add(initialState.Build())

	for g.states.Changed() {
		g.states.ResetChanged()

		g.states.ForEach(func(fromState *itemSet) {
			for _, sym := range g.transitionSymbols(fromState) {
				toState := g.gotoState(fromState, sym)
				g.transitions.Add(fromState, toState, sym)
			}
		})
	}
}

func (g *Grammar) closure(i *itemSetBuilder) {
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
			beta := syms(prod.Terms[item.Dot+1:])
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

func (g *Grammar) gotoState(i *itemSet, x Symbol) *itemSet {
	j := newItemSetBuilder()
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

func (g *Grammar) transitionSymbols(s *itemSet) []Symbol {
	symSet := new(set.Set[Symbol])
	for _, item := range s.Items {
		prod := g.prods[item.Prod]
		if item.Dot >= len(prod.Terms) {
			continue
		}
		symSet.Add(prod.Terms[item.Dot].sym)
	}
	syms := symSet.Elements()

	// The order of the symbols determine the order states are created.
	// Make the analysis deterministic by sorting the symbols.
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})

	return syms
}

func (g *Grammar) printStateGraph(w io.Writer) {
	fmt.Fprintf(w, "digraph G {\n")
	g.states.ForEach(func(state *itemSet) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			state.Index,
			fmt.Sprintf("I%d\n%v", state.Index, state.ToString(g)),
		)
	})
	g.transitions.ForEach(func(from, to *itemSet, sym Symbol) {
		fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
			from.Index,
			to.Index,
			sym.SymName())
	})
	fmt.Fprintf(w, "}\n")
}

func syms(terms []*Term) []Symbol {
	syms := make([]Symbol, len(terms))
	for i, term := range terms {
		syms[i] = term.sym
	}
	return syms
}
