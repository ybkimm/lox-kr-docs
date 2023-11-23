package lr1

import (
	"testing"

	"github.com/dcaiafa/lox/internal/base/set"
)

func TestFirst(t *testing.T) {
	names := func(g *Grammar, syms set.Set[*Terminal]) []string {
		names := make([]string, 0, syms.Len())
		symElems := syms.Elements()
		SortTerms(symElems)
		for _, sym := range symElems {
			names = append(names, sym.TermName())
		}
		return names
	}
	assertTerminalSetEq := func(t *testing.T, g *Grammar, result set.Set[*Terminal], expected ...*Terminal) {
		t.Helper()
		expectedSet := set.New[*Terminal](expected...)
		if !result.Equal(expectedSet) {
			t.Log("Expected: ", names(g, expectedSet))
			t.Log("Actual: ", names(g, result))
			t.Fatalf("Terminal set did not match expectation")
		}
	}

	t.Run("1", func(t *testing.T) {
		// Dragon book section 4.11.
		//	E  -> TE'
		//	E' -> +TE' | ε
		//	T  -> FT'
		//	T' -> *FT' | ε
		//	F  -> ( E ) | id

		g := NewGrammar()
		var (
			tId = g.AddTerminal("id")
			tA  = g.AddTerminal("+")
			tM  = g.AddTerminal("*")
			tOP = g.AddTerminal("(")
			tCP = g.AddTerminal(")")

			rE  = g.AddRule("E")
			rEp = g.AddRule("E'")
			rT  = g.AddRule("T")
			rTp = g.AddRule("T'")
			rF  = g.AddRule("F")
		)

		g.SetStart(rE)
		g.AddProd(rE, rT, rEp)
		g.AddProd(rEp /* ε */)
		g.AddProd(rEp, tA, rT, rEp)
		g.AddProd(rT, rF, rTp)
		g.AddProd(rTp, tM, rF, rTp)
		g.AddProd(rTp /* ε */)
		g.AddProd(rF, tOP, rE, tCP)
		g.AddProd(rF, tId)

		assertTerminalSetEq(t, g, First(g, []Term{rE}), tOP, tId)
		assertTerminalSetEq(t, g, First(g, []Term{rEp}), tA, Epsilon)
		assertTerminalSetEq(t, g, First(g, []Term{rTp}), tM, Epsilon)
		assertTerminalSetEq(t, g, First(g, []Term{rEp, rE}), tA, tOP, tId)
		assertTerminalSetEq(t, g, First(g, []Term{rEp, rTp}), tA, tM, Epsilon)
		assertTerminalSetEq(t, g, First(g, []Term{rEp, tId}), tA, tId)
	})

	t.Run("2", func(t *testing.T) {
		// X -> Y Z '*'
		// Y -> '+' | ε
		// Z -> '-' | ε

		g := NewGrammar()
		var (
			tA = g.AddTerminal("+")
			tS = g.AddTerminal("-")
			tM = g.AddTerminal("*")

			rX = g.AddRule("X")
			rY = g.AddRule("Y")
			rZ = g.AddRule("Z")
		)

		g.SetStart(rX)
		g.AddProd(rX, rY, rZ, tM)
		g.AddProd(rY /* ε */)
		g.AddProd(rY, tA)
		g.AddProd(rZ, tS)
		g.AddProd(rZ /* ε */)

		assertTerminalSetEq(t, g, First(g, []Term{rX}), tA, tS, tM)
	})

	t.Run("3", func(t *testing.T) {
		// XS -> XS | X
		// X -> '+'

		g := NewGrammar()
		var (
			tA = g.AddTerminal("+")

			rXS = g.AddRule("XS")
			rX  = g.AddRule("X")
		)

		g.SetStart(rXS)
		g.AddProd(rXS, rXS, rX)
		g.AddProd(rXS, rX)
		g.AddProd(rX, tA)

		assertTerminalSetEq(t, g, First(g, []Term{rXS}), tA)
	})

	t.Run("4", func(t *testing.T) {
		// A = B C '%' E | D | '+'
		// B = '-' | ε
		// C = '/' | ε
		// D = '*' | ε
		// E = '$'

		g := NewGrammar()
		var (
			tAdd = g.AddTerminal("+")
			tSub = g.AddTerminal("-")
			tMul = g.AddTerminal("*")
			tDiv = g.AddTerminal("/")
			tRem = g.AddTerminal("%")
			tDlr = g.AddTerminal("$")

			rA = g.AddRule("A")
			rB = g.AddRule("B")
			rC = g.AddRule("C")
			rD = g.AddRule("D")
			rE = g.AddRule("E")
		)

		g.SetStart(rA)
		g.AddProd(rA, rB, rC, tRem, rE)
		g.AddProd(rA, rD)
		g.AddProd(rA, tAdd)
		g.AddProd(rB, tSub)
		g.AddProd(rB /* ε */)
		g.AddProd(rC, tDiv)
		g.AddProd(rC /* ε */)
		g.AddProd(rD, tMul)
		g.AddProd(rD /* ε */)
		g.AddProd(rE, tDlr)

		assertTerminalSetEq(t, g, First(g, []Term{rB}), tSub, Epsilon)
		assertTerminalSetEq(t, g, First(g, []Term{rB, tMul}), tSub, tMul)
		assertTerminalSetEq(t, g, First(g, []Term{rA}), tSub, tDiv, tRem, tMul, tAdd, Epsilon)
	})
}
