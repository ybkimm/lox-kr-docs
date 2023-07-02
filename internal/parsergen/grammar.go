package parsergen

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/util/set"
)

var epsilon = &Terminal{Name: "ε", index: -1}

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Symbol interface {
	SymName() string
}

type Grammar struct {
	Terminals   []*Terminal
	Rules       []*Rule
	eof         *Terminal
	syms        map[string]Symbol
	prods       []*Prod
	sp          *Rule
	states      *stateSet
	transitions *transitions
	errs        Errors
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

type Rule struct {
	Name  string
	Prods []*Prod

	firstSet *set.Set[*Terminal]
}

func (r *Rule) Print(w io.Writer) {
	fmt.Fprintf(w, "%v = ", r.Name)
	for i, prod := range r.Prods {
		if i != 0 {
			fmt.Fprintf(w, "\n    | ")
		}
		for j, term := range prod.Terms {
			if j != 0 {
				fmt.Fprintf(w, " ")
			}
			qualifier := ""
			switch term.Qualifier {
			case NoQualifier:
				qualifier = ""
			case ZeroOrMore:
				qualifier = "*"
			case OneOrMore:
				qualifier = "+"
			case ZeroOrOne:
				qualifier = "?"
			default:
				panic("not reached")
			}
			fmt.Fprintf(w, "%s%s", term.Name, qualifier)
		}
	}
	fmt.Fprintf(w, " .\n")
}

func (r *Rule) SymName() string {
	return r.Name
}

type Prod struct {
	Terms []*Term

	rule  *Rule
	index int
}

type Term struct {
	Name      string
	Qualifier Qualifier

	sym Symbol
}

type Terminal struct {
	Name string

	index int
}

func (t *Terminal) SymName() string {
	return t.Name
}
