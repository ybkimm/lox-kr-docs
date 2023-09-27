package ast_test

import (
	gotoken "go/token"
	"os"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/lexergen/parser"
)

func parse(t *testing.T, input string) (*ast.Spec, *ast.Context) {
	fset := gotoken.NewFileSet()
	errs := errlogger.New()
	file := fset.AddFile("input.lox", -1, len(input))
	unit := parser.Parse(file, []byte(input), errs)
	if errs.HasError() {
		t.Fatalf("Failed to parse")
	}
	spec := new(ast.Spec)
	spec.Units = []*ast.Unit{unit}
	ctx := ast.NewContext(fset, errs)
	if !ctx.Analyze(spec) {
		t.Fatalf("Failed to analyze")
	}
	return spec, ctx
}

func TestTermLiteral(t *testing.T) {
	spec, ctx := parse(t, `
@macro foo = 'abc' ;
	`)

	term := spec.Units[0].Statements[0].(*ast.MacroRule).Expr.Factors[0].Terms[0].Term.(*ast.TermLiteral)
	nfaCons := term.NFACons(ctx)
	if ctx.Errs.HasError() {
		t.Fatalf("Failed to generate NFACons")
	}

	nfa := ctx.Mode().NFA
	nfa.Start = nfaCons.B
	ctx.Mode().NFA.Print(os.Stdout)
}
