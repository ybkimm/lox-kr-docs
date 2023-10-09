package lr2

import (
	"cmp"
	"slices"
)

type TransitionMap struct {
	transitions map[Term]*ItemSet // symbol => state
}

func (m *TransitionMap) Add(symbol Term, to *ItemSet) {
	if m.transitions == nil {
		m.transitions = make(map[Term]*ItemSet)
	}
	m.transitions[symbol] = to
}

func (m *TransitionMap) Get(input Term) *ItemSet {
	to, ok := m.transitions[input]
	if !ok {
		panic("no transition for input")
	}
	return to
}

func (m *TransitionMap) Inputs() []Term {
	inputs := make([]Term, 0, len(m.transitions))
	for input := range m.transitions {
		inputs = append(inputs, input)
	}
	slices.SortFunc(inputs, func(a, b Term) int {
		return cmp.Compare(a.TermName(), b.TermName())
	})
	return inputs
}
