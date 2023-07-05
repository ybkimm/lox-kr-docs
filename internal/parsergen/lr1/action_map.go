package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
	"github.com/dcaiafa/lox/internal/util/set"
)

type actionKey struct {
	state *ItemSet
	sym   grammar.Symbol
}

type ActionMap struct {
	actions map[actionKey]*set.Set[Action]
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		actions: make(map[actionKey]*set.Set[Action]),
	}
}

func (m *ActionMap) Add(
	state *ItemSet,
	sym grammar.Symbol,
	action Action,
	logger *logger.Logger,
) {
	key := actionKey{state, sym}
	actionSet := m.actions[key]
	if actionSet == nil {
		actionSet = new(set.Set[Action])
		m.actions[key] = actionSet
	}
	if actionSet.Has(action) {
		return
	}

	logger.Logf(
		"state %v with %v: %v",
		state.Index,
		sym.SymName(),
		action.String())

	actionSet.Add(action)
	if actionSet.Len() > 1 {
		panic("ambiguous")
	}
}
