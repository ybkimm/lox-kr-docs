package grammar

import (
	"fmt"
	"io"
)

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Decl interface {
	DeclName() string
	Print(w io.Writer)
}

type Syntax struct {
	Decls []Decl
}

type Rule struct {
	Name  string
	Prods []*Prod
}

func (r *Rule) DeclName() string { return r.Name }

func (r *Rule) Print(w io.Writer) {
	fmt.Fprintf(w, "%v = ", r.Name)
	for i, prod := range r.Prods {
		if i != 0 {
			fmt.Fprintf(w, "\n    |")
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

type Prod struct {
	Terms []*Term
	Label *Label
}

type Term struct {
	Name      string
	Literal   string
	Qualifier Qualifier
}

type Label struct {
	Label string
}

type Token struct {
	Name string
}
