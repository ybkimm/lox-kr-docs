package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
)

func TestTermLiteral(t *testing.T) {
	spec, ctx := parseAndAnalyze(t, `
@lexer
@macro FOO = 'abc' ;
	`)

	term := spec.Units[0].Statements[0].(*ast.MacroRule).Expr.Factors[0].Terms[0].Term.(*ast.LexerTermLiteral)
	nfaCons := term.NFACons(ctx)
	if ctx.Errs.HasError() {
		t.Fatalf("Failed to generate NFACons")
	}

	var nfaStr strings.Builder
	nfaCons.B.Print(&nfaStr)
	requireEqualStr(t, nfaStr.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 3 [label="c"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
}
	`)
}
