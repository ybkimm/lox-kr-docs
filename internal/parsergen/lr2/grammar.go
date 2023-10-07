package lr2

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/util/logger"
)

const (
	EOF        = 0
	Error      = 1
	Epsilon    = math.MaxInt
	SPrime     = -1
	SPrimeProd = 0
)

type Associativity int

const (
	Left Associativity = iota
	Right
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
	Rule          int
	Terms         []int
	Precedence    int
	Associativity Associativity
	UserData      any
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
	n = g.AddProd(SPrime)
	assert.True(n == SPrimeProd)

	return g
}

// SetStart sets the Start rule for a grammar. This is the actual thing we are
// trying to derive. If a Rule is not in the transitive closure of things
// derivable from the start rule, it will never be derived.
func (g *Grammar) SetStart(ruleIndex int) {
	assert.True(IsRule(ruleIndex))
	g.prods[SPrimeProd].Terms = []int{ruleIndex}
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
		Rule: ruleIndex,
	}

	g.prods = append(g.prods, p)
	prodIndex := len(g.prods) - 1
	rule.Prods = append(rule.Prods, prodIndex)
	p.Terms = append(p.Terms, terms...)

	return prodIndex
}

func (g *Grammar) LastProd() *Prod {
	return g.prods[len(g.prods)-1]
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

func (g *Grammar) GetSymbolNames(symIndex []int) []string {
	names := make([]string, len(symIndex))
	for i, symIndex := range symIndex {
		names[i] = g.GetSymbolName(symIndex)
	}
	return names
}

func (g *Grammar) GetProd(prodIndex int) *Prod {
	return g.prods[prodIndex]
}

// Print will write a visual representation of the grammar to an io.Writer for
// debugging purposes.
func (g *Grammar) Print(w io.Writer) {
	l := logger.New(w)
	l.Logf("Terminals")
	l.Logf("=========")
	for _, t := range g.terminals {
		l.Logf("%v", t.Name)
	}
	l.Logf("")
	l.Logf("Rules")
	l.Logf("=====")

	writeProd := func(buf *strings.Builder, pi int) {
		p := g.prods[pi]
		for j, ti := range p.Terms {
			if j != 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(g.GetSymbolName(ti))
		}
		if len(p.Terms) == 0 {
			buf.WriteString("ε")
		}
		if p.Precedence > 0 {
			ass := "@left"
			if p.Associativity == Right {
				ass = "@right"
			}
			fmt.Fprintf(buf, "  %v(%v)", ass, p.Precedence)
		}
	}

	var buf strings.Builder
	for _, r := range g.rules {
		buf.Reset()
		fmt.Fprintf(&buf, "%v = ", r.Name)
		if len(r.Prods) == 0 {
			buf.WriteString("<rule has no prods>")
			l.Logf("%v", buf.String())
			continue
		}
		writeProd(&buf, r.Prods[0])
		l.Logf("%v", buf.String())
		for _, pi := range r.Prods[1:] {
			buf.Reset()
			buf.WriteString("| ")
			writeProd(&buf, pi)
			l.WithIndent().Logf("%v", buf.String())
		}
	}
}
