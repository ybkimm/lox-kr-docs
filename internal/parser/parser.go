package parser

import (
	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
)

//go:generate goyacc parser.y

func Parse(filename string, input []byte, errs *errs.Errs) *ast.Spec {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(filename, input, errs)
	p := yyNewParser()
	p.Parse(l)
	return l.Spec
}
