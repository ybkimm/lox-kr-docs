package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

// Action is the interface implemented by ASTs that define actions for tokens
// and fragments.
type Action interface {
	AST

	// GetAction returns the mode.Action corresponding to the AST.
	GetAction() mode.Action
}

// ActionDiscard is the AST for the action @discard.
//
// E.g.
//
//	// Discard whitespaces.
//	@frag [ \n\r\t]+  @discard;
type ActionDiscard struct {
	baseAST
}

func (a *ActionDiscard) RunPass(ctx *Context, pass Pass) {}

func (a *ActionDiscard) GetAction() mode.Action {
	return mode.Action{
		Type: mode.ActionDiscard,
	}
}

// ActionPushMode is the AST for the action @push_mode.
//
// Example:
//
//	@frag '"'  @push_mode(StringLiteral) ;
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

// ActionPopMode is the AST for the action @pop_mode.
//
// Example:
//
//	@mode StringLiteral {
//	  STRING = '"' @pop_mode ;
//	  @frag [\u0020-\U0010FFFF]* ;
//	}
type ActionPopMode struct {
	baseAST
}

func (a *ActionPopMode) RunPass(ctx *Context, pass Pass) {}

func (a *ActionPopMode) GetAction() mode.Action {
	return mode.Action{
		Type: mode.ActionPopMode,
	}
}
