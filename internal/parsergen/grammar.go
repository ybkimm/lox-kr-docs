package parsergen

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/util/set"
)

type Grammar struct {
	Terminals []*Terminal
	Rules     []*Rule
}

type Symbol interface {
	SymName() string
}

type Rule struct {
	Name  string
	Prods []*Prod

	firstSet *set.Set[*Terminal]
}

func (r *Rule) SymName() string {
	return r.Name
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
			cardinality := ""
			switch term.Cardinality {
			case One:
				cardinality = ""
			case ZeroOrMore:
				cardinality = "*"
			case OneOrMore:
				cardinality = "+"
			case ZeroOrOne:
				cardinality = "?"
			default:
				panic("not reached")
			}
			fmt.Fprintf(w, "%s%s", term.Name, cardinality)
		}
	}
	fmt.Fprintf(w, " .\n")
}

type Prod struct {
	Terms []*Term

	rule  *Rule
	index int
}

func newProd(terms ...*Term) *Prod {
	return &Prod{
		Terms: terms,
	}
}

type Cardinality int

const (
	One        Cardinality = iota
	ZeroOrMore             // *
	OneOrMore              // +
	ZeroOrOne              // ?
)

type Term struct {
	Name        string
	Cardinality Cardinality

	sym Symbol
}

func newTerm(sym Symbol, q ...Cardinality) *Term {
	t := &Term{
		Name: sym.SymName(),
		sym:  sym,
	}
	if len(q) != 0 {
		t.Cardinality = q[0]
	}
	return t
}

func termSymbols(terms []*Term) []Symbol {
	syms := make([]Symbol, len(terms))
	for i, term := range terms {
		syms[i] = term.sym
	}
	return syms
}

var epsilon = &Terminal{Name: "ε", index: -1}

type Terminal struct {
	Name  string
	index int
}

func (t *Terminal) SymName() string {
	return t.Name
}
