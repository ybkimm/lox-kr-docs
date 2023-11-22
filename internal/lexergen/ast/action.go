package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

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
	return mode.Action{
		Type: mode.ActionPushMode,
		Mode: a.Mode,
	}
}

type ActionPopMode struct {
	baseAST
}

func (a *ActionPopMode) RunPass(ctx *Context, pass Pass) {}

func (a *ActionPopMode) GetAction() mode.Action {
	return mode.Action{
		Type: mode.ActionPopMode,
	}
}
