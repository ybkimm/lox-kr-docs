package ast

import "github.com/dcaiafa/lox/internal/parsergen/lr2"

type ParserProd struct {
	baseAST
	Terms     []*ParserTerm
	Qualifier *ProdQualifier

	ProdIndex int
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
		ruleIndex := ctx.CurrentParserRule.Peek().RuleIndex
		terms := make([]int, len(p.Terms))
		for i, termAst := range p.Terms {
			terms[i] = termAst.SymbolIndex
		}
		prodIndex := ctx.Grammar.AddProd(ruleIndex, terms...)
		if p.Qualifier != nil {
			prod := ctx.Grammar.GetProd(prodIndex)
			prod.Precedence = p.Qualifier.Precedence
			switch p.Qualifier.Associativity {
			case Left:
				prod.Associativity = lr2.Left
			case Right:
				prod.Associativity = lr2.Right
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
