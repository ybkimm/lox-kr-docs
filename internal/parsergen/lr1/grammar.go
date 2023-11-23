package lr1

import (
	"cmp"
	"fmt"
	gotoken "go/token"
	"io"
	"slices"
	"strings"

	"github.com/dcaiafa/lox/internal/base/assert"
	"github.com/dcaiafa/lox/internal/base/logger"
)

const (
	eofIndex        = 0
	sPrimeProdIndex = 0
)

const SPrime = "S'"

type Term interface {
	TermName() string
}

func TermNames[T Term](ts []T) []string {
	names := make([]string, len(ts))
	for i, t := range ts {
		names[i] = t.TermName()
	}
	return names
}

func SortTerms[T Term](ts []T) {
	slices.SortFunc(ts, func(a, b T) int {
		return cmp.Compare(a.TermName(), b.TermName())
	})
}

type Associativity int

const (
	Left Associativity = iota
	Right
)

type Terminal struct {
	Index    int
	Name     string
	Alias    string
	UserData any
}

func (t *Terminal) TermName() string { return t.Name }

var Epsilon = &Terminal{
	Name: "ε",
}

// Rule, also known as "non-terminal", is a named collection of productions.
// For example, the following is a rule with three productions:
//
//	expr = expr '+' expr @left(1) | '(' expr ')' | NUM
type Rule struct {
	Index    int
	Name     string
	Prods    []*Prod
	Position gotoken.Position
	UserData any
}

func (r *Rule) TermName() string { return r.Name }

// Prod, or production, is a ordered set of terms belonging to a Rule.
// For example:
//
//	        +--------Prod--------+
//	        |                    |
//		expr = expr '+' expr @left(1) | '(' expr ')' |   NUM
//	         ^
//	       term
type Prod struct {
	Index         int
	Rule          *Rule
	Terms         []Term
	Precedence    int
	Associativity Associativity
	Position      gotoken.Position
	UserData      any
}

// Grammar represents a LR1 grammar.
type Grammar struct {
	Terminals     []*Terminal
	Rules         []*Rule
	Prods         []*Prod
	EOFTerminal   *Terminal
	ErrorTerminal *Terminal
}

// NewGrammar creates a new Grammar.
func NewGrammar() *Grammar {
	g := &Grammar{}

	g.EOFTerminal = g.AddTerminal("EOF")
	g.ErrorTerminal = g.AddTerminal("ERROR")

	sprime := g.AddRule(SPrime)
	sprimeProd := g.AddProd(sprime)
	assert.True(sprimeProd.Index == sPrimeProdIndex)

	return g
}

// SetStart sets the Start rule for a grammar. This is the actual thing we are
// trying to derive. If a Rule is not in the transitive closure of things
// derivable from the start rule, it will never be derived.
func (g *Grammar) SetStart(rule *Rule) {
	g.Prods[0].Terms = []Term{rule}
}

// AddTerminal adds a Terminal to the grammar, and returns its symbol id.
// GetTerminal can be used to a retrieve a `Terminal` object from a symbol id.
// IsTerminal can be used to determine whether a symbol id references a
// Terminal.
func (g *Grammar) AddTerminal(name string) *Terminal {
	t := &Terminal{
		Index: len(g.Terminals),
		Name:  name,
	}
	g.Terminals = append(g.Terminals, t)
	return t
}

// AddRule adds a Rule to the grammar, and returns its symbol id. GetRule can be
// used to retrieve a `Rule` object from a symbol id. IsRule can be used to
// determine whether a symbol id references a Rule.
func (g *Grammar) AddRule(name string) *Rule {
	r := &Rule{
		Index: len(g.Rules),
		Name:  name,
	}
	g.Rules = append(g.Rules, r)
	return r
}

// AddProd adds a Prod to a Rule.
func (g *Grammar) AddProd(rule *Rule, terms ...Term) *Prod {
	p := &Prod{
		Index: len(g.Prods),
		Rule:  rule,
	}

	g.Prods = append(g.Prods, p)
	rule.Prods = append(rule.Prods, p)
	p.Terms = append(p.Terms, terms...)

	return p
}

func (g *Grammar) LastProd() *Prod {
	return g.Prods[len(g.Prods)-1]
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

	writeProd := func(buf *strings.Builder, p *Prod) {
		for j, ti := range p.Terms {
			if j != 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(ti.TermName())
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
