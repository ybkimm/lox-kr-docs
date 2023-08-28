package grammar

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/util/set"
)

var epsilon = &Terminal{Name: "ε"}

type AugmentedGrammar struct {
	Grammar
	Prods  []*Prod
	EOF    *Terminal
	Error  *Terminal
	Sprime *Rule

	nameToSymbol    map[string]Symbol
	termToSymbol    map[*Term]Symbol
	terminalToIndex map[*Terminal]int
	prodToIndex     map[*Prod]int
	prodToRule      map[*Prod]*Rule
	ruleToIndex     map[*Rule]int
	firstSets       map[*Rule]*set.Set[*Terminal]
}

func (g *Grammar) ToAugmentedGrammar(errs *errlogger.ErrLogger) *AugmentedGrammar {
	ag := &AugmentedGrammar{
		nameToSymbol:    make(map[string]Symbol),
		termToSymbol:    make(map[*Term]Symbol),
		terminalToIndex: make(map[*Terminal]int),
		prodToIndex:     make(map[*Prod]int),
		prodToRule:      make(map[*Prod]*Rule),
		ruleToIndex:     make(map[*Rule]int),
		firstSets:       make(map[*Rule]*set.Set[*Terminal]),
	}

	ag.EOF = &Terminal{Name: "EOF"}
	ag.Error = &Terminal{Name: "ERROR"}
	ag.Terminals = append(
		[]*Terminal{ag.EOF, ag.Error},
		g.Terminals...)

	ag.Sprime = &Rule{
		Name:      "S'",
		Prods:     []*Prod{NewProd(NewTermS(g.Rules[0]))},
		Generated: GeneratedSPrime,
	}
	ag.Rules = append([]*Rule{ag.Sprime}, g.Rules...)

	// Resolve references before calling normalize() to detect reference errors
	// before altering the grammar.
	ag.resolveReferences(errs)
	if errs.HasError() {
		return nil
	}

	ag.normalize()

	// We have to resolve references again because normalize() might have changed
	// the grammar. This is guaranteed to succeed, though.
	ag.resolveReferences(errs)
	if errs.HasError() {
		panic("unreachable")
	}

	ag.assignIndex()

	return ag
}

func (g *AugmentedGrammar) GetSymbol(name string) Symbol {
	return g.nameToSymbol[name]
}

func (g *AugmentedGrammar) ProdRule(prod *Prod) *Rule {
	rule := g.prodToRule[prod]
	if rule == nil {
		panic("invalid prod")
	}
	return rule
}

func (g *AugmentedGrammar) resolveReferences(errs *errlogger.ErrLogger) {
	g.nameToSymbol = make(map[string]Symbol)
	g.termToSymbol = make(map[*Term]Symbol)

	for _, terminal := range g.Terminals {
		if other := g.nameToSymbol[terminal.SymName()]; other != nil {
			errs.Errorf(terminal.Pos, "%q redeclared", terminal.Name)
			errs.Infof(other.Position(), "other %q declared here", terminal.Name)
			continue
		}
		g.nameToSymbol[terminal.SymName()] = terminal
	}
	for _, rule := range g.Rules {
		if other := g.nameToSymbol[rule.SymName()]; other != nil {
			errs.Errorf(rule.Pos, "%q redeclared", rule.Name)
			errs.Infof(other.Position(), "other %q declared here", rule.Name)
			continue
		}
		g.nameToSymbol[rule.SymName()] = rule
	}

	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				g.resolveTerm(term, errs)
			}
		}
	}
}

func (g *AugmentedGrammar) resolveTerm(term *Term, errs *errlogger.ErrLogger) {
	sym := g.nameToSymbol[term.Name]
	if sym == nil {
		errs.Errorf(term.Pos, "%q undefined", term.Name)
		return
	}
	g.termToSymbol[term] = sym
	if term.Separator != nil {
		g.resolveTerm(term.Separator, errs)
	}
}

func (g *AugmentedGrammar) assignIndex() {
	g.terminalToIndex = make(map[*Terminal]int, len(g.Terminals))
	for i, terminal := range g.Terminals {
		g.terminalToIndex[terminal] = i
	}

	g.Prods = nil
	g.ruleToIndex = make(map[*Rule]int)
	g.prodToIndex = make(map[*Prod]int, len(g.Prods))
	g.prodToRule = make(map[*Prod]*Rule, len(g.Prods))
	for ruleIndex, rule := range g.Rules {
		g.ruleToIndex[rule] = ruleIndex
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

func (g *AugmentedGrammar) RuleIndex(rule *Rule) int {
	index, ok := g.ruleToIndex[rule]
	if !ok {
		panic("invalid Rule")
	}
	return index
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
	newRule := func(namePrefix string, generated Generated) *Rule {
		r := &Rule{
			Name:      fmt.Sprintf("%s$%d", namePrefix, len(g.Rules)),
			Generated: generated,
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
						srule := newRule(rule.Name, GeneratedZeroOrOne)
						srule.Prods = []*Prod{
							NewProd(NewTerm(term.Name, OneOrMore)),
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
						srule := newRule(rule.Name, GeneratedOneOrMore)
						srule.Prods = []*Prod{
							NewProd(NewTerm(srule.Name), NewTerm(term.Name)),
							NewProd(NewTerm(term.Name)),
						}
						prod.Terms[i] = NewTermS(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a'
						// a' = c | e
						srule := newRule(rule.Name, GeneratedZeroOrOne)
						srule.Prods = []*Prod{
							NewProd(NewTerm(term.Name)),
							NewProd(),
						}
						prod.Terms[i] = NewTerm(srule.Name)
						changed = true
					default:
						panic("not reached")
					}
				}
			}
		}
	}
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
