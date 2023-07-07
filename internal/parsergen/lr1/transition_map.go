package lr1

import (
	"sort"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type symTransitions map[grammar.Symbol]*ItemSet

type TransitionMap struct {
	states map[*ItemSet]symTransitions
}

func NewTransitionMap() *TransitionMap {
	return &TransitionMap{
		states: make(map[*ItemSet]symTransitions),
	}
}

func (m *TransitionMap) Add(from *ItemSet, to *ItemSet, sym grammar.Symbol) {
	symTrans := m.states[from]
	if symTrans == nil {
		symTrans = make(symTransitions)
		m.states[from] = symTrans
	}
	existing := symTrans[sym]
	if existing != nil {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	symTrans[sym] = to
}

func (m *TransitionMap) Get(from *ItemSet, sym grammar.Symbol) *ItemSet {
	symTrans := m.states[from]
	if symTrans == nil {
		panic("invalid state")
	}
	toState := symTrans[sym]
	if toState == nil {
		panic("no transition")
	}
	return toState
}

func (m *TransitionMap) ForEach(from *ItemSet, fn func(sym grammar.Symbol, to *ItemSet)) {
	symTrans := m.states[from]
	if symTrans == nil {
		return
	}
	keys := make([]grammar.Symbol, 0, len(symTrans))
	for sym := range symTrans {
		keys = append(keys, sym)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].SymName() < keys[j].SymName()
	})
	for _, sym := range keys {
		to := symTrans[sym]
		fn(sym, to)
	}
}
