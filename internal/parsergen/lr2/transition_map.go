package lr2

type TransitionMap struct {
	transitions map[int]int // key: symbol_index value: state_index
}

func (m *TransitionMap) Add(symbol, to int) {
	if m.transitions == nil {
		m.transitions = make(map[int]int)
	}
	m.transitions[symbol] = to
}
