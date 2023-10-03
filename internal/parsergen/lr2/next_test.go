package lr2

import (
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestNext(t *testing.T) {
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

		result := Next(g, is)
		testutil.RequireEqual(t, g.GetSymbolNames(result), []string{"C", "S", "c", "d"})
	})

	t.Run("2", func(t *testing.T) {
		// C = c .C, EOF
		// C = .c C, EOF
		// C = .d, EOF
		var is ItemSet
		is.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 1, Lookahead: EOF})
		is.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 0, Lookahead: EOF})
		is.Add(Item{Prod: g.GetRule(C).Prods[1], Dot: 0, Lookahead: EOF})

		result := Next(g, is)
		testutil.RequireEqual(t, g.GetSymbolNames(result), []string{"C", "c", "d"})
	})

	t.Run("3", func(t *testing.T) {
		// C = d., EOF
		var is ItemSet
		is.Add(Item{Prod: g.GetRule(C).Prods[1], Dot: 1, Lookahead: EOF})

		result := Next(g, is)
		testutil.RequireEqual(t, g.GetSymbolNames(result), []string{})
	})
}
