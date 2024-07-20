package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
)

func TestExpr(t *testing.T) {
	spec, ctx := parseAndAnalyze(t, `
@lexer
FOOBAR = 'foo' | 'bar'
`)
	term := spec.Units[0].Statements[0].(*ast.TokenRule).Expr
	nfaCons := term.NFACons(ctx)
	if ctx.Errs.HasError() {
		t.Fatalf("Failed to generate NFACons")
	}

	var nfaStr strings.Builder
	nfaCons.B.Print(&nfaStr)
	requireEqualStr(t, nfaStr.String(), `
digraph G {
  rankdir="LR";
  0 -> 2 [label="ε"];
  0 -> 6 [label="ε"];
  2 -> 3 [label="f"];
  3 -> 4 [label="o"];
  4 -> 5 [label="o"];
  5 -> 1 [label="ε"];
  6 -> 7 [label="b"];
  7 -> 8 [label="a"];
  8 -> 9 [label="r"];
  9 -> 1 [label="ε"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="circle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="circle"];
  7 [label="7", shape="circle"];
  8 [label="8", shape="circle"];
  9 [label="9", shape="circle"];
}
`)
}
