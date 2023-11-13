package ast

import (
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type LexerTermCharClass struct {
	baseAST

	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *LexerTermCharClass) RunPass(ctx *Context, pass Pass) {}

func (t *LexerTermCharClass) NFACons(ctx *Context) *mode.NFAComposite {
	ranges := make([]mode.Range, len(t.CharClassItems))
	for i, item := range t.CharClassItems {
		ranges[i] = mode.Range{
			B: item.From,
			E: item.To,
		}
	}
	ranges = mode.FlattenRanges(ranges)
	if t.Neg {
		ranges = mode.NegateRanges(ranges)
	}

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

type CharClassItem struct {
	baseAST
	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *Context, pass Pass) {}
