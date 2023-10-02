package lr2

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/assert"
)

const (
	EOFIndex   = 0
	ErrorIndex = 1
)

func IsTerminal(n int) bool {
	return n >= 0
}

func IsRule(n int) bool {
	return !IsTerminal(n)
}

type Terminal struct {
	Name     string
	UserData any
}

type Rule struct {
	Name     string
	Prods    []int
	UserData any
}

type Prod struct {
	Terms    []int
	UserData any
}

type Grammar struct {
	terminals []*Terminal
	rules     []*Rule
	prods     []*Prod
}

func NewGrammar() *Grammar {
	g := &Grammar{}
	n := g.AddTerminal("EOF", nil)
	assert.True(n == EOFIndex)
	n = g.AddTerminal("ERROR", nil)
	assert.True(n == ErrorIndex)
	return g
}

func (g *Grammar) AddTerminal(name string, userData any) int {
	t := &Terminal{
		Name:     name,
		UserData: userData,
	}
	g.terminals = append(g.terminals, t)
	return len(g.terminals) - 1
}

func (g *Grammar) AddRule(name string, userData any) int {
	r := &Rule{
		Name:     name,
		UserData: userData,
	}
	g.rules = append(g.rules, r)
	return -len(g.rules)
}

func (g *Grammar) AddProd(userData any) int {
	p := &Prod{
		UserData: userData,
	}

	g.prods = append(g.prods, p)
	prodIndex := len(g.prods) - 1

	r := g.rules[len(g.rules)-1]
	r.Prods = append(r.Prods, prodIndex)

	return prodIndex
}

func (g *Grammar) AddTerm(t int) {
	p := g.prods[len(g.prods)-1]
	p.Terms = append(p.Terms, t)
}

func (g *Grammar) GetTerminal(t int) *Terminal {
	assert.True(IsTerminal(t))
	return g.terminals[t]
}

func (g *Grammar) GetRule(r int) *Rule {
	assert.True(IsRule(r))
	r = -r - 1
	return g.rules[r]
}

func (g *Grammar) GetTermName(t int) string {
	if IsTerminal(t) {
		return g.GetTerminal(t).Name
	} else {
		return g.GetRule(t).Name
	}
}

func (g *Grammar) Print(w io.Writer) {
	fmt.Fprintf(w, "Terminals:\n")
	for _, t := range g.terminals {
		fmt.Fprintf(w, " %v\n", t.Name)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Rules:\n")
	for _, r := range g.rules {
		for _, pi := range r.Prods {
			fmt.Fprintf(w, " %v = ", r.Name)
			p := g.prods[pi]
			for i, ti := range p.Terms {
				if i != 0 {
					fmt.Fprintf(w, " | ")
				}
				fmt.Fprint(w, g.GetTermName(ti))
			}
			fmt.Fprintf(w, "\n")
		}
	}
}
