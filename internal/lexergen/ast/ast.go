package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type Pass int

const (
	CreateNames Pass = iota
	Check
	BuildNFA
)

var passes = []Pass{
	CreateNames,
	Check,
	BuildNFA,
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

type Spec struct {
	baseAST

	Units []*Unit
}

func (s *Spec) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, s.Units, pass)
}

type Unit struct {
	baseAST

	Statements []Statement
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, u.Statements, pass)
}

type Statement interface {
	AST
	isStatement()
}

type baseStatement struct {
	baseAST
}

func (s *baseStatement) isStatement() {}

type Mode struct {
	baseStatement

	Name  string
	Rules []Statement
}

func (m *Mode) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(m.Name, m) {
			return
		}
	}
	RunPass(ctx, m.Rules, pass)
}

type TokenRule struct {
	baseStatement

	Name    string
	Expr    *Expr
	Actions []Action
}

func (r *TokenRule) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}

type FragRule struct {
	baseStatement
	Expr    *Expr
	Actions []Action
}

func (r *FragRule) RunPass(ctx *Context, pass Pass) {
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}

type Term interface {
	AST
	NFACons(ctx *Context) *mode.NFAComposite
}

type TermCharClass struct {
	baseAST

	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *TermCharClass) RunPass(ctx *Context, pass Pass) {}

func (t *TermCharClass) NFACons(ctx *Context) *mode.NFAComposite {
	panic("not implemented")
}

type CharClassItem struct {
	baseAST
	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *Context, pass Pass) {}

type Card int

const (
	One Card = iota
	ZeroOrOne
	ZeroOrMore
	OneOrMore
)

type Action interface {
	AST
	isAction()
}

type baseAction struct {
	baseAST
}

func (a *baseAction) isAction() {}

type ActionSkip struct {
	baseAction
}

func (a *ActionSkip) RunPass(ctx *Context, pass Pass) {}

type ActionPushMode struct {
	baseAction
	Mode string

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

type ActionPopMode struct {
	baseAction
}

func (a *ActionPopMode) RunPass(ctx *Context, pass Pass) {
}
