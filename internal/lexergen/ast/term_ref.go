package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

type TermRef struct {
	baseAST
	Ref      string
	refMacro *MacroRule
}

func (t *TermRef) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		ast := ctx.Lookup(t.Ref)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Ref)
			return
		}
		macro, ok := ast.(*MacroRule)
		if !ok {
			ctx.Errs.Errorf(ctx.Position(t), "term is not a macro: %v", t.Ref)
			return
		}
		t.refMacro = macro
	}
}

func (t *TermRef) NFACons(ctx *Context) *mode.NFAComposite {
	return t.refMacro.NFACons(ctx)
}
