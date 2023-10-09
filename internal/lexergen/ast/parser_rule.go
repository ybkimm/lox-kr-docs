package ast

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

type ParserRule struct {
	baseStatement

	IsStart   bool
	Name      string
	Prods     []*ParserProd
	Generated lr2.Generated

	Rule *lr2.Rule
}

func (r *ParserRule) RunPass(ctx *Context, pass Pass) {
	ctx.CurrentParserRule.Push(r)
	defer ctx.CurrentParserRule.Pop()

	switch pass {
	case CreateNames:
		ctx.HasParserRules = true
		if !ctx.RegisterName(r.Name, r) {
			return
		}
		r.Rule = ctx.Grammar.AddRule(r.Name)
		r.Rule.Position = ctx.Position(r)
		r.Rule.Generated = r.Generated

		if r.IsStart {
			if ctx.StartParserRule != nil {
				ctx.Errs.Errorf(ctx.Position(r), "@start redefined: %v", r.Name)
				ctx.Errs.Infof(
					ctx.Position(ctx.StartParserRule), "@start previously defined: %v",
					ctx.StartParserRule.Name)
				return
			}
			ctx.StartParserRule = r
		}

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		generated := ""
		if r.Generated != lr2.NotGenerated {
			generated = fmt.Sprintf(" Generated: %v", r.Generated)
		}
		printer.Printf("ParserRule: Name: %v%v", r.Name, generated)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	RunPass(ctx, r.Prods, pass)
}
