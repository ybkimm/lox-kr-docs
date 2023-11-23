package main

import (
	gotoken "go/token"
	"math"
	"strconv"

	"github.com/dcaiafa/lox/internal/util/baselexer"
)

func Eval(expr string) (float64, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile("expr", -1, len(expr))
	errs := &ErrLogger{
		Fset: fset,
	}

	onError := func(l *baselexer.Lexer) {
		errs.Errorf(l.Pos(), "unexpected character: %c", l.Peek())
	}

	var parser calcParser
	parser.errLogger = errs
	lex := baselexer.New(new(_LexerStateMachine), onError, file, []byte(expr))
	_ = parser.parse(lex)
	return parser.result, errs.Err()
}

type Token = baselexer.Token

type calcParser struct {
	lox
	errLogger *ErrLogger
	result    float64
}

func (p *calcParser) on_S__foo(e float64) any {
	p.result = e
	return nil
}

func (p *calcParser) on_expr__binary(left float64, op Token, right float64) float64 {
	switch op.Type {
	case ADD:
		return left + right
	case SUB:
		return left - right
	case MUL:
		return left * right
	case DIV:
		return left / right
	case REM:
		return math.Mod(left, right)
	case POW:
		return math.Pow(left, right)
	default:
		panic("not reached")
	}
}

func (p *calcParser) on_expr__paren(_ Token, e float64, _ Token) float64 {
	return e
}

func (p *calcParser) on_expr__num(e float64) float64 {
	return e
}

func (p *calcParser) on_num(num Token) float64 {
	v, err := strconv.ParseFloat(string(num.Str), 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return v
}

func (p *calcParser) on_num__minus(_ Token, num Token) float64 {
	v, err := strconv.ParseFloat(string(num.Str), 64)
	if err != nil {
		p.errLogger.Errorf(num.Pos, "invalid float: %v", err)
		return 0
	}
	return -v
}

func (p *calcParser) onError() {
	if p.errorToken().Type != ERROR {
		p.errLogger.Errorf(p.errorToken().Pos, "unexpected token %v", p.errorToken())
	}
}
