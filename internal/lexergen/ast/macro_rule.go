package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

type MacroRule struct {
	baseStatement
	Name string
	Expr *LexerExpr

	cycleDetect bool
}

func (r *MacroRule) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}
	r.Expr.RunPass(ctx, pass)
}

func (r *MacroRule) NFACons(ctx *Context) *mode.NFAComposite {
	if r.cycleDetect {
		ctx.Errs.Errorf(ctx.Position(r), "macro cycle detected")
		nfaFactory := ctx.Mode().StateFactory
		nfaCons := &mode.NFAComposite{B: nfaFactory.NewState()}
		nfaCons.E = nfaCons.B
		return nfaCons
	}

	r.cycleDetect = true
	nfaCons := r.Expr.NFACons(ctx)
	r.cycleDetect = false
	return nfaCons
}
