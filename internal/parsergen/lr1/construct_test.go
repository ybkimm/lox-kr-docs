package lr1

import (
	"fmt"
	gotoken "go/token"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
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
		errs := errlogger.New(gotoken.NewFileSet())

		ag := g.ToAugmentedGrammar(errs)
		if errs.HasError() {
			t.Fatalf("ToAugmentedGrammar failed")
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

	runAllConstructTest(t, "precedence",
		// E -> E + E
		//    | E - E
		//    | E * E
		//    | num
		&grammar.Grammar{
			Terminals: []*grammar.Terminal{
				{Name: "+"},
				{Name: "-"},
				{Name: "*"},
				{Name: "/"},
				{Name: "^"},
				{Name: "num"},
			},
			Rules: []*grammar.Rule{
				{
					Name: "E",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{term("E"), term("+"), term("E")}, Precence: 1},
						{Terms: []*grammar.Term{term("E"), term("-"), term("E")}, Precence: 1},
						{Terms: []*grammar.Term{term("E"), term("*"), term("E")}, Precence: 2},
						{Terms: []*grammar.Term{term("num")}},
					},
				},
			},
		},
	)

	runAllConstructTest(t, "repro1",
		// StmtsPP = StmtsP SEMICOLON? ;
		// StmtsP = StmtsP SEMICOLON Stmt | Stmt ;
		// Stmt = NUMBER ;
		&grammar.Grammar{
			Terminals: []*grammar.Terminal{
				{Name: "NUMBER"},
				{Name: "SEMICOLON"},
			},
			Rules: []*grammar.Rule{
				{
					Name: "StmtsPP",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{term("StmtsP"), term("SemicolonOpt")}},
					},
				},
				{
					Name: "SemicolonOpt",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{term("SEMICOLON")}},
					},
				},
				{
					Name: "StmtsP",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{term("StmtsP"), term("SEMICOLON"), term("Stmt")}},
						{Terms: []*grammar.Term{term("Stmt")}},
					},
				},
				{
					Name: "Stmt",
					Prods: []*grammar.Prod{
						{Terms: []*grammar.Term{term("NUMBER")}},
					},
				},
			},
		},
	)
}
