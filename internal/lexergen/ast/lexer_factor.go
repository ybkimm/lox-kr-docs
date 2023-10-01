package ast

import (
	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type LexerFactor struct {
	baseAST
	Terms []*LexerTermCard
}

func (f *LexerFactor) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, f.Terms, pass)
}

func (f *LexerFactor) NFACons(ctx *Context) *mode.NFAComposite {
	assert.True(len(f.Terms) > 0)

	// Build the following NFACons{b,e}:
	//
	//  T0b -...-> T0e  -> T1b -...-> T1e  -> ...  -> TNb -...-> TNe
	//   ^                                                        ^
	//   b                                                        e
	//
	// For all {Tnb, Tne} where Tn ∈ Terms.

	termCons := make([]*mode.NFAComposite, len(f.Terms))
	for i, term := range f.Terms {
		termCons[i] = term.NFACons(ctx)
	}
	for i := 0; i < len(termCons)-1; i++ {
		termCons[i].E.AddTransition(termCons[i+1].B, nfa.Epsilon)
	}

	return &mode.NFAComposite{
		B: termCons[0].B,
		E: termCons[len(termCons)-1].E,
	}
}
