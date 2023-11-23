package codegen

import (
	gotoken "go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/testutil"
	"github.com/dcaiafa/lox/internal/base/baseline"
)

func TestParseLox(t *testing.T) {
	const file1 = `
@parser

@start
S = expr ;

expr = expr '+' expr  @left(1)
     | expr '-' expr  @left(1)
     | expr '*' expr  @left(2)
     | expr '/' expr  @left(2)
     | expr '%' expr  @left(2)
     | expr '^' expr  @right(3)
     | '(' expr ')'
     | num ;

num = NUM
    | '-' NUM ;
`

	const file2 = `
@lexer

NUM = [0-9]+ ;
ADD = '+' ;
SUB = '-' ;
MUL = '*' ;
DIV = '/' ;
REM = '%' ;
POW = '^' ;
O_PAREN = '(' ;
C_PAREN = ')' ;
`

	tmpDir, err := os.MkdirTemp("", "lox")
	testutil.RequireNoError(t, err)
	defer os.RemoveAll(tmpDir)

	writeFile := func(name, data string) {
		t.Helper()
		err := os.WriteFile(filepath.Join(tmpDir, name), []byte(data), 0666)
		testutil.RequireNoError(t, err)
	}
	writeFile("file1.lox", file1)
	writeFile("file2.lox", file2)
	writeFile("ignored.go", `
	package foo
	func ignored() {}
`)

	fset := gotoken.NewFileSet()
	errs := errlogger.New(os.Stderr)

	ctx := &context{
		Fset: fset,
		Dir:  tmpDir,
		Errs: errs,
	}

	ctx.ParseLox()
	testutil.RequireFalse(t, ctx.Errs.HasError())

	var output strings.Builder
	ctx.ParserGrammar.Print(&output)
	baseline.Assert(t, output.String())
}
