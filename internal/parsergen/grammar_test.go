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

	assertTerminalSetEq(t, g.first(syms["E"]), "(", "id")
	assertTerminalSetEq(t, g.first(syms["E'"]), "+", "ε")
	assertTerminalSetEq(t, g.first(syms["T'"]), "*", "ε")
	assertTerminalSetEq(t, g.First([]Symbol{syms["E'"], syms["E"]}), "+", "(", "id")
	assertTerminalSetEq(t, g.First([]Symbol{syms["E'"], syms["T'"]}), "+", "*", "ε")
	assertTerminalSetEq(t, g.First([]Symbol{syms["E'"], syms["id"]}), "+", "id")
}
