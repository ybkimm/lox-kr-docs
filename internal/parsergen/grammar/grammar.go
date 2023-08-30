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
	GeneratedList
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
			fmt.Fprint(w, term.String())
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

type TermType int

const (
	Simple     TermType = iota
	ZeroOrMore          // *
	OneOrMore           // +
	ZeroOrOne           // ?
	List                // @list(term, sep)
)

type Term struct {
	Pos   gotoken.Position
	Type  TermType
	Name  string
	Child *Term
	Sep   *Term
}

func NewTerm(symName string) *Term {
	return &Term{
		Name: symName,
	}
}

func NewTermS(sym Symbol) *Term {
	return NewTerm(sym.SymName())
}

func (t *Term) String() string {
	switch t.Type {
	case Simple:
		return t.Name
	case ZeroOrMore:
		return t.Child.String() + "*"
	case OneOrMore:
		return t.Child.String() + "+"
	case ZeroOrOne:
		return t.Child.String() + "?"
	case List:
		return fmt.Sprintf("@list(%v, %v)", t.Child.String(), t.Sep.String())
	default:
		panic("not-reached")
	}
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
