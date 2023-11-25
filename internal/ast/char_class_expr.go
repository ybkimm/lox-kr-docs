package ast

import "github.com/dcaiafa/lox/internal/lexergen/rang3"

type CharClassExpr interface {
	AST
	GetRanges() []rang3.Range
}

type CharClassBinaryExprOp int

const (
	CharClassBinaryExprAdd CharClassBinaryExprOp = 0
	CharClassBinaryExprSub CharClassBinaryExprOp = 1
)

type CharClassBinaryExpr struct {
	baseAST

	Op    CharClassBinaryExprOp
	Left  CharClassExpr
	Right CharClassExpr
}

func (e *CharClassBinaryExpr) RunPass(ctx *Context, pass Pass) {
	e.Left.RunPass(ctx, pass)
	e.Right.RunPass(ctx, pass)
}

func (e *CharClassBinaryExpr) GetRanges() []rang3.Range {
	rangesLeft := e.Left.GetRanges()
	rangesRight := e.Right.GetRanges()

	switch e.Op {
	case CharClassBinaryExprAdd:
		return rang3.Flatten(append(rangesLeft, rangesRight...), nil)
	case CharClassBinaryExprSub:
		return rang3.Subtract(rangesLeft, rangesRight)
	default:
		panic("unreachable")
	}
}
