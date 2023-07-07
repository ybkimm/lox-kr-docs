package grammar

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/util/multierror"
	"github.com/dcaiafa/lox/internal/util/set"
)

var epsilon = &Terminal{Name: "ε"}

type AugmentedGrammar struct {
	Grammar
	Prods  []*Prod
	EOF    *Terminal
	Sprime *Rule

	nameToSymbol    map[string]Symbol
	termToSymbol    map[*Term]Symbol
	terminalToIndex map[*Terminal]int
	prodToIndex     map[*Prod]int
	prodToRule      map[*Prod]*Rule
	firstSets       map[*Rule]*set.Set[*Terminal]
}

func (g *Grammar) ToAugmentedGrammar() (*AugmentedGrammar, error) {
	ag := &AugmentedGrammar{
		nameToSymbol:    make(map[string]Symbol),
		termToSymbol:    make(map[*Term]Symbol),
		terminalToIndex: make(map[*Terminal]int),
		prodToIndex:     make(map[*Prod]int),
		prodToRule:      make(map[*Prod]*Rule),
		firstSets:       make(map[*Rule]*set.Set[*Terminal]),
	}

	ag.EOF = &Terminal{Name: "$"}
	ag.Terminals = append(
		[]*Terminal{ag.EOF},
		g.Terminals...)

	ag.Sprime = &Rule{
		Name:  "S'",
		Prods: []*Prod{NewProd(NewTermS(g.Rules[0]))},
	}
	ag.Rules = append([]*Rule{ag.Sprime}, g.Rules...)

	ag.normalize()
	err := ag.resolveReferences()
	if err != nil {
		return nil, err
	}
	ag.assignIndex()
	return ag, nil
}

func (g *AugmentedGrammar) GetSymbol(name string) Symbol {
	sym := g.nameToSymbol[name]
	if sym == nil {
		panic("invalid symbol")
	}
	return sym
}

func (g *AugmentedGrammar) ProdRule(prod *Prod) *Rule {
	rule := g.prodToRule[prod]
	if rule == nil {
		panic("invalid prod")
	}
	return rule
}

func (g *AugmentedGrammar) resolveReferences() error {
	var errs multierror.MultiError

	for _, terminal := range g.Terminals {
		if other := g.nameToSymbol[terminal.SymName()]; other != nil {
			errs.Add(&RedeclaredError{Sym: terminal, Other: other})
			continue
		}
		g.nameToSymbol[terminal.SymName()] = terminal
	}
	for _, rule := range g.Rules {
		if other := g.nameToSymbol[rule.SymName()]; other != nil {
			errs.Add(&RedeclaredError{Sym: rule, Other: other})
			continue
		}
		g.nameToSymbol[rule.SymName()] = rule
	}

	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				sym := g.nameToSymbol[term.Name]
				if sym == nil {
					errs.Add(&UndefinedError{Term: term, Prod: prod, Rule: rule})
					continue
				}
				g.termToSymbol[term] = sym
			}
		}
	}
	return errs.ToError()
}

func (g *AugmentedGrammar) assignIndex() {
	g.terminalToIndex = make(map[*Terminal]int, len(g.Terminals))
	for i, terminal := range g.Terminals {
		g.terminalToIndex[terminal] = i
	}

	g.Prods = nil
	g.prodToIndex = make(map[*Prod]int, len(g.Prods))
	g.prodToRule = make(map[*Prod]*Rule, len(g.Prods))
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			g.prodToIndex[prod] = len(g.Prods)
			g.prodToRule[prod] = rule
			g.Prods = append(g.Prods, prod)
		}
	}
}

func (g *AugmentedGrammar) First(syms []Symbol) *set.Set[*Terminal] {
	var fullSet *set.Set[*Terminal]
	for i, sym := range syms {
		symSet := g.first(sym)
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

func (g *AugmentedGrammar) first(s Symbol) *set.Set[*Terminal] {
	switch s := s.(type) {
	case *Terminal:
		terminalSet := new(set.Set[*Terminal])
		terminalSet.Add(s)
		return terminalSet
	case *Rule:
		rule := s
		firstSet := g.firstSets[rule]
		if firstSet != nil {
			return firstSet
		}
		firstSet = new(set.Set[*Terminal])
		g.firstSets[rule] = firstSet
		for _, prod := range s.Prods {
			if len(prod.Terms) == 0 {
				firstSet.Add(epsilon)
			} else {
				termSym := g.TermSymbol(prod.Terms[0])
				firstSet.AddSet(g.first(termSym))
			}
		}
		return firstSet
	default:
		panic("not-reached")
	}
}

func (g *AugmentedGrammar) TermSymbol(term *Term) Symbol {
	sym := g.termToSymbol[term]
	if sym == nil {
		panic("invalid Term")
	}
	return sym
}

func (g *AugmentedGrammar) TermSymbols(terms []*Term) []Symbol {
	syms := make([]Symbol, len(terms))
	for i := range terms {
		syms[i] = g.TermSymbol(terms[i])
	}
	return syms
}

func (g *AugmentedGrammar) ProdIndex(prod *Prod) int {
	index, ok := g.prodToIndex[prod]
	if !ok {
		panic("invalid Prod")
	}
	return index
}

func (g *AugmentedGrammar) TerminalIndex(terminal *Terminal) int {
	index, ok := g.terminalToIndex[terminal]
	if !ok {
		panic("invalid Terminal")
	}
	return index
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
							NewProd(NewTermS(g.TermSymbol(term), OneOrMore)),
							NewProd(),
						}
						prod.Terms[i] = NewTermS(srule)
						changed = true
					case OneOrMore:
						// a = b c+
						//  =>
						// a = b a'
						// a' = a' c
						//    | c
						srule := newRule(rule.Name)
						srule.Prods = []*Prod{
							NewProd(NewTermS(srule), NewTermS(g.TermSymbol(term))),
							NewProd(NewTermS(g.TermSymbol(term))),
						}
						prod.Terms[i] = NewTermS(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a'
						// a' = c | e
						srule := newRule(rule.Name)
						srule.Prods = []*Prod{
							NewProd(NewTermS(g.TermSymbol(term))),
							NewProd(),
						}
						prod.Terms[i] = NewTermS(srule)
						changed = true
					default:
						panic("not reached")
					}
				}
			}
		}
	}
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
