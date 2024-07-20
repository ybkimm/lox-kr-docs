package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	RunPass(ctx *Context, pass Pass)
	SetBounds(b Bounds)
	Bounds() Bounds
}

type baseAST struct {
	bounds Bounds
}

func (a *baseAST) SetBounds(b Bounds) {
	a.bounds = b
}

func (a *baseAST) Bounds() Bounds {
	return a.bounds
}

type Statement interface {
	AST
	isStatement()
}

type baseStatement struct {
	baseAST
}

func (s *baseStatement) isStatement() {}

type Noop struct {
	baseStatement
}

func (n *Noop) RunPass(ctx *Context, pass Pass) {}

type LexerTerm interface {
	AST
	NFACons(ctx *Context) *mode.NFAComposite
}

type Card int

const (
	One Card = iota
	ZeroOrOne
	ZeroOrMore
	OneOrMore
)

type Associativity int

const (
	Left  Associativity = 0
	Right Associativity = 1
)

type ProdQualifier struct {
	baseAST
	Precedence    int
	Associativity Associativity
}
