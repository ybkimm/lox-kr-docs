package lr2

import "github.com/dcaiafa/lox/internal/util/stack"

type ActionType int

const (
	ActionShift ActionType = iota
	ActionReduce
	ActionAccept
)

type Action struct {
	Type       ActionType
	ShiftState int
	Prods      []int
}

type ActionMap struct {
	actions map[int]*stack.Stack[*Action]
}

/*
func (m *ActionMap) AddShift(terminal int, toState int, prod int) {
	actions := m.getMap()[terminal]
	for i := range actions {
		if actions[i].Type == ActionShift {
			panic("shift-shift conflict")
		}
	}

}

func (m *ActionMap) getMap() map[int][]*Action {
	if m.actions == nil {
		m.actions = make(map[int][]*Action)
	}
	return m.actions
}
*/
