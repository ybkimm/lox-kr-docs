package ast

type TokenRule struct {
	baseStatement

	Name    string
	Expr    *LexerExpr
	Actions []Action
}

func (r *TokenRule) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		printer.Printf("LexerTokenRule: Name: %v", r.Name)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}
