package lr2

import (
	"sort"

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
func Next(g *Grammar, is ItemSet) []int {
	var set set.Set[int]
	is.ForEach(func(i Item) {
		prod := g.GetProd(i.Prod)
		if i.Dot >= len(prod.Terms) {
			return
		}
		set.Add(prod.Terms[i.Dot])
	})
	syms := set.Elements()
	sort.Ints(syms)
	return syms
}
