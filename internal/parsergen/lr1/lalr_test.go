package lr1

import (
	"testing"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func TestConstructLALR(t *testing.T) {
	newGrammar := func() *grammar.AugmentedGrammar {
		g := &grammar.Grammar{
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
		}
		ag, err := g.ToAugmentedGrammar()
		if err != nil {
			t.Fatalf("ToAugmentedGrammar failed: %v", err)
		}
		return ag
	}

	_ = newGrammar

	/*
		t.Run("lr0Items", func(t *testing.T) {
			g := newGrammar()
			stateSet := lr0Items(g)
			report := strings.Builder{}
			logger := logger.New(&report)
			stateSet.ForEach(func(s *lr0.State) {
				logger := logger
				logger.Logf("I%d:", s.Index)
				logger = logger.WithIndent()
				logger.Logf("%v", s.ToString(g))
			})
			baseline.Assert(t, report.String())
		})
	*/

}
