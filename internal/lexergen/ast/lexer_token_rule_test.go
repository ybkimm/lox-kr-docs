package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/testutil"
)

func TestLexerTokenRule(t *testing.T) {
	t.Run("cant-discard", func(t *testing.T) {
		spec, ctx := parse(t, `
@lexer
FOO = 'foo' @discard ;
`)
		ctx.Analyze(spec, ast.AllPasses)
		testutil.RequireTrue(t, ctx.Errs.HasError())
		msg := ctx.Errs.Output().(*strings.Builder).String()
		testutil.RequireTrue(t, strings.Contains(msg, "tokens cannot be discarded"))
	})
}
