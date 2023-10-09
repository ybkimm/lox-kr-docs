package lr2

import (
	"cmp"
	"slices"

	"github.com/dcaiafa/lox/internal/util/set"
)

// Next returns the list of unique symbols on the right side of the dot in Items
// in a ItemSet.
//
// For example, given the following ItemSet(I):
//
//	C = c .C
//	C = .c C
//	C = .d
//
// Then Next(I) is [C, c, d].
func Next(g *Grammar, is ItemSet) []Term {
	var set set.Set[Term]
	is.ForEach(func(i Item) {
		prod := g.Prods[i.Prod]
		if i.Dot >= len(prod.Terms) {
			return
		}
		set.Add(prod.Terms[i.Dot])
	})
	syms := set.Elements()
	slices.SortFunc(syms, func(a, b Term) int {
		return cmp.Compare(a.TermName(), b.TermName())
	})
	return syms
}
