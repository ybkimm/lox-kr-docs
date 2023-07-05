package lr1

import (
	"sort"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type transitionKey struct {
	From *ItemSet
	Sym  grammar.Symbol
}

type TransitionMap struct {
	transitions map[transitionKey]*ItemSet
}

func NewTransitionMap() *TransitionMap {
	return &TransitionMap{
		transitions: make(map[transitionKey]*ItemSet),
	}
}

func (m *TransitionMap) Add(from *ItemSet, to *ItemSet, sym grammar.Symbol) {
	key := transitionKey{from, sym}
	if existing, ok := m.transitions[key]; ok {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	m.transitions[key] = to
}

func (m *TransitionMap) Get(from *ItemSet, sym grammar.Symbol) *ItemSet {
	key := transitionKey{from, sym}
	toState := m.transitions[key]
	if toState == nil {
		panic("no transition")
	}
	return toState
}

func (m *TransitionMap) ForEach(fn func(from *ItemSet, to *ItemSet, sym grammar.Symbol)) {
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
