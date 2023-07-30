package parser

import (
	gotoken "go/token"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	fset := gotoken.NewFileSet()
	data, err := os.ReadFile("parser.lox")
	if err != nil {
		t.Fatal(err)
	}
	file := fset.AddFile("foo.lox", -1, len(data))
	errLogger := &ErrLogger{
		Fset: fset,
	}
	spec, ok := Parse(file, []byte(data), errLogger)
	if !ok {
		t.Fatal("Parse failed")
	}
	if spec == nil {
		t.Fatal("spec is nil")
	}
}
