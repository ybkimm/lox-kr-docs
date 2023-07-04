package lr1

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type Item struct {
	Prod uint32
	Dot  uint32
}

func NewItem(g *grammar.AugmentedGrammar, prod *grammar.Prod, dot int) Item {
	return Item{
		Prod: uint32(g.ProdIndex(prod)),
		Dot:  uint32(dot),
	}
}

func (i Item) IsKernel() bool {
	// Assumes that [S' -> S] is Prod 0.
	return i.Dot != 0 || i.Prod == 0
}

func (i Item) ToString(g *grammar.AugmentedGrammar) string {
	var str strings.Builder
	prod := g.Prods[i.Prod]
	rule := g.ProdRule(prod)

	fmt.Fprintf(&str, "%v -> ", rule.Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if uint32(j) == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(g.TermSymbol(term).SymName())
	}
	if i.Dot == uint32(len(prod.Terms)) {
		str.WriteString(".")
	}

	return str.String()
}

func sortItems(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		switch {
		case items[i].Prod < items[j].Prod:
			return true
		case items[i].Prod > items[j].Prod:
			return false
		default:
			return items[i].Dot < items[j].Dot
		}
	})
}
