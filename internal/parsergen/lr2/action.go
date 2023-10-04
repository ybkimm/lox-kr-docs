package lr2

import (
	"github.com/dcaiafa/lox/internal/util/array"
)

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
	actions map[int]*array.Array[*Action]
}

func (m *ActionMap) AddShift(terminal int, toState int, prod int) {
	actions := m.getMap()[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
		m.getMap()[terminal] = actions
	}
	for _, action := range actions.Elements() {
		if action.Type == ActionShift {
			if action.ShiftState != toState {
				panic("impossible shift-shift conflict")
			}
			action.Prods = append(action.Prods, prod)
			return
		}
	}
	action := &Action{
		Type:       ActionShift,
		ShiftState: toState,
		Prods:      []int{prod},
	}
	actions.Add(action)
}

func (m *ActionMap) AddReduce(terminal int, prod int) {
	actions := m.getMap()[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
		m.getMap()[terminal] = actions
	}
	for _, action := range actions.Elements() {
		if action.Type == ActionReduce {
			action.Prods = append(action.Prods, prod)
			return
		}
	}
	action := &Action{
		Type:  ActionReduce,
		Prods: []int{prod},
	}
	actions.Add(action)
}

func (m *ActionMap) AddAccept(terminal int) {
	actions := m.getMap()[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
		m.getMap()[terminal] = actions
	}
	for _, action := range actions.Elements() {
		if action.Type == ActionAccept {
			panic("impossible accept-accept conflict")
		}
	}
	action := &Action{
		Type: ActionAccept,
	}
	actions.Add(action)
}

func (m *ActionMap) getMap() map[int]*array.Array[*Action] {
	if m.actions == nil {
		m.actions = make(map[int]*array.Array[*Action])
	}
	return m.actions
}
