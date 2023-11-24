package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
)

func TestCharClass(t *testing.T) {
	t.Run("simple1", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
FOOBAR = [A-Za-z0-9-_] ;
`)
		it := spec.Units[0].Statements[0].(*ast.TokenRule).Expr
		nfaCons := it.NFACons(ctx)
		if ctx.Errs.HasError() {
			t.Fatalf("Failed to generate NFACons")
		}

		var nfaStr strings.Builder
		nfaCons.B.Print(&nfaStr)
		requireEqualStr(t, nfaStr.String(), `
digraph G {
  rankdir="LR";
  0 -> 2 [label="ε"];
  0 -> 4 [label="ε"];
  0 -> 6 [label="ε"];
  0 -> 8 [label="ε"];
  0 -> 10 [label="ε"];
  2 -> 3 [label="\\-"];
  3 -> 1 [label="ε"];
  4 -> 5 [label="0-9"];
  5 -> 1 [label="ε"];
  6 -> 7 [label="A-Z"];
  7 -> 1 [label="ε"];
  8 -> 9 [label="_"];
  9 -> 1 [label="ε"];
  10 -> 11 [label="a-z"];
  11 -> 1 [label="ε"];
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
  10 [label="10", shape="circle"];
  11 [label="11", shape="circle"];
}
	`)
	})
	t.Run("negated", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@lexer
FOOBAR = ~[b-d1-8\n] ;
`)
		it := spec.Units[0].Statements[0].(*ast.TokenRule).Expr
		nfaCons := it.NFACons(ctx)
		if ctx.Errs.HasError() {
			t.Fatalf("Failed to generate NFACons")
		}

		var nfaStr strings.Builder
		nfaCons.B.Print(&nfaStr)
		requireEqualStr(t, nfaStr.String(), `
digraph G {
  rankdir="LR";
  0 -> 2 [label="ε"];
  0 -> 4 [label="ε"];
  0 -> 6 [label="ε"];
  0 -> 8 [label="ε"];
  2 -> 3 [label="\\u0000-\\t"];
  3 -> 1 [label="ε"];
  4 -> 5 [label="\\u000b-0"];
  5 -> 1 [label="ε"];
  6 -> 7 [label="9-a"];
  7 -> 1 [label="ε"];
  8 -> 9 [label="e-\\u10ffff"];
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
	})
}
