package parser

import (
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/grammar"
)

//go:generate goyacc parser.y

func Parse(filename string, input []byte, errs *errlogger.ErrLogger) *grammar.Syntax {
	//yyDebug = 10
	yyErrorVerbose = true
	l := newLex(filename, input, errs)
	p := yyNewParser()
	p.Parse(l)
	return l.Syntax
}
