package parser

import (
	gotoken "go/token"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
)

var testInput = `
@lexer

CHAR = ['\r\n\\]+ ;
`

func TestParser(t *testing.T) {
	fset := gotoken.NewFileSet()
	data := []byte(testInput)
	file := fset.AddFile("input.loxl", -1, len(data))
	errs := errlogger.New()
	Parse(file, []byte(data), errs)
	if errs.HasError() {
		t.Fatal("Parse failed")
	}
}
