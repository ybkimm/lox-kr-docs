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
		is.Add(Item{Prod: sPrimeProdIndex, Dot: 0, Lookahead: eofIndex})
		is.Add(Item{Prod: S.Prods[0].Index, Dot: 0, Lookahead: eofIndex})
		is.Add(Item{Prod: C.Prods[0].Index, Dot: 0, Lookahead: c.Index})
		is.Add(Item{Prod: C.Prods[0].Index, Dot: 0, Lookahead: d.Index})
		is.Add(Item{Prod: C.Prods[1].Index, Dot: 0, Lookahead: c.Index})
		is.Add(Item{Prod: C.Prods[1].Index, Dot: 0, Lookahead: d.Index})

		result := Next(g, is)
		testutil.RequireEqual(t, TermNames(result), []string{"C", "S", "c", "d"})
	})

	t.Run("2", func(t *testing.T) {
		// C = c .C, EOF
		// C = .c C, EOF
		// C = .d, EOF
		var is ItemSet
		is.Add(Item{Prod: C.Prods[0].Index, Dot: 1, Lookahead: eofIndex})
		is.Add(Item{Prod: C.Prods[0].Index, Dot: 0, Lookahead: eofIndex})
		is.Add(Item{Prod: C.Prods[1].Index, Dot: 0, Lookahead: eofIndex})

		result := Next(g, is)
		testutil.RequireEqual(t, TermNames(result), []string{"C", "c", "d"})
	})

	t.Run("3", func(t *testing.T) {
		// C = d., EOF
		var is ItemSet
		is.Add(Item{Prod: C.Prods[1].Index, Dot: 1, Lookahead: eofIndex})

		result := Next(g, is)
		testutil.RequireEqual(t, TermNames(result), []string{})
	})
}
