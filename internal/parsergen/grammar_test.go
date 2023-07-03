package parsergen

import (
	"testing"
)

func TestFirst(t *testing.T) {
	g := grammar4_11()
	g.preAnalysis()
	if g.failed() {
		t.Fatalf("preAnalysis failed: %v", g.errs)
	}

	assertTerminalSetEq(t, g.first1(g.syms["E"]), "(", "id")
	assertTerminalSetEq(t, g.first1(g.syms["E'"]), "+", "ε")
	assertTerminalSetEq(t, g.first1(g.syms["T'"]), "*", "ε")
	assertTerminalSetEq(t, g.first([]Symbol{g.syms["E'"], g.syms["E"]}), "+", "(", "id")
	assertTerminalSetEq(t, g.first([]Symbol{g.syms["E'"], g.syms["T'"]}), "+", "*", "ε")
	assertTerminalSetEq(t, g.first([]Symbol{g.syms["E'"], g.syms["id"]}), "+", "id")
}
