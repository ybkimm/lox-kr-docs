package lr0

import (
	"strings"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type State struct {
	Items []Item
	Key   string
	Index int
}

func (s *State) ItemSet(g *grammar.AugmentedGrammar) *ItemSet {
	itemSet := NewItemSet(g)
	for _, item := range s.Items {
		itemSet.Add(item)
	}
	itemSet.Closure()
	return itemSet
}

func (s *State) ToString(g *grammar.AugmentedGrammar) string {
	var str strings.Builder
	for i := range s.Items {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Items[i].ToString(g))
	}
	return str.String()
}
