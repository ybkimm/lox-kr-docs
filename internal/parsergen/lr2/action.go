package lr2

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/dcaiafa/lox/internal/base/array"
)

type ActionType int

const (
	ActionShift ActionType = iota
	ActionReduce
	ActionAccept
)

type Action struct {
	Type       ActionType
	ShiftState *ItemSet
	Prods      []*Prod
}

func (a Action) ToString(g *Grammar) string {
	switch a.Type {
	case ActionShift:
		return fmt.Sprintf("shift I%v", a.ShiftState.Index)
	case ActionReduce:
		return fmt.Sprintf("reduce %v", a.Prods[0].Rule.Name)
	case ActionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
}

type ActionMap struct {
	actions map[*Terminal]*array.Array[*Action]
}

func (m *ActionMap) AddShift(terminal *Terminal, toState *ItemSet, prod *Prod) {
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
		Prods:      []*Prod{prod},
	}
	actions.Add(action)
}

func (m *ActionMap) AddReduce(terminal *Terminal, prod *Prod) {
	actions := m.getMap()[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
		m.getMap()[terminal] = actions
	}
	action := &Action{
		Type:  ActionReduce,
		Prods: []*Prod{prod},
	}
	actions.Add(action)
}

func (m *ActionMap) AddAccept(terminal *Terminal) {
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

func (m *ActionMap) Terminals() []*Terminal {
	inputs := make([]*Terminal, 0, len(m.actions))
	for input := range m.actions {
		inputs = append(inputs, input)
	}
	slices.SortFunc(inputs, func(a, b *Terminal) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return inputs
}

func (m *ActionMap) Get(terminal *Terminal) *array.Array[*Action] {
	actions := m.actions[terminal]
	if actions == nil {
		actions = new(array.Array[*Action])
	}
	return actions
}

func (m *ActionMap) getMap() map[*Terminal]*array.Array[*Action] {
	if m.actions == nil {
		m.actions = make(map[*Terminal]*array.Array[*Action])
	}
	return m.actions
}
