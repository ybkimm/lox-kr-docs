package lr2

import "slices"

type TransitionMap struct {
	transitions map[int]int // key: symbol_index value: state_index
}

func (m *TransitionMap) Add(symbol, to int) {
	if m.transitions == nil {
		m.transitions = make(map[int]int)
	}
	m.transitions[symbol] = to
}

func (m *TransitionMap) Get(input int) int {
	to, ok := m.transitions[input]
	if !ok {
		panic("no transition for input")
	}
	return to
}

func (m *TransitionMap) Inputs() []int {
	inputs := make([]int, 0, len(m.transitions))
	for input := range m.transitions {
		inputs = append(inputs, input)
	}
	slices.Sort(inputs)
	return inputs
}
