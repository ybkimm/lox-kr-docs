package ast

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

type ParserTermType int

const (
	ParserTermSimple     ParserTermType = iota
	ParserTermZeroOrMore                // *
	ParserTermOneOrMore                 // +
	ParserTermZeroOrOne                 // ?
	ParserTermList                      // @list
	ParserTermError                     // @error
)

func (t ParserTermType) String() string {
	switch t {
	case ParserTermSimple:
		return "Simple"
	case ParserTermZeroOrMore:
		return "ZeroOrMore"
	case ParserTermOneOrMore:
		return "OneOrMore"
	case ParserTermZeroOrOne: // ?
		return "ZeroOrOne"
	case ParserTermList: // @list
		return "List"
	case ParserTermError: // @error
		return "Error"
	default:
		return "???"
	}
}

type ParserTerm struct {
	baseAST
	Type  ParserTermType
	Name  string
	Alias string
	Child *ParserTerm
	Sep   *ParserTerm

	Symbol lr2.Term
}

func (t *ParserTerm) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if !t.check(ctx) {
			return
		}
	case Normalize:
		t.normalize(ctx)

	case Print:
		printer := ctx.CurrentPrinter.Peek()
		alias := ""
		if t.Alias != "" {
			alias = fmt.Sprintf("Alias: %v ", t.Alias)
		}
		printer.Printf(
			"Term: Name: %v%v Type: %v",
			t.Name, alias, t.Type)
		ctx.CurrentPrinter.Push(printer.WithIndent(2))
		defer ctx.CurrentPrinter.Pop()
	}

	if t.Child != nil {
		t.Child.RunPass(ctx, pass)
	}
	if t.Sep != nil {
		t.Sep.RunPass(ctx, pass)
	}
}

func (t *ParserTerm) check(ctx *Context) bool {
	if t.Name != "" {
		ast := ctx.Lookup(t.Name)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Name)
			return false
		}
		switch ast := ast.(type) {
		case *ParserRule:
			t.Symbol = ast.Rule
		case *TokenRule:
			t.Symbol = ast.Terminal
		default:
			ctx.Errs.Errorf(ctx.Position(t), "%v is not a parser or token rule", t.Name)
			return false
		}
	} else if t.Alias != "" {
		ast := ctx.LookupAlias(t.Alias)
		switch ast {
		case nil:
			ctx.Errs.Errorf(ctx.Position(t), "unknown token literal: '%v'", t.Alias)
			return false
		case AmbiguousAlias:
			ctx.Errs.Errorf(ctx.Position(t), "ambiguous token literal: '%v'", t.Alias)
			return false
		}
		t.Symbol = ast.Terminal
	}
	return true
}

func (t *ParserTerm) normalize(ctx *Context) {
	generate := func(f func(r *ParserRule)) {
		r := &ParserRule{}
		r.Name = fmt.Sprintf(
			"%s$%d",
			ctx.CurrentParserRule.Peek().Name,
			len(ctx.CurrentUnit.Peek().Statements))

		f(r)

		unit := ctx.CurrentUnit.Peek()
		unit.Statements = append(unit.Statements, r)

		t.Name = r.Name
		t.Type = ParserTermSimple
		t.Child = nil
		t.Sep = nil

		r.RunPass(ctx, CreateNames)
		r.RunPass(ctx, Check)
		r.RunPass(ctx, Normalize)
		t.RunPass(ctx, Check)
		assert.False(ctx.Errs.HasError())
	}

	switch t.Type {
	case ParserTermSimple, ParserTermError:
		// No changes required.

	case ParserTermZeroOrMore:
		// a = b c*
		//   =>
		// a = b a'
		// a' = c+ | ε
		generate(func(r *ParserRule) {
			r.Generated = lr2.GeneratedZeroOrOne
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Type: ParserTermOneOrMore, Child: t.Child},
					},
				},
				{},
			}
		})

	case ParserTermOneOrMore:
		// a = b c+
		//   =>
		// a = b a'
		// a' = a' c
		//    | c
		generate(func(r *ParserRule) {
			r.Generated = lr2.GeneratedOneOrMore
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Name: r.Name},
						t.Child,
					},
				},
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
			}
		})

	case ParserTermZeroOrOne:
		// a = b c?
		//   =>
		// a = b a'
		// a' = c | ε
		generate(func(r *ParserRule) {
			r.Generated = lr2.GeneratedZeroOrOne
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
				{},
			}
		})

	case ParserTermList:
		// a = b @list(c, sep)
		//   =>
		// a = b a'
		// a' = a' sep c
		//    | c
		generate(func(r *ParserRule) {
			r.Generated = lr2.GeneratedList
			r.Prods = []*ParserProd{
				{
					Terms: []*ParserTerm{
						{Name: r.Name},
						t.Sep,
						t.Child,
					},
				},
				{
					Terms: []*ParserTerm{
						t.Child,
					},
				},
			}
		})
	}
}
