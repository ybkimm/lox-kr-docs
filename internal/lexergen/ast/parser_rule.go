package ast

import "fmt"

type ParserRuleGenerated int

const (
	ParserRuleNotGenerated ParserRuleGenerated = iota
	ParserRuleGeneratedSPrime
	ParserRuleGeneratedZeroOrOne
	ParserRuleGeneratedOneOrMore
	ParserRuleGeneratedList
)

func (g ParserRuleGenerated) String() string {
	switch g {
	case ParserRuleNotGenerated:
		return "NotGenerated"
	case ParserRuleGeneratedSPrime:
		return "SPrime"
	case ParserRuleGeneratedZeroOrOne:
		return "ZeroOrOne"
	case ParserRuleGeneratedOneOrMore:
		return "OneOrMore"
	case ParserRuleGeneratedList:
		return "List"
	default:
		return "???"
	}
}

type ParserRule struct {
	baseStatement

	Name      string
	Prods     []*ParserProd
	Generated ParserRuleGenerated
}

func (r *ParserRule) RunPass(ctx *Context, pass Pass) {
	ctx.CurrentParserRule.Push(r)
	defer ctx.CurrentParserRule.Pop()

	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		generated := ""
		if r.Generated != ParserRuleNotGenerated {
			generated = fmt.Sprintf(" Generated: %v", r.Generated)
		}
		printer.Printf("ParserRule: Name: %v%v", r.Name, generated)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	RunPass(ctx, r.Prods, pass)
}
