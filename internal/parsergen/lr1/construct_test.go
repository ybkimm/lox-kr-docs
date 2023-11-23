package lr1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/base/baseline"
)

func runConstructTest(t *testing.T, name string, g *Grammar) {
	t.Run(name, func(t *testing.T) {
		pt := ConstructLALR(g)
		report := strings.Builder{}
		pt.Print(&report)
		fmt.Fprintln(&report, "")
		pt.PrintGraph(&report)
		baseline.Assert(t, report.String())
	})
}

func TestConstruct(t *testing.T) {
	{
		// S = C C
		// C = c C | d
		var (
			g = NewGrammar()
			c = g.AddTerminal("c")
			d = g.AddTerminal("d")
			S = g.AddRule("S")
			C = g.AddRule("C")
		)
		g.SetStart(S)
		g.AddProd(S, C, C)
		g.AddProd(C, c, C)
		g.AddProd(C, d)
		runConstructTest(t, "1", g)
	}

	{
		// S = L '=' R | R
		// L = '*' R | id
		// R = L
		var (
			g   = NewGrammar()
			eq  = g.AddTerminal("=")
			mul = g.AddTerminal("*")
			id  = g.AddTerminal("id")
			S   = g.AddRule("S")
			L   = g.AddRule("L")
			R   = g.AddRule("R")
		)
		g.SetStart(S)
		g.AddProd(S, L, eq, R)
		g.AddProd(S, R)
		g.AddProd(L, mul, R)
		g.AddProd(L, id)
		g.AddProd(R, L)
		runConstructTest(t, "2", g)
	}

	{
		// E -> E + E
		//    | E - E
		//    | E * E
		//    | num
		var (
			g   = NewGrammar()
			mul = g.AddTerminal("*")
			add = g.AddTerminal("+")
			sub = g.AddTerminal("-")
			num = g.AddTerminal("num")
			E   = g.AddRule("E")
		)
		g.SetStart(E)
		g.AddProd(E, E, add, E)
		g.LastProd().Precedence = 1
		g.AddProd(E, E, sub, E)
		g.LastProd().Precedence = 1
		g.AddProd(E, E, mul, E)
		g.LastProd().Precedence = 2
		g.AddProd(E, num)
		runConstructTest(t, "precedence", g)
	}
}
