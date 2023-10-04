package lr2

import (
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestGoto(t *testing.T) {
	// S = CC
	// C = cC | d
	g := NewGrammar()
	var (
		c = g.AddTerminal("c")
		d = g.AddTerminal("d")

		S = g.AddRule("S")
		C = g.AddRule("C")
	)

	g.SetStart(S)
	g.AddProd(S, C, C)
	g.AddProd(C, c, C)
	g.AddProd(C, d)

	t.Run("1", func(t *testing.T) {
		// S' = .S EOF, EOF
		// S = .C C, EOF
		// C = .c C, c
		// C = .c C, d
		// C = .d, c
		// C = .d, d
		var is ItemSet
		is.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
		is.Add(Item{Prod: g.GetRule(S).Prods[0], Dot: 0, Lookahead: EOF})
		is.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 0, Lookahead: c})
		is.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 0, Lookahead: d})
		is.Add(Item{Prod: g.GetRule(C).Prods[1], Dot: 0, Lookahead: c})
		is.Add(Item{Prod: g.GetRule(C).Prods[1], Dot: 0, Lookahead: d})

		gS := Goto(g, &is, S)
		testutil.RequireEqualStr(t, gS.ToString(g), `
S' = S., EOF
		`)

		gC := Goto(g, &is, C)
		testutil.RequireEqualStr(t, gC.ToString(g), `
S = C .C, EOF
C = .c C, EOF
C = .d, EOF
		`)

		gc := Goto(g, &is, c)
		testutil.RequireEqualStr(t, gc.ToString(g), `
C = .c C, c
C = .c C, d
C = c .C, c
C = c .C, d
C = .d, c
C = .d, d
		`)

		gd := Goto(g, &is, d)
		testutil.RequireEqualStr(t, gd.ToString(g), `
C = d., c
C = d., d
		`)
	})

}
