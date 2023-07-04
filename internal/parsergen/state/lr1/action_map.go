package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

type actionKey struct {
	state *State
	sym   grammar.Symbol
}

type ActionMap struct {
	actions map[actionKey]Action
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		actions: make(map[actionKey]Action),
	}
}

func (m *ActionMap) Add(
	state *State,
	sym grammar.Symbol,
	action Action,
	logger *logger.Logger,
) bool {
	key := actionKey{state, sym}
	action2, exists := m.actions[key]
	if exists && action == action2 {
		return true
	}

	logger.Logf(
		"state %v with %v: %v",
		state.Index,
		sym.SymName(),
		action.String())

	if exists {
		if action2.Type > action.Type {
			action, action2 = action2, action
		}
		switch {
		case action.Type == ActionShift && action2.Type == ActionReduce:
			logger.Logf("CONFLICT: shift/reduce")
		case action.Type == ActionReduce && action2.Type == ActionReduce:
			logger.Logf("CONFLICT: reduce/reduce")
		default:
			panic("invalid conflict")
		}
		return false
	}

	m.actions[key] = action
	return true
}
