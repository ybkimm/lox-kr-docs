package lr0

import (
	"fmt"
	"testing"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/baseline"
)

func TestItemSet_Closure(t *testing.T) {
	g := &grammar.Grammar{
		Terminals: []*grammar.Terminal{
			{Name: "id"},
			{Name: "+"},
			{Name: "("},
			{Name: ")"},
		},
		Rules: []*grammar.Rule{
			// E -> E + T | T
			{
				Name: "E",
				Prods: []*grammar.Prod{
					{
						Terms: []*grammar.Term{
							{Name: "E"}, {Name: "+"}, {Name: "T"},
						},
					},
					{
						Terms: []*grammar.Term{
							{Name: "T"},
						},
					},
				},
			},

			// T -> T * F | F
			{
				Name: "T",
				Prods: []*grammar.Prod{
					{
						Terms: []*grammar.Term{
							{Name: "T"}, {Name: "+"}, {Name: "F"},
						},
					},
					{
						Terms: []*grammar.Term{
							{Name: "F"},
						},
					},
				},
			},

			// F -> (E) | id
			{
				Name: "F",
				Prods: []*grammar.Prod{
					{
						Terms: []*grammar.Term{
							{Name: "("}, {Name: "E"}, {Name: ")"},
						},
					},
					{
						Terms: []*grammar.Term{
							{Name: "id"},
						},
					},
				},
			},
		},
	}

	ag, err := g.ToAugmentedGrammar()
	if err != nil {
		t.Fatalf("ToAugmentedGrammar failed: %v", err)
	}

	run := func(name string, items ...Item) {
		t.Run(name, func(t *testing.T) {
			itemSet := NewItemSet(ag)
			for _, item := range items {
				itemSet.Add(item)
			}
			before := itemSet.String()
			itemSet.Closure()
			result := fmt.Sprintf("ItemSet:\n%v\nClosure:\n%v", before, itemSet)
			t.Log("\n" + result)
			baseline.Assert(t, result)
		})
	}

	run("1", NewItem(ag, ag.Sprime.Prods[0], 0))
	run("2", NewItem(ag, ag.GetSymbol("E").(*grammar.Rule).Prods[0], 2))
	run("3",
		NewItem(ag, ag.GetSymbol("T").(*grammar.Rule).Prods[0], 2),
		NewItem(ag, ag.GetSymbol("E").(*grammar.Rule).Prods[0], 2))
	run("4", NewItem(ag, ag.GetSymbol("F").(*grammar.Rule).Prods[1], 0))
	run("5", NewItem(ag, ag.GetSymbol("F").(*grammar.Rule).Prods[0], 1))
	run("6", NewItem(ag, ag.GetSymbol("E").(*grammar.Rule).Prods[1], 1))
}
