package ast

import (
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type LexerTermCharClass struct {
	baseAST
	Expr CharClassExpr
}

func (t *LexerTermCharClass) RunPass(ctx *Context, pass Pass) {
	t.Expr.RunPass(ctx, pass)
}

func (t *LexerTermCharClass) NFACons(ctx *Context) *mode.NFAComposite {
	ranges := t.Expr.GetRanges()
	nfaFactory := ctx.Mode().StateFactory
	nfaCons := &mode.NFAComposite{}
	nfaCons.B = nfaFactory.NewState()
	nfaCons.E = nfaFactory.NewState()
	for _, r := range ranges {
		rc := mode.NFAComposite{}
		rc.B = nfaFactory.NewState()
		rc.E = nfaFactory.NewState()
		rc.B.AddTransition(rc.E, r)

		nfaCons.B.AddTransition(rc.B, nfa.Epsilon)
		rc.E.AddTransition(nfaCons.E, nfa.Epsilon)
	}
	return nfaCons
}
