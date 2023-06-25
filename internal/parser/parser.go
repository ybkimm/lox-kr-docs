package parser

import (
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/grammar"
)

//go:generate goyacc parser.y

func Parse(filename string, input []byte, errs *errs.Errs) *grammar.Spec {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(filename, input, errs)
	p := yyNewParser()
	p.Parse(l)
	return l.Spec
}
