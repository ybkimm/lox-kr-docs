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
		for _, actAST := range r.Actions {
			act := actAST.GetAction()
			if act.Type == mode.ActionDiscard {
				hasDiscard = true
			}
			actions.Actions = append(actions.Actions, act)
		}
		if !hasDiscard {
			actions.Actions = append(actions.Actions, mode.Action{
				Type: mode.ActionAccum,
			})
		}
		nfaCons.E.Data = actions
		ctx.CurrentLexerMode.Peek().AddRule(*nfaCons)
	}
}
