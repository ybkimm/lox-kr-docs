package parser

import (
	"fmt"
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
	errLogger := errlogger.New()
	spec, ok := Parse(file, []byte(data), errLogger)
	if !ok {
		t.Fatal("Parse failed")
	}
	if spec == nil {
		t.Fatal("spec is nil")
	}
	fmt.Printf("%+v\n", spec)
}
