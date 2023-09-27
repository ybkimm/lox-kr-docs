package ast

import (
	"unicode/utf8"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type TermLiteral struct {
	baseAST

	Literal string
}

func (t *TermLiteral) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if len(t.Literal) == 0 {
			ctx.Errs.Errorf(ctx.Position(t), "literal cannot be empty")
			return
		}

	case BuildNFA:
	}
}

func (t *TermLiteral) NFACons(ctx *Context) *mode.NFAComposite {
	// For a literal "foo", build the NFACons:
	//     f      o      o
	//  B --> s1 --> s2 --> E
	nfa := ctx.Mode().NFA
	nfaCons := new(mode.NFAComposite)
	nfaCons.B = nfa.NewState()
	nfaCons.E = nfaCons.B
	str := t.Literal
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		s := nfa.NewState()
		nfaCons.E.AddTransition(s, mode.Range{B: r, E: r})
		nfaCons.E = s
	}
	return nfaCons
}
