package ast_test

import (
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

func TestLexerFragRule(t *testing.T) {
	// When there is no @discard and no @emit, ActionAccum will be inferred.
	t.Run("accum", func(t *testing.T) {
		_, ctx := parseAndAnalyze(t, `
@lexer
@frag 'abc' @pop_mode;
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
		if actions.Actions[1].Type != mode.ActionAccum {
			t.Fatalf("Expected ActionAccum; actual %v", actions.Actions[1].Type)
		}
	})

	// When @discard is specified, there will be no ActionAccum.
	t.Run("discard", func(t *testing.T) {
		_, ctx := parseAndAnalyze(t, `
@lexer
@frag 'abc' @pop_mode @discard;
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
		if actions.Actions[1].Type != mode.ActionDiscard {
			t.Fatalf("Expected ActionDiscard; actual %v", actions.Actions[1].Type)
		}
	})

	// When @emit is specified, there will be no ActionAccum.
	t.Run("emit", func(t *testing.T) {
		_, ctx := parseAndAnalyze(t, `
@lexer
FOO = 'foo';
@frag 'abc' @pop_mode @emit(FOO);
		`)

		defaultMode := ctx.LexerModes[ast.DefaultModeName]
		if len(defaultMode.Rules) != 2 {
			t.Fatalf("Expected 1 mode rule, actual %v", len(ctx.Mode().Rules))
		}
		actions, ok := defaultMode.Rules[1].E.Data.(*mode.Actions)
		if !ok || actions == nil || len(actions.Actions) == 0 {
			t.Fatalf("Rule has no actions")
		}
		if len(actions.Actions) != 2 {
			t.Fatalf("Expected 1 action; actual %v", len(actions.Actions))
		}
		if actions.Actions[0].Type != mode.ActionPopMode {
			t.Fatalf("Expected ActionPopMode; actual %v", actions.Actions[0].Type)
		}
		if actions.Actions[1].Type != mode.ActionAccept {
			t.Fatalf("Expected ActionDiscard; actual %v", actions.Actions[1].Type)
		}
	})

	t.Run("cannot_discard_and_emit", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
FOO = 'foo';
@frag 'abc' @pop_mode @emit(FOO) @discard;
		`, "cannot be discarded and emitted")
	})
	t.Run("only_one_emit", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
FOO = 'foo';
@frag 'abc' @pop_mode @emit(FOO) @emit(FOO);
		`, "one @emit")
	})
	t.Run("only_one_discard", func(t *testing.T) {
		parseButFailAnalyze(t, `
@lexer
FOO = 'foo';
@frag 'abc' @pop_mode @discard @discard;
		`, "one @discard")
	})
}
