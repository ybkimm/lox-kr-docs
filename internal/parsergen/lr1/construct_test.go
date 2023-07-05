package lr1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/baseline"
)

func runAllConstructTest(t *testing.T, name string, g *grammar.Grammar) {
	t.Run(name, func(t *testing.T) {
		runConstructTest(t, "LR", ConstructLR, g)
		runConstructTest(t, "LALR", ConstructLALR, g)
	})
}

func runConstructTest(
	t *testing.T,
	name string,
	constructFunc func(*grammar.AugmentedGrammar) *ParserTable,
	g *grammar.Grammar,
) {
	t.Run(name, func(t *testing.T) {
		ag, err := g.ToAugmentedGrammar()
		if err != nil {
			t.Fatalf("ToAugmentedGrammar failed: %v", err)
		}

		pt := constructFunc(ag)

		report := strings.Builder{}
		pt.Print(&report)
		fmt.Fprintln(&report, "")
		pt.PrintStateGraph(&report)

		baseline.Assert(t, report.String())
	})
}

func TestConstruct(t *testing.T) {
	runAllConstructTest(t, "1",
		&grammar.Grammar{
			Terminals: []*grammar.Terminal{
				{Name: "c"},
				{Name: "d"},
			},
			Rules: []*grammar.Rule{
				{
					Name: "S",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{{Name: "C"}, {Name: "C"}}},
					},
				},
				{
					Name: "C",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{{Name: "c"}, {Name: "C"}}},
						{Terms: []*grammar.Term{{Name: "d"}}},
					},
				},
			},
		})
	runAllConstructTest(t, "2",
		&grammar.Grammar{
			Terminals: []*grammar.Terminal{
				{Name: "="},
				{Name: "*"},
				{Name: "id"},
			},
			Rules: []*grammar.Rule{
				// S -> L = R | R
				{
					Name: "S",
					Prods: []*grammar.Prod{
						{
							Terms: []*grammar.Term{
								{Name: "L"}, {Name: "="}, {Name: "R"},
							},
						},
						{
							Terms: []*grammar.Term{
								{Name: "R"},
							},
						},
					},
				},

				// L -> * R | id
				{
					Name: "L",
					Prods: []*grammar.Prod{
						{
							Terms: []*grammar.Term{
								{Name: "*"}, {Name: "R"},
							},
						},
						{
							Terms: []*grammar.Term{
								{Name: "id"},
							},
						},
					},
				},

				// R -> L
				{
					Name: "R",
					Prods: []*grammar.Prod{
						{
							Terms: []*grammar.Term{
								{Name: "L"},
							},
						},
					},
				},
			},
		})
}
