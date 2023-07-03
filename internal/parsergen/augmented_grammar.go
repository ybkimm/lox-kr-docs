package parsergen

import (
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/lox/internal/util/set"
)

type AugmentedGrammar struct {
	Grammar
	EOF    *Terminal
	Prods  []*Prod
	Sprime *Rule
}

func (g *Grammar) ToAugmentedGrammar() (*AugmentedGrammar, error) {
	ag := new(AugmentedGrammar)
	ag.EOF = &Terminal{Name: "$"}
	ag.Terminals = append([]*Terminal{ag.EOF}, g.Terminals...)
	ag.Sprime = &Rule{
		Name:  "S'",
		Prods: []*Prod{newProd(newTerm(g.Rules[0]))},
	}
	ag.Rules = append([]*Rule{ag.Sprime}, g.Rules...)
	err := ag.resolveReferences()
	if err != nil {
		return nil, err
	}
	ag.assignIndex()
	return ag, nil
}

func (g *AugmentedGrammar) symbolMap() (map[string]Symbol, error) {
	var errs Errors
	syms := make(map[string]Symbol, len(g.Terminals)+len(g.Rules))
	for _, terminal := range g.Terminals {
		if other := syms[terminal.SymName()]; other != nil {
			errs.Add(&RedeclaredError{Sym: terminal, Other: other})
			continue
		}
		syms[terminal.SymName()] = terminal
	}

	for _, rule := range g.Rules {
		if other := syms[rule.SymName()]; other != nil {
			errs.Add(&RedeclaredError{Sym: rule, Other: other})
			continue
		}
		syms[rule.SymName()] = rule
	}
	return syms, errs.ToError()
}

func (g *AugmentedGrammar) resolveReferences() error {
	syms, err := g.symbolMap()
	if err != nil {
		return err
	}

	var errs Errors
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				sym := syms[term.Name]
				if sym == nil {
					errs.Add(&UndefinedError{Term: term, Prod: prod, Rule: rule})
					continue
				}
				term.sym = sym
			}
		}
	}
	return errs.ToError()
}

func (g *AugmentedGrammar) normalize() {
	newRule := func(namePrefix string) *Rule {
		r := &Rule{
			Name: fmt.Sprintf("%s__%d", namePrefix, len(g.Rules)),
		}
		g.Rules = append(g.Rules, r)
		return r
	}

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
						// a = b a'
						// a' = c+ | e
						srule := newRule(rule.Name)
						srule.Prods = []*Prod{
							newProd(newTerm(term.sym, OneOrMore)),
							newProd(),
						}
						prod.Terms[i] = newTerm(srule)
						changed = true
					case OneOrMore:
						// a = b c+
						//  =>
						// a = b a'
						// a' = a' c
						//    | c
						srule := newRule(rule.Name)
						srule.Prods = []*Prod{
							newProd(newTerm(srule), newTerm(term.sym)),
							newProd(newTerm(term.sym)),
						}
						prod.Terms[i] = newTerm(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a'
						// a' = c | e
						srule := newRule(rule.Name)
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

func (g *AugmentedGrammar) assignIndex() {
	for i, terminal := range g.Terminals {
		terminal.index = i
	}
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			prod.index = len(g.Prods)
			prod.rule = rule
			g.Prods = append(g.Prods, prod)
		}
	}
}

func (g *AugmentedGrammar) first1(s Symbol) *set.Set[*Terminal] {
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

func (g *AugmentedGrammar) first(syms []Symbol) *set.Set[*Terminal] {
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

func (g *AugmentedGrammar) constructLALR() {
	//g.constructLR0Kernels()
}

/*
func (g *Grammar) constructLR0Kernels() {
	initialState := newStateBuilder()
	initialState.Add(newItem(g.sp.Prods[0].index, 0, 0))
	g.states = newStateSet()
	g.states.Add(initialState.Build())

	for g.states.Changed() {
		g.states.ResetChanged()

		g.states.ForEach(func(s *state) {
			for _, item := range s.Items {
				prod := g.prods[item.Prod]
				if item.Dot == len(prod.Terms) {
					continue
				}
				term := prod.Terms[item.Dot]


			}
		})
	}
}
*/

func (g *AugmentedGrammar) closure(i *stateBuilder) {
	changed := true
	for changed {
		changed = false
		// For each item [A -> α.Bβ, a]:
		for _, item := range i.items {
			prod := g.Prods[item.Prod]
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

func (g *AugmentedGrammar) transitionSymbols(s *state) []Symbol {
	symSet := new(set.Set[Symbol])
	for _, item := range s.Items {
		prod := g.Prods[item.Prod]
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

func (g *AugmentedGrammar) Print(w io.Writer) {
	fmt.Fprintf(w, "Terminals:\n")
	for _, terminal := range g.Terminals {
		fmt.Fprintf(w, "  %v\n", terminal.Name)
	}
	fmt.Fprintf(w, "Rules:\n")
	for _, rule := range g.Rules {
		rule.Print(w)
	}
}
