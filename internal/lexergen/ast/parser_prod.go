package ast

import "github.com/dcaiafa/lox/internal/parsergen/lr1"

type ParserProd struct {
	baseAST
	Terms     []*ParserTerm
	Qualifier *ProdQualifier

	Prod *lr1.Prod
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

	case GenerateGrammar:
		rule := ctx.CurrentParserRule.Peek().Rule
		terms := make([]lr1.Term, len(p.Terms))
		for i, termAst := range p.Terms {
			terms[i] = termAst.Symbol
		}
		p.Prod = ctx.Grammar.AddProd(rule, terms...)
		p.Prod.Position = ctx.Position(p)
		if p.Qualifier != nil {
			p.Prod.Precedence = p.Qualifier.Precedence
			switch p.Qualifier.Associativity {
			case Left:
				p.Prod.Associativity = lr1.Left
			case Right:
				p.Prod.Associativity = lr1.Right
			default:
				panic("not-reached")
			}
		}
	}

	RunPass(ctx, p.Terms, pass)

	if ctx.Errs.HasError() {
		return
	}
}
