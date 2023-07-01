package parsergen

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

type Def interface {
	DefName() string
}

type Grammar struct {
	Rules     []*Rule
	Terminals []*Terminal
	defs      map[string]Def
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

func (r *Rule) DefName() string {
	return r.Name
}

type Prod struct {
	Terms []*Term
}

type Term struct {
	Name      string
	Qualifier Qualifier

	def Def
}

type Terminal struct {
	Name string
}

func (t *Terminal) DefName() string {
	return t.Name
}
