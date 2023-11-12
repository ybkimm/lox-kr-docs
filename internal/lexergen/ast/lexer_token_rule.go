package ast

import (
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

type TokenRule struct {
	baseStatement

	Name    string
	Expr    *LexerExpr
	Actions []Action

	Terminal *lr2.Terminal
}

func (r *TokenRule) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}

		// If the token rule is just a simple literal:
		// E.g.: ADD = '+'
		// Then use the literal as an alias for the token.
		if len(r.Expr.Factors) == 1 &&
			len(r.Expr.Factors[0].Terms) == 1 &&
			r.Expr.Factors[0].Terms[0].Card == One {
			if literal, ok := r.Expr.Factors[0].Terms[0].Term.(*LexerTermLiteral); ok {
				ctx.CreateAlias(literal.Literal, r)
			}
		}
		r.Terminal = ctx.Grammar.AddTerminal(r.Name)

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		printer.Printf("LexerTokenRule: Name: %v", r.Name)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)

	switch pass {
	case GenerateGrammar:
		nfaCons := r.Expr.NFACons(ctx)
		nfaCons.E.Accept = true
		actions := make([]*mode.Action, 0, len(r.Actions)+1)
		for _, actAST := range r.Actions {
			actions = append(actions, actAST.GetAction())
		}
		actions = append(actions, &mode.Action{
			Type:     mode.ActionEmit,
			Terminal: r.Terminal.Index,
			Pos:      r.bounds.Begin,
		})
		nfaCons.E.Data = actions
		ctx.CurrentLexerMode.Peek().AddRule(*nfaCons)
	}
}
