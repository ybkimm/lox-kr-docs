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

type state struct {
	Items []item
	Key   string
	Index int
}

func (s *state) ToString(g *Grammar) string {
	var str strings.Builder
	for i := range s.Items {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Items[i].ToString(g))
	}
	return str.String()
}

type stateBuilder struct {
	items map[string]item
}

func newStateBuilder() *stateBuilder {
	return &stateBuilder{
		items: make(map[string]item),
	}
}

func (b *stateBuilder) Add(item item) bool {
	if _, ok := b.items[item.Key]; ok {
		return false
	}
	b.items[item.Key] = item
	return true
}

func (b *stateBuilder) Build() *state {
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

	return &state{
		Items: items,
		Key:   key.String(),
	}
}

type stateSet struct {
	stateMap map[string]*state
	states   []*state
	changed  bool
}

func newStateSet() *stateSet {
	return &stateSet{
		stateMap: make(map[string]*state),
	}
}

func (c *stateSet) Changed() bool {
	return c.changed
}

func (c *stateSet) ResetChanged() {
	c.changed = false
}

func (c *stateSet) Add(s *state) *state {
	if existing, ok := c.stateMap[s.Key]; ok {
		return existing
	}
	c.changed = true
	s.Index = len(c.states)
	c.states = append(c.states, s)
	c.stateMap[s.Key] = s
	return s
}

func (c *stateSet) ForEach(fn func(s *state)) {
	for _, state := range c.states {
		fn(state)
	}
}

type transitionKey struct {
	From *state
	Sym  Symbol
}

type transitions struct {
	transitions map[transitionKey]*state
}

func newTransitions() *transitions {
	return &transitions{
		transitions: make(map[transitionKey]*state),
	}
}

func (m *transitions) Add(from *state, to *state, sym Symbol) {
	key := transitionKey{from, sym}
	if existing, ok := m.transitions[key]; ok {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	m.transitions[key] = to
}

func (m *transitions) ForEach(fn func(from *state, to *state, sym Symbol)) {
	keys := make([]transitionKey, 0, len(m.transitions))
	for key := range m.transitions {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		switch {
		case keys[i].From.Index < keys[j].From.Index:
			return true
		case keys[i].From.Index > keys[j].From.Index:
			return false
		default:
			return keys[i].Sym.SymName() < keys[j].Sym.SymName()
		}
	})
	for _, key := range keys {
		fn(key.From, m.transitions[key], key.Sym)
	}
}
