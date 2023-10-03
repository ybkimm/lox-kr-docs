package lr2

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
		i.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
		result := Closure(g, i)
		testutil.RequireEqualStr(t, result.ToString(g), `
S' = .S EOF, EOF
S = .C C, EOF
C = .c C, c
C = .c C, d
C = .d, c
C = .d, d
`)
	})

	t.Run("2", func(t *testing.T) {
		var i ItemSet
		i.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 1, Lookahead: c})
		i.Add(Item{Prod: g.GetRule(C).Prods[0], Dot: 1, Lookahead: d})
		result := Closure(g, i)
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
