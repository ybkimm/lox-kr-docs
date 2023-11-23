package lr1

import (
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestClosure(t *testing.T) {
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
		var i ItemSet
		i.Add(Item{Prod: sPrimeProdIndex, Dot: 0, Lookahead: eofIndex})
		result := Closure(g, &i)
		testutil.RequireEqualStr(t, result.ToString(g), `
S' = .S, EOF
S = .C C, EOF
C = .c C, c
C = .c C, d
C = .d, c
C = .d, d
`)
	})

	t.Run("2", func(t *testing.T) {
		var i ItemSet
		i.Add(Item{Prod: C.Prods[0].Index, Dot: 1, Lookahead: c.Index})
		i.Add(Item{Prod: C.Prods[0].Index, Dot: 1, Lookahead: d.Index})
		result := Closure(g, &i)
		testutil.RequireEqualStr(t, result.ToString(g), `
C = .c C, c
C = .c C, d
C = c .C, c
C = c .C, d
C = .d, c
C = .d, d
`)
	})
}
