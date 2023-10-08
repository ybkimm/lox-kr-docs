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
	Alias    string
	UserData any
}

var epsilon = &Terminal{
	Name: "ε",
}

// Generated is a hint for the code-generator that determines how the rule was
// produced.
type Generated int

const (
	// NotGenerated: the rule was not generated - it was specified in the grammar.
	NotGenerated Generated = iota

	// GeneratedSPrime: the rule is the S' rule generated in the process of
	// creating the augmented grammar.
	GeneratedSPrime

	// GeneratedZeroOrOne: the rule was generated to normalize a T* term.
	GeneratedZeroOrOne

	// GeneratedOneOrMore: the rule was generated to normalize a T+ term.
	GeneratedOneOrMore

	// GeneratedList: the rule was generated to normalize a @list(T, S) term.
	GeneratedList
)

func (g Generated) String() string {
	switch g {
	case NotGenerated:
		return "NotGenerated"
	case GeneratedSPrime:
		return "SPrime"
	case GeneratedZeroOrOne:
		return "ZeroOrOne"
	case GeneratedOneOrMore:
		return "OneOrMore"
	case GeneratedList:
		return "List"
	default:
		return "???"
	}
}

// Rule, also known as "non-terminal", is a named collection of productions.
// For example, the following is a rule:
//
//	expr = expr '+' expr @left(1) | '(' expr ')' | NUM
type Rule struct {
	Name      string
	Prods     []int
	Generated Generated
	UserData  any
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
	Terminals []*Terminal
	Rules     []*Rule
	Prods     []*Prod
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
	g.GetRule(n).Generated = GeneratedSPrime

	n = g.AddProd(SPrime)
	assert.True(n == SPrimeProd)

	return g
}

// SetStart sets the Start rule for a grammar. This is the actual thing we are
// trying to derive. If a Rule is not in the transitive closure of things
// derivable from the start rule, it will never be derived.
func (g *Grammar) SetStart(ruleID int) {
	assert.True(IsRule(ruleID))
	g.Prods[SPrimeProd].Terms = []int{ruleID}
}

// AddTerminal adds a Terminal to the grammar, and returns its symbol id.
// GetTerminal can be used to a retrieve a `Terminal` object from a symbol id.
// IsTerminal can be used to determine whether a symbol id references a
// Terminal.
func (g *Grammar) AddTerminal(name string) int {
	t := &Terminal{
		Name: name,
	}
	g.Terminals = append(g.Terminals, t)
	return len(g.Terminals) - 1
}

// AddRule adds a Rule to the grammar, and returns its symbol id. GetRule can be
// used to retrieve a `Rule` object from a symbol id. IsRule can be used to
// determine whether a symbol id references a Rule.
func (g *Grammar) AddRule(name string) int {
	r := &Rule{
		Name: name,
	}
	g.Rules = append(g.Rules, r)
	return -len(g.Rules)
}

// AddProd adds a Prod to a Rule.
func (g *Grammar) AddProd(ruleID int, terms ...int) int {
	rule := g.GetRule(ruleID)

	p := &Prod{
		Rule: ruleID,
	}

	g.Prods = append(g.Prods, p)
	prodIndex := len(g.Prods) - 1
	rule.Prods = append(rule.Prods, prodIndex)
	p.Terms = append(p.Terms, terms...)

	return prodIndex
}

func (g *Grammar) LastProd() *Prod {
	return g.Prods[len(g.Prods)-1]
}

// GetTerminal returns the `Terminal` referenced by a symbol id.
// symID must reference a Terminal, not a Rule.
func (g *Grammar) GetTerminal(symID int) *Terminal {
	assert.True(IsTerminal(symID))
	if symID == Epsilon {
		return epsilon
	}
	return g.Terminals[symID]
}

// GetRule returns the `Rule` referenced by a symbol id.
// symID must reference a Rule, not a Terminal.
func (g *Grammar) GetRule(symID int) *Rule {
	assert.True(IsRule(symID))
	symID = -symID - 1
	return g.Rules[symID]
}

// GetSymbolName returns the name of a rule or symbol referenced
// by the symbol id.
func (g *Grammar) GetSymbolName(symID int) string {
	if IsTerminal(symID) {
		return g.GetTerminal(symID).Name
	} else {
		return g.GetRule(symID).Name
	}
}

func (g *Grammar) GetSymbolNames(symIDs []int) []string {
	names := make([]string, len(symIDs))
	for i, symID := range symIDs {
		names[i] = g.GetSymbolName(symID)
	}
	return names
}

func (g *Grammar) GetProd(prodIndex int) *Prod {
	return g.Prods[prodIndex]
}

// Print will write a visual representation of the grammar to an io.Writer for
// debugging purposes.
func (g *Grammar) Print(w io.Writer) {
	l := logger.New(w)
	l.Logf("Terminals")
	l.Logf("=========")
	for _, t := range g.Terminals {
		l.Logf("%v", t.Name)
	}
	l.Logf("")
	l.Logf("Rules")
	l.Logf("=====")

	writeProd := func(buf *strings.Builder, pi int) {
		p := g.Prods[pi]
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
	for _, r := range g.Rules {
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
