package parsergen

import (
	"testing"

	"github.com/dcaiafa/lox/internal/util/set"
)

// grammar4_11 creates the grammar example based on Dragon Book section 4.11.
//
//	E  -> TE'
//	E' -> +TE' | ε
//	T  -> FT'
//	T' -> *FT' | ε
//	F  -> ( E ) | id
func grammar4_11() *Grammar {
	return &Grammar{
		Terminals: []*Terminal{
			{Name: "id"},
			{Name: "+"},
			{Name: "*"},
			{Name: "("},
			{Name: ")"},
		},
		Rules: []*Rule{
			{
				Name: "E",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "T"}, {Name: "E'"}}},
				},
			},
			{
				Name: "E'",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "+"}, {Name: "T"}, {Name: "E'"}}},
					{Terms: []*Term{}},
				},
			},
			{
				Name: "T",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "F"}, {Name: "T'"}}},
				},
			},
			{
				Name: "T'",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "*"}, {Name: "F"}, {Name: "T'"}}},
					{Terms: []*Term{}},
				},
			},
			{
				Name: "F",
				Prods: []*Prod{
					{Terms: []*Term{{Name: "("}, {Name: "E"}, {Name: ")"}}},
					{Terms: []*Term{{Name: "id"}}},
				},
			},
		},
	}
}

func assertTerminalSetEq(t *testing.T, symSet *set.Set[*Terminal], symNames ...string) {
	t.Helper()

	expected := new(set.Set[string])
	expected.AddSlice(symNames)
	actual := new(set.Set[string])
	symSet.ForEach(func(terminal *Terminal) {
		actual.Add(terminal.Name)
	})

	if !actual.Equal(expected) {
		t.Log("Expected: ", expected.Elements())
		t.Log("Actual: ", actual.Elements())
		t.Fatalf("Terminal set did not match expectation")
	}
}
