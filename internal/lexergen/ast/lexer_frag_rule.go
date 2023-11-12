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
		actions := make([]*mode.Action, 0, len(r.Actions))
		for _, actAST := range r.Actions {
			actions = append(actions, actAST.GetAction())
		}
		nfaCons.E.Data = actions
		ctx.CurrentLexerMode.Peek().AddRule(*nfaCons)
	}
}
