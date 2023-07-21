package parser

import (
	"errors"
	"fmt"
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
	spec, err := Parse(file, []byte(data))
	if err != nil {
		var unexpectedToken *_lxUnexpectedTokenError
		if errors.As(err, &unexpectedToken) {
			t.Fatalf("%v: %v", fset.Position(unexpectedToken.Token.Pos), err)
		} else {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	fmt.Println(spec)
}
