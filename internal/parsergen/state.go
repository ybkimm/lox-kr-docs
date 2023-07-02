package parsergen

import (
	"sort"
)

type stateSet struct {
	stateMap map[string]*itemSet
	states   []*itemSet
	changed  bool
}

func newStateSet() *stateSet {
	return &stateSet{
		stateMap: make(map[string]*itemSet),
	}
}

func (c *stateSet) Changed() bool {
	return c.changed
}

func (c *stateSet) ResetChanged() {
	c.changed = false
}

func (c *stateSet) Add(state *itemSet) *itemSet {
	if existing, ok := c.stateMap[state.Key]; ok {
		return existing
	}
	c.changed = true
	state.Index = len(c.states)
	c.states = append(c.states, state)
	c.stateMap[state.Key] = state
	return state
}

func (c *stateSet) ForEach(fn func(s *itemSet)) {
	for _, state := range c.states {
		fn(state)
	}
}

type transitionKey struct {
	From *itemSet
	Sym  Symbol
}

type transitions struct {
	transitions map[transitionKey]*itemSet
}

func newTransitions() *transitions {
	return &transitions{
		transitions: make(map[transitionKey]*itemSet),
	}
}

func (m *transitions) Add(from *itemSet, to *itemSet, sym Symbol) {
	key := transitionKey{from, sym}
	if existing, ok := m.transitions[key]; ok {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	m.transitions[key] = to
}

func (m *transitions) ForEach(fn func(from *itemSet, to *itemSet, sym Symbol)) {
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
