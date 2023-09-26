package ast

import (
	gotoken "go/token"
)

type Pass int

const (
	CreateNames Pass = iota
	Check
)

var passes = []Pass{
	CreateNames,
	Check,
}

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	RunPass(ctx *context, pass Pass)
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

func (s *Spec) RunPass(ctx *context, pass Pass) {
	RunPass(ctx, s.Units, pass)
}

type Unit struct {
	baseAST

	Statements []Statement
}

func (u *Unit) RunPass(ctx *context, pass Pass) {
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

func (m *Mode) RunPass(ctx *context, pass Pass) {
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

func (r *TokenRule) RunPass(ctx *context, pass Pass) {
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

func (r *FragRule) RunPass(ctx *context, pass Pass) {
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}

type MacroRule struct {
	baseStatement
	Name string
	Expr *Expr
}

func (r *MacroRule) RunPass(ctx *context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}
	r.Expr.RunPass(ctx, pass)
}

type Expr struct {
	baseTerm

	Factors []*Factor
}

func (e *Expr) RunPass(ctx *context, pass Pass) {
	RunPass(ctx, e.Factors, pass)
}

type Factor struct {
	baseAST

	Terms []*TermCard
}

func (f *Factor) RunPass(ctx *context, pass Pass) {
	RunPass(ctx, f.Terms, pass)
}

type TermCard struct {
	baseAST

	Term Term
	Card Card
}

func (t *TermCard) RunPass(ctx *context, pass Pass) {
	t.Term.RunPass(ctx, pass)
}

type Term interface {
	AST
	isTerm()
}

type baseTerm struct {
	baseAST
}

func (t *baseTerm) isTerm() {}

type TermLiteral struct {
	baseTerm

	Literal string
}

func (t *TermLiteral) RunPass(ctx *context, pass Pass) {}

type TermRef struct {
	baseTerm

	Ref string

	refAST AST
}

func (t *TermRef) RunPass(ctx *context, pass Pass) {
	switch pass {
	case Check:
		ast := ctx.Lookup(t.Ref)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Ref)
			return
		}

		switch ast.(type) {
		case *MacroRule:
		default:
			ctx.Errs.Errorf(ctx.Position(t), "invalid term: %v", t.Ref)
			return
		}
		t.refAST = ast
	}
}

type TermCharClass struct {
	baseTerm

	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *TermCharClass) RunPass(ctx *context, pass Pass) {}

type CharClassItem struct {
	baseAST

	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *context, pass Pass) {}

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

func (a *ActionSkip) RunPass(ctx *context, pass Pass) {}

type ActionPushMode struct {
	baseAction
	Mode string

	modeAST *Mode
}

func (a *ActionPushMode) RunPass(ctx *context, pass Pass) {
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

func (a *ActionPopMode) RunPass(ctx *context, pass Pass) {
}
