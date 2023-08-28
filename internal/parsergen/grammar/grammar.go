package grammar

import (
	"fmt"
	"io"

	gotoken "go/token"
)

type Generated int

const (
	NotGenerated Generated = iota
	GeneratedSPrime
	GeneratedZeroOrOne
	GeneratedOneOrMore
)

type Grammar struct {
	Terminals []*Terminal
	Rules     []*Rule
}

type Symbol interface {
	SymName() string
	Position() gotoken.Position
}

type Rule struct {
	Name      string
	Prods     []*Prod
	Generated Generated
	Pos       gotoken.Position
}

func (r *Rule) SymName() string {
	return r.Name
}

func (r *Rule) Position() gotoken.Position {
	return r.Pos
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

type Associativity int

const (
	Left  Associativity = 0
	Right Associativity = 1
)

type Prod struct {
	Terms         []*Term
	Precence      int
	Associativity Associativity
	Pos           gotoken.Position
}

func NewProd(terms ...*Term) *Prod {
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
	List                   // @list(term, sep)
)

type Term struct {
	Name        string
	Cardinality Cardinality
	Pos         gotoken.Position
	Separator   *Term
}

func NewTerm(symName string, q ...Cardinality) *Term {
	t := &Term{
		Name: symName,
	}
	if len(q) != 0 {
		t.Cardinality = q[0]
	}
	return t
}

func NewTermS(sym Symbol, q ...Cardinality) *Term {
	return NewTerm(sym.SymName(), q...)
}

type Terminal struct {
	Name string
	Pos  gotoken.Position
}

func (t *Terminal) SymName() string {
	return t.Name
}

func (t *Terminal) Position() gotoken.Position {
	return t.Pos
}
