package ast

import "github.com/dcaiafa/lox/internal/lexergen/rang3"

type CharClassItem struct {
	baseAST
	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *Context, pass Pass) {}

type CharClass struct {
	baseAST
	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *CharClass) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, t.CharClassItems, pass)
}

func (t *CharClass) GetRanges() []rang3.Range {
	ranges := make([]rang3.Range, len(t.CharClassItems))
	for i, item := range t.CharClassItems {
		ranges[i] = rang3.Range{
			B: item.From,
			E: item.To,
		}
	}
	ranges = rang3.Flatten(ranges, nil)
	if t.Neg {
		ranges = rang3.Subtract([]rang3.Range{{B: 0, E: rang3.MaxRune}}, ranges)
	}
	return ranges
}
