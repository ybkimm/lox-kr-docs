package parsergen

import (
	"strings"
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

func TestConstruct(t *testing.T) {
	g := &Grammar{
		Terminals: []*Terminal{
			{Name: "c"},
			{Name: "d"},
		},
		Rules: []*Rule{
			{
				Name: "S",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "C"}, {Name: "C"}}},
				},
			},
			{
				Name: "C",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "c"}, {Name: "C"}}},
					{Terms: []*Term{{Name: "d"}}},
				},
			},
		},
	}

	g.preAnalysis()
	if g.failed() {
		t.Fatalf("preAnalysis failed: %v", g.errs)
	}

	g.construct()

	var graph strings.Builder
	g.printStateGraph(&graph)

	expected := `
digraph G {
  I0 [label="I0\nS -> .C C, $\nC -> .c C, c\nC -> .c C, d\nC -> .d, c\nC -> .d, d\nS' -> .S, $"];
  I1 [label="I1\nS -> C .C, $\nC -> .c C, $\nC -> .d, $"];
  I2 [label="I2\nS' -> S., $"];
  I3 [label="I3\nC -> .c C, c\nC -> .c C, d\nC -> c .C, c\nC -> c .C, d\nC -> .d, c\nC -> .d, d"];
  I4 [label="I4\nC -> d., c\nC -> d., d"];
  I5 [label="I5\nS -> C C., $"];
  I6 [label="I6\nC -> .c C, $\nC -> c .C, $\nC -> .d, $"];
  I7 [label="I7\nC -> d., $"];
  I8 [label="I8\nC -> c C., c\nC -> c C., d"];
  I9 [label="I9\nC -> c C., $"];
  I0 -> I1 [label="C"];
  I0 -> I2 [label="S"];
  I0 -> I3 [label="c"];
  I0 -> I4 [label="d"];
  I1 -> I5 [label="C"];
  I1 -> I6 [label="c"];
  I1 -> I7 [label="d"];
  I3 -> I8 [label="C"];
  I3 -> I3 [label="c"];
  I3 -> I4 [label="d"];
  I6 -> I9 [label="C"];
  I6 -> I6 [label="c"];
  I6 -> I7 [label="d"];
}
`
	expected = strings.TrimSpace(expected)
	actual := strings.TrimSpace(graph.String())

	if expected != actual {
		t.Log("Expected:\n", expected)
		t.Log("Actual:\n", actual)
		t.Fatal("State graph does not match expectation")
	}
}
