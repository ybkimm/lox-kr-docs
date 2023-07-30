package main

import (
	"errors"
	"fmt"
	gotoken "go/token"
	"math"
	"strconv"
)

func Eval(expr string) (float64, error) {
	fset := gotoken.NewFileSet()
	file := fset.AddFile("expr", -1, len(expr))
	errLogger := &ErrLogger{
		Fset: fset,
	}

	var parser parser
	parser.errLogger = errLogger
	lex := newLex(file, []byte(expr), errLogger)
	ok := parser.parse(lex, errLogger)
	if !ok {
		errLogger.Error(0, errors.New("failed to parse"))
	}
	return parser.result, errLogger.Err()
}

type parser struct {
	loxParser
	errLogger _lxErrorLogger
	result    float64
}

func (p *parser) reduceS(e float64) any {
	p.result = e
	return nil
}

func (p *parser) reduceExpr_binary(left float64, op Token, right float64) float64 {
	switch op.Type {
	case PLUS:
		return left + right
	case MINUS:
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

func (p *parser) reduceExpr_paren(_ Token, e float64, _ Token) float64 {
	return e
}

func (p *parser) reduceExpr_num(e float64) float64 {
	return e
}

func (p *parser) reduceNum(num Token) float64 {
	v, err := strconv.ParseFloat(num.Str, 64)
	if err != nil {
		p.errLogger.Error(num.Pos, fmt.Errorf("invalid float: %v", err))
		return 0
	}
	return v
}

func (p *parser) reduceNum_minus(_ Token, num Token) float64 {
	v, err := strconv.ParseFloat(num.Str, 64)
	if err != nil {
		p.errLogger.Error(num.Pos, fmt.Errorf("invalid float: %v", err))
		return 0
	}
	return -v
}
