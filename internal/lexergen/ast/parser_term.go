package ast

type ParserTermType int

const (
	ParserTermSimple     ParserTermType = iota
	ParserTermZeroOrMore                // *
	ParserTermOneOrMore                 // +
	ParserTermZeroOrOne                 // ?
	ParserTermList                      // @list
	ParserTermError                     // @error
)

type ParserTerm struct {
	baseAST
	Type  ParserTermType
	Name  string
	Alias string
	Child *ParserTerm
	Sep   *ParserTerm

	refRule  *ParserRule
	refToken *TokenRule
}

func (t *ParserTerm) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		if t.Name != "" {
			ast := ctx.Lookup(t.Name)
			if ast == nil {
				ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Name)
				return
			}
			switch ast := ast.(type) {
			case *ParserRule:
				t.refRule = ast
			case *TokenRule:
				t.refToken = ast
			default:
				ctx.Errs.Errorf(ctx.Position(t), "%v is not a parser or token rule", t.Name)
				return
			}
		} else if t.Alias != "" {
			ast := ctx.LookupAlias(t.Name)
			switch ast {
			case nil:
				ctx.Errs.Errorf(ctx.Position(t), "unknown token literal: '%v'", t.Alias)
				return
			case AmbiguousAlias:
				ctx.Errs.Errorf(ctx.Position(t), "ambiguous token literal: '%v'", t.Alias)
				return
			}
			t.refToken = ast
		}
	}
	if t.Child != nil {
		t.Child.RunPass(ctx, pass)
	}
	if t.Sep != nil {
		t.Sep.RunPass(ctx, pass)
	}
}
