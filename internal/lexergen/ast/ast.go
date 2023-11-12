package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type Pass int

const (
	CreateNames Pass = iota
	Check
	Normalize
	GenerateGrammar
)

const AllPasses Pass = GenerateGrammar

const Print = 1000

var passes = []Pass{
	CreateNames,
	Check,
	Normalize,
	GenerateGrammar,
}

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

type Action interface {
	AST
	GetAction() mode.Action
}

type ActionSkip struct {
	baseAST
}

func (a *ActionSkip) RunPass(ctx *Context, pass Pass) {}

func (a *ActionSkip) GetAction() mode.Action {
	return mode.Action{
		Type: mode.ActionDiscard,
	}
}

type ActionPushMode struct {
	baseAST
	Mode    string
	modeAST *Mode
}

func (a *ActionPushMode) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		ast := ctx.Lookup(a.Mode)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(a), "undefined: %v", a.Mode)
			return
		}
		modeAST, ok := ast.(*Mode)
		if !ok {
			ctx.Errs.Errorf(ctx.Position(a), "not a mode: %v", a.Mode)
			return
		}
		a.modeAST = modeAST
	}
}

func (a *ActionPushMode) GetAction() mode.Action {
	panic("not implemented")
}

type ActionPopMode struct {
	baseAST
}

func (a *ActionPopMode) RunPass(ctx *Context, pass Pass) {}

func (a *ActionPopMode) GetAction() mode.Action {
	panic("not implemented")
}

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
