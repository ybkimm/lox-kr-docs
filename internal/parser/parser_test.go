package parser

import (
	gotoken "go/token"
	"os"
	"testing"

	"github.com/dcaiafa/lox/internal/errlogger"
)

func TestParser(t *testing.T) {
	fset := gotoken.NewFileSet()
	data, err := os.ReadFile("parser.lox")
	if err != nil {
		t.Fatal(err)
	}
	file := fset.AddFile("foo.lox", -1, len(data))
	errs := errlogger.New()
	spec := Parse(file, []byte(data), errs)
	if errs.HasError() {
		t.Fatal("Parse failed")
	}
	if spec == nil {
		t.Fatal("spec is nil")
	}
}
