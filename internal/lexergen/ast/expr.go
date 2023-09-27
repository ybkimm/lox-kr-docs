package ast

import (
	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type Expr struct {
	baseAST

	Factors []*Factor
}

func (e *Expr) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, e.Factors, pass)
}

func (e *Expr) NFACons(ctx *Context) *mode.NFAComposite {
	assert.True(len(e.Factors) > 0)

	if len(e.Factors) == 1 {
		return e.Factors[0].NFACons(ctx)
	}

	// Build the following NFACons{b,e}:
	//        ε                      ε
	//      +----> F0b --> ... F0e -----+
	//      | ε                      ε  |
	//  b --+----> F1b --> ... F1e -----+--> e
	//      | ε                      ε  |
	//      +----> F2b --> ... F2e -----+
	//
	// For all {Fnb, Fne} where Fn ∈ Factors.
	nfaFactory := ctx.Mode().NFA
	nfaCons := &mode.NFAComposite{}
	nfaCons.B = nfaFactory.NewState()
	nfaCons.E = nfaFactory.NewState()
	for _, f := range e.Factors {
		fc := f.NFACons(ctx)
		nfaCons.B.AddTransition(fc.B, nfa.Epsilon)
		fc.E.AddTransition(nfaCons.E, nfa.Epsilon)
	}
	return nfaCons
}
