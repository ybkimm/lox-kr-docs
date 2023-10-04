package lr2

import (
	"fmt"
	"slices"

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

func (a Action) ToString(g *Grammar) string {
	switch a.Type {
	case ActionShift:
		return fmt.Sprintf("shift I%v", a.ShiftState)
	case ActionReduce:
		return fmt.Sprintf("reduce %v", g.GetSymbolName(g.GetProd(a.Prods[0]).Rule))
	case ActionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
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

func (m *ActionMap) Terminals() []int {
	inputs := make([]int, 0, len(m.actions))
	for input := range m.actions {
		inputs = append(inputs, input)
	}
	slices.Sort(inputs)
	return inputs
}

func (m *ActionMap) Get(terminal int) *array.Array[*Action] {
	actions := m.actions[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
	}
	return actions
}

func (m *ActionMap) getMap() map[int]*array.Array[*Action] {
	if m.actions == nil {
		m.actions = make(map[int]*array.Array[*Action])
	}
	return m.actions
}
