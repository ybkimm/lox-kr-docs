package ast

type ParserProd struct {
	baseAST
	Terms     []*ParserTerm
	Qualifier *ProdQualifier
}

func (p *ParserProd) RunPass(ctx *Context, pass Pass) {
	ctx.CurrentParserProd.Push(p)
	defer ctx.CurrentParserProd.Pop()

	switch pass {
	case Print:
		printer := ctx.CurrentPrinter.Peek()
		printer.Printf("Prod:")
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	RunPass(ctx, p.Terms, pass)
}
