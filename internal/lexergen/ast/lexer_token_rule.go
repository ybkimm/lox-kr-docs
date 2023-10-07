package ast

type TokenRule struct {
	baseStatement

	Name    string
	Expr    *LexerExpr
	Actions []Action

	TerminalIndex int
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
		r.TerminalIndex = ctx.Grammar.AddTerminal(r.Name)

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
		ctx.CurrentLexerMode.Peek().AddRule(*nfaCons)
	}
}
