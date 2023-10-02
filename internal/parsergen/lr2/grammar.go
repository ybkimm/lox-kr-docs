package lr2

import (
	"fmt"
	"io"
	"math"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/util/set"
)

const (
	EOF     = 0
	Error   = 1
	Epsilon = math.MaxInt
	SPrime  = -1
)

func IsTerminal(sym int) bool {
	return sym >= 0
}

func IsRule(sym int) bool {
	return !IsTerminal(sym)
}

type Terminal struct {
	Name     string
	UserData any
}

var epsilon = &Terminal{
	Name: "ε",
}

// Rule, also known as "non-terminal", is a named collection of productions.
// For example, the following is a rule:
//
//	expr = expr '+' expr @left(1) | '(' expr ')' | NUM
type Rule struct {
	Name     string
	Prods    []int
	UserData any
}

// Prod, or production, is a ordered set of terms belonging to a Rule.
// For example:
//
//	        +--------Prod--------+
//	        |                    |
//		expr = expr '+' expr @left(1) | '(' expr ')' |   NUM
//	         ^
//	       term
type Prod struct {
	Rule     *Rule
	Terms    []int
	UserData any
}

// Grammar represents a LR1 grammar.
type Grammar struct {
	terminals []*Terminal
	rules     []*Rule
	prods     []*Prod
}

// NewGrammar creates a new Grammar.
func NewGrammar() *Grammar {
	g := &Grammar{}
	n := g.AddTerminal("EOF")
	assert.True(n == EOF)
	n = g.AddTerminal("ERROR")
	assert.True(n == Error)
	n = g.AddRule("S'")
	assert.True(n == SPrime)
	return g
}

// SetStart sets the Start rule for a grammar. This is the actual thing we are
// trying to derive. If a Rule is not in the transitive closure of things
// derivable from the start rule, it will never be derived. SetStart must be
// called once and only once.
func (g *Grammar) SetStart(ruleIndex int) {
	assert.True(len(g.GetRule(SPrime).Prods) == 0) // start can only be called once
	assert.True(IsRule(ruleIndex))
	g.AddProd(SPrime, ruleIndex)
}

// AddTerminal adds a Terminal to the grammar, and returns its symbol index.
// GetTerminal can be used to a retrieve a `Terminal` object from a symbol
// index. IsTerminal can be used to determine whether a symbol index references
// a Terminal.
func (g *Grammar) AddTerminal(name string) int {
	t := &Terminal{
		Name: name,
	}
	g.terminals = append(g.terminals, t)
	return len(g.terminals) - 1
}

// AddRule adds a Rule to the grammar, and returns its symbol index. GetRule can
// be used to retrieve a `Rule` object from a symbol index. IsRule can be used
// to determine whether a symbol index references a Rule.
func (g *Grammar) AddRule(name string) int {
	r := &Rule{
		Name: name,
	}
	g.rules = append(g.rules, r)
	return -len(g.rules)
}

// AddProd adds a Prod to a Rule.
func (g *Grammar) AddProd(ruleIndex int, terms ...int) int {
	rule := g.GetRule(ruleIndex)

	p := &Prod{
		Rule: rule,
	}

	g.prods = append(g.prods, p)
	prodIndex := len(g.prods) - 1
	rule.Prods = append(rule.Prods, prodIndex)
	p.Terms = append(p.Terms, terms...)

	return prodIndex
}

// GetTerminal returns the `Terminal` referenced by a symbol index.
// symIndex must reference a Terminal, not a Rule.
func (g *Grammar) GetTerminal(symIndex int) *Terminal {
	assert.True(IsTerminal(symIndex))
	if symIndex == Epsilon {
		return epsilon
	}
	return g.terminals[symIndex]
}

// GetRule returns the `Rule` referenced by a symbol index.
// symIndex must reference a Rule, not a Terminal.
func (g *Grammar) GetRule(r int) *Rule {
	assert.True(IsRule(r))
	r = -r - 1
	return g.rules[r]
}

// GetSymbolName returns the name of a rule or symbol referenced by the symbol
// index.
func (g *Grammar) GetSymbolName(symIndex int) string {
	if IsTerminal(symIndex) {
		return g.GetTerminal(symIndex).Name
	} else {
		return g.GetRule(symIndex).Name
	}
}

// First returns the set of Terminals that could be derived first from a set of
// symbols composed of Rules or Terminals. The Dragon Book section 4.4 has a
// rigorous albeit hard-to-follow definition for the FIRST function, but it is
// easier to understand with examples:
//
// Given:
//
//		A = B C '%' E | D | '+'
//		B = '-' | ε
//		C = '/' | ε
//		D = '*' | ε
//	  E = '$'
//
//	First(B) = { '-', ε }
//
// This should be intuitive.
//
//	First(B, '*') = { '-', '*' }
//
// Because First(B) includes ε, the result must also include First('*') but not
// ε since First('*') does not include ε.
//
//	First(A) = { '-', '/', '%', '*', '+', ε }
//
// Notice that First(B), and First(C) are included because both include ε, but
// First(E) is not included as First('%') does not include ε. '*' is included by
// First(D), and '+' by First('+'). Finally ε is in the final result only
// because First(D) includes it.
func (g *Grammar) First(syms []int) set.Set[int] {
	visited := new(set.Set[int])
	if len(syms) == 1 {
		return g.first(visited, syms[0])
	}
	var first set.Set[int]
	for _, sym := range syms {
		partialFirst := g.first(visited, sym)
		first.AddSet(partialFirst)

		// If sym[i] includes ε, include FIRST(sym[i+1]) in FIRST(syms).
		// Otherwise, stop now.
		if !partialFirst.Has(Epsilon) {
			first.Remove(Epsilon)
			break
		}
	}
	return first
}

func (g *Grammar) first(visited *set.Set[int], s int) set.Set[int] {
	if IsTerminal(s) {
		return set.New[int](s)
	}

	// Productions can contain recursion.
	// E.g.: xs = xs x | x
	if visited.Has(s) {
		return set.Set[int]{}
	}
	visited.Add(s)

	rule := g.GetRule(s)
	first := set.Set[int]{}
	for _, prodIndex := range rule.Prods {
		prod := g.prods[prodIndex]
		if len(prod.Terms) == 0 {
			first.Add(Epsilon)
			continue
		}

		addEpsilon := true
		for _, term := range prod.Terms {
			termFirst := g.first(visited, term)
			hasEpsilon := false
			termFirst.ForEach(func(s int) {
				if s == Epsilon {
					hasEpsilon = true
					return
				}
				first.Add(s)
			})
			if !hasEpsilon {
				addEpsilon = false
				break
			}
		}
		if addEpsilon {
			first.Add(Epsilon)
		}
	}
	return first
}

// Print will write a visual representation of the grammar to an io.Writer for
// debugging purposes.
func (g *Grammar) Print(w io.Writer) {
	fmt.Fprintf(w, "Terminals:\n")
	for _, t := range g.terminals {
		fmt.Fprintf(w, " %v\n", t.Name)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Rules:\n")
	for _, r := range g.rules {
		fmt.Fprintf(w, " %v = ", r.Name)
		for i, pi := range r.Prods {
			if i != 0 {
				fmt.Fprintf(w, " | ")
			}
			p := g.prods[pi]
			for j, ti := range p.Terms {
				if j != 0 {
					fmt.Fprintf(w, " ")
				}
				fmt.Fprint(w, g.GetSymbolName(ti))
			}
			if len(p.Terms) == 0 {
				fmt.Fprintf(w, "ε")
			}
		}
		fmt.Fprintf(w, "\n")
	}
}
