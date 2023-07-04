package clr

import (
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
	_ = ConstructParserTable(g, logger.New(&report))
	baseline.Compare(t, report.String())
}
