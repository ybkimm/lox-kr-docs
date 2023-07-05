package lr1

import (
	"fmt"
	"sort"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
	"github.com/dcaiafa/lox/internal/util/set"
)

type stateActions map[grammar.Symbol]*set.Set[Action]

type ActionMap struct {
	states map[*ItemSet]stateActions
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		states: make(map[*ItemSet]stateActions),
	}
}

func (m *ActionMap) Add(
	state *ItemSet,
	sym grammar.Symbol,
	action Action,
	logger *logger.Logger,
) {
	stateActs := m.states[state]
	if stateActs == nil {
		stateActs = make(stateActions)
		m.states[state] = stateActs
	}
	actionSet := stateActs[sym]
	if actionSet == nil {
		actionSet = new(set.Set[Action])
		stateActs[sym] = actionSet
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

func (m *ActionMap) ForEachActionSet(
	state *ItemSet,
	fn func(grammar.Symbol, []Action)) {
	stateActions := m.states[state]
	if len(stateActions) == 0 {
		panic("state has no actions")
	}
	syms := make([]grammar.Symbol, 0, len(stateActions))
	for sym := range stateActions {
		syms = append(syms, sym)
	}
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})
	for _, sym := range syms {
		actions := stateActions[sym].Elements()
		sort.Slice(actions, func(i, j int) bool {
			symName := func(a Action) string {
				switch {
				case a.Reduce != nil:
					return a.Reduce.SymName()
				case a.Shift != nil:
					return fmt.Sprintf("I%v", a.Shift.Index)
				default:
					return ""
				}
			}
			switch {
			case actions[i].Type < actions[j].Type:
				return true
			case actions[i].Type > actions[j].Type:
				return false
			default:
				return symName(actions[i]) < symName(actions[j])
			}
		})
		fn(sym, actions)
	}
}
