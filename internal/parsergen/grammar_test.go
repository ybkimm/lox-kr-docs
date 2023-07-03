package parsergen

import (
	"testing"
)

func TestFirst(t *testing.T) {
	sg := grammar4_11()
	g, err := sg.ToAugmentedGrammar()
	if err != nil {
		t.Fatalf("ToAugmentedGrammar failed: %v", err)
	}

	syms, _ := g.symbolMap()

	assertTerminalSetEq(t, g.first1(syms["E"]), "(", "id")
	assertTerminalSetEq(t, g.first1(syms["E'"]), "+", "ε")
	assertTerminalSetEq(t, g.first1(syms["T'"]), "*", "ε")
	assertTerminalSetEq(t, g.first([]Symbol{syms["E'"], syms["E"]}), "+", "(", "id")
	assertTerminalSetEq(t, g.first([]Symbol{syms["E'"], syms["T'"]}), "+", "*", "ε")
	assertTerminalSetEq(t, g.first([]Symbol{syms["E'"], syms["id"]}), "+", "id")
}
