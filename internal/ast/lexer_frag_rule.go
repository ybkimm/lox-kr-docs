package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

type FragRule struct {
	baseStatement
	Expr    *LexerExpr
	Actions []Action
}

func (r *FragRule) RunPass(ctx *Context, pass Pass) {
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)

	switch pass {
	case GenerateGrammar:
		nfaCons := r.Expr.NFACons(ctx)
		nfaCons.E.Accept = true
		actions := &mode.Actions{
			Pos: r.Bounds().Begin,
		}

		hasDiscard := false
		hasEmit := false
		for _, actAST := range r.Actions {
			act := actAST.GetAction()
			switch act.Type {
			case mode.ActionDiscard:
				hasDiscard = true
			case mode.ActionAccept:
				hasEmit = true
			}
			actions.Actions = append(actions.Actions, act)
		}

		if !hasDiscard {
			actions.Actions = append(actions.Actions, mode.Action{
				Type: mode.ActionAccum,
			})
		}

		if hasDiscard && hasEmit {
			ctx.Errs.Errorf(
				ctx.Position(r),
				"@frag cannot be discarded and emitted at the same time")
			return
		}

		nfaCons.E.Data = actions
		ctx.CurrentLexerMode.Peek().AddRule(*nfaCons)
	}
}
