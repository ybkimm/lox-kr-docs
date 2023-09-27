package ast_test

import (
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/lexergen/ast"
)

func TestTermRef(t *testing.T) {
	t.Run("simple1", func(t *testing.T) {
		spec, ctx := parseAndAnalyze(t, `
@macro FOO = 'foo' ;
FOOBAR = FOO 'bar' ;
`)
		it := spec.Units[0].Statements[1].(*ast.TokenRule).Expr
		nfaCons := it.NFACons(ctx)
		if ctx.Errs.HasError() {
			t.Fatalf("Failed to generate NFACons")
		}

		var nfaStr strings.Builder
		nfaCons.B.Print(&nfaStr)
		requireEqualStr(t, nfaStr.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="f"];
  1 -> 2 [label="o"];
  2 -> 3 [label="o"];
  3 -> 4 [label="ε"];
  4 -> 5 [label="b"];
  5 -> 6 [label="a"];
  6 -> 7 [label="r"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="circle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="circle"];
  7 [label="7", shape="circle"];
}
	`)
	})
	t.Run("undefined", func(t *testing.T) {
		spec, ctx := parse(t, `
@macro FOO = 'foo' ;
FOOBAR = FO0 'bar' ;
`)
		ctx.Analyze(spec)
		if !ctx.Errs.HasError() {
			t.Fatalf("Errors expected")
		}
	})
	t.Run("not-macro", func(t *testing.T) {
		spec, ctx := parse(t, `
FOO = 'foo' ;
FOOBAR = FOO 'bar' ;
`)
		ctx.Analyze(spec)
		if !ctx.Errs.HasError() {
			t.Fatalf("Errors expected")
		}
	})
}
