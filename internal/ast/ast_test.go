package ast_test

import (
	gotoken "go/token"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/parser"
)

func requireEqualStr(t *testing.T, actual, expected string) {
	t.Helper()
	actual = strings.TrimSpace(actual)
	expected = strings.TrimSpace(expected)
	if actual != expected {
		t.Fatalf("not equal:\nexpected:\n%v\nactual:\n%v", expected, actual)
	}
}

func parse(t *testing.T, input string) (*ast.Spec, *ast.Context) {
	fset := gotoken.NewFileSet()
	errs := errlogger.New(fset, &strings.Builder{})
	file := fset.AddFile("input.lox", -1, len(input))
	unit := parser.Parse(file, []byte(input), errs)
	if errs.HasError() {
		t.Log("Parser errors:\n",
			errs.Output().(*strings.Builder).String())
		t.Fatalf("Failed to parse")
	}
	spec := new(ast.Spec)
	spec.Units = []*ast.Unit{unit}
	ctx := ast.NewContext(fset, errs)
	return spec, ctx
}

func failParse(t *testing.T, input string, errContains string) {
	t.Helper()
	fset := gotoken.NewFileSet()
	errs := errlogger.New(fset, &strings.Builder{})
	file := fset.AddFile("input.lox", -1, len(input))
	_ = parser.Parse(file, []byte(input), errs)
	if !errs.HasError() {
		t.Fatal("Expected parser error, but it succeed")
	}
	if errContains != "" {
		errStr := errs.Output().(*strings.Builder).String()
		if !strings.Contains(errStr, errContains) {
			t.Log("Error was:\n", errStr)
			t.Fatalf("Error did not contain %q:", errContains)
		}
	}
}

func parseAndAnalyze(t *testing.T, input string) (*ast.Spec, *ast.Context) {
	t.Helper()
	spec, ctx := parse(t, input)
	if !ctx.Analyze(spec, ast.AllPasses) {
		t.Log("Analysis errors:\n",
			ctx.Errs.Output().(*strings.Builder).String())
		t.Fatalf("Failed to analyze")
	}
	return spec, ctx
}

func parseButFailAnalyze(t *testing.T, input string, errContains string) {
	t.Helper()
	spec, ctx := parse(t, input)
	if ctx.Analyze(spec, ast.AllPasses) {
		t.Fatalf("Expected Analyze to fail, but it succeeded")
		return
	}
	if errContains != "" {
		errStr := ctx.Errs.Output().(*strings.Builder).String()
		if !strings.Contains(errStr, errContains) {
			t.Log("Error was:\n", errStr)
			t.Fatalf("Error did not contain %q:", errContains)
		}
	}
}
