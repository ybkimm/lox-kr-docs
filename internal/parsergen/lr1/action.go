package lr1

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type ActionType int

const (
	ActionShift ActionType = iota
	ActionReduce
	ActionAccept
)

type Action struct {
	Type   ActionType
	Reduce *grammar.Rule
	Shift  *ItemSet
}

func (a Action) String() string {
	switch a.Type {
	case ActionShift:
		return fmt.Sprintf("shift I%v", a.Shift.Index)
	case ActionReduce:
		return fmt.Sprintf("reduce %v", a.Reduce.SymName())
	case ActionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
}
