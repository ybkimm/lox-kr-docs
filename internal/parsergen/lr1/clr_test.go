package lr1

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/baseline"
	"github.com/dcaiafa/lox/internal/util/logger"
)

var cleanOutputRegex = regexp.MustCompile(`(?m)\s+$`)

func cleanOutput(s string) string {
	return cleanOutputRegex.ReplaceAllString(strings.TrimSpace(s), "")
}

func TestCLR(t *testing.T) {
	sg := &grammar.Grammar{
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
	}

	g, err := sg.ToAugmentedGrammar()
	if err != nil {
		t.Fatalf("ToAugmentedGrammar failed: %v", err)
	}

	report := strings.Builder{}
	_ = ConstructCLR(g, logger.New(&report))
	baseline.Assert(t, report.String())
}

func TestLALR(t *testing.T) {
	sg := &grammar.Grammar{
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
	}

	g, err := sg.ToAugmentedGrammar()
	if err != nil {
		t.Fatalf("ToAugmentedGrammar failed: %v", err)
	}

	report := strings.Builder{}
	pt := ConstructLALR(g, logger.New(&report))

	pt.PrintStateGraph(os.Stdout)

	baseline.Assert(t, report.String())
}
