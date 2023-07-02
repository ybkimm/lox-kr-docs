package parsergen

import (
	"fmt"
	"sort"
	"strings"
)

type item struct {
	Prod     int
	Dot      int
	Terminal int
	Key      string
}

func newItem(prod, dot, terminal int) item {
	key := fmt.Sprintf("%06x%04x%06x", prod, dot, terminal)
	return item{
		Prod:     prod,
		Dot:      dot,
		Terminal: terminal,
		Key:      key,
	}
}

func (i *item) ToString(g *Grammar) string {
	var str strings.Builder
	prod := g.prods[i.Prod]
	rule := prod.rule

	fmt.Fprintf(&str, "%v -> ", rule.Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if j == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(term.sym.SymName())
	}
	if i.Dot == len(prod.Terms) {
		str.WriteString(".")
	}
	str.WriteString(", ")
	terminal := g.Terminals[i.Terminal]
	str.WriteString(terminal.Name)

	return str.String()
}

type itemSet struct {
	Items []item
	Key   string
	Index int
}

func (s *itemSet) ToString(g *Grammar) string {
	var str strings.Builder
	for i := range s.Items {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Items[i].ToString(g))
	}
	return str.String()
}

type itemSetBuilder struct {
	items map[string]item
}

func newItemSetBuilder() *itemSetBuilder {
	return &itemSetBuilder{
		items: make(map[string]item),
	}
}

func (b *itemSetBuilder) Add(item item) bool {
	if _, ok := b.items[item.Key]; ok {
		return false
	}
	b.items[item.Key] = item
	return true
}

func (b *itemSetBuilder) Build() *itemSet {
	items := make([]item, 0, len(b.items))
	for _, item := range b.items {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		switch {
		case items[i].Prod < items[j].Prod:
			return true
		case items[i].Prod > items[j].Prod:
			return false
		case items[i].Dot < items[j].Dot:
			return true
		case items[i].Dot > items[j].Dot:
			return false
		default:
			return items[i].Terminal < items[j].Terminal
		}
	})

	key := strings.Builder{}
	const itemKeySize = 16
	key.Grow(16 * len(items))
	for _, item := range items {
		key.WriteString(item.Key)
	}

	return &itemSet{
		Items: items,
		Key:   key.String(),
	}
}
