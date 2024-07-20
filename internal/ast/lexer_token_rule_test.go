package ast_test

import (
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

func TestLexerTokenRule(t *testing.T) {
	// ActionAccept is implied with token.
	t.Run("accept", func(t *testing.T) {
		_, ctx := parseAndAnalyze(t, `
@lexer
ABC = 'abc' @pop_mode
		`)

		defaultMode := ctx.LexerModes[ast.DefaultModeName]
		if len(defaultMode.Rules) != 1 {
			t.Fatalf("Expected 1 mode rule, actual %v", len(ctx.Mode().Rules))
		}
		actions, ok := defaultMode.Rules[0].E.Data.(*mode.Actions)
		if !ok || actions == nil || len(actions.Actions) == 0 {
			t.Fatalf("Rule has no actions")
		}
		if len(actions.Actions) != 2 {
			t.Fatalf("Expected 1 action; action %v", len(actions.Actions))
		}
		if actions.Actions[0].Type != mode.ActionPopMode {
			t.Fatalf("Expected ActionPopMode; actual %v", actions.Actions[0].Type)
		}
		if actions.Actions[1].Type != mode.ActionAccept {
			t.Fatalf("Expected ActionPopMode; actual %v", actions.Actions[1].Type)
		}
	})
	t.Run("cannot_discard", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
FOO = 'foo' @discard
`, "tokens cannot be discarded")
	})
	t.Run("cannot_emit", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
FOO = 'foo' @emit(FOO)
`, "@emit is not allowed")
	})
}
