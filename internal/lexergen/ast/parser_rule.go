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

	IsStart   bool
	Name      string
	Prods     []*ParserProd
	Generated ParserRuleGenerated

	RuleIndex int
}

func (r *ParserRule) RunPass(ctx *Context, pass Pass) {
	ctx.CurrentParserRule.Push(r)
	defer ctx.CurrentParserRule.Pop()

	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}
		r.RuleIndex = ctx.Grammar.AddRule(r.Name)
		if r.IsStart {
			if ctx.StartRule != nil {
				ctx.Errs.Errorf(ctx.Position(r), "@start redefined: %v", r.Name)
				ctx.Errs.Infof(
					ctx.Position(ctx.StartRule), "@start previously defined: %v",
					ctx.StartRule.Name)
				return
			}
			ctx.StartRule = r
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
