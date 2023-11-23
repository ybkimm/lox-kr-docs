package lr1

import (
	"sort"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/base/set"
)

type symActions map[grammar.Symbol]ActionSet

type ActionMap struct {
	states map[*ItemSet]symActions
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		states: make(map[*ItemSet]symActions),
	}
}

func (m *ActionMap) Add(
	state *ItemSet,
	sym grammar.Symbol,
	action Action,
	prod *grammar.Prod,
) {
	symActs := m.states[state]
	if symActs == nil {
		symActs = make(symActions)
		m.states[state] = symActs
	}
	actionSet := symActs[sym]
	if actionSet == nil {
		actionSet = make(ActionSet)
		symActs[sym] = actionSet
	}
	actionSet.Add(action, prod)
}

func (m *ActionMap) Remove(
	state *ItemSet,
	sym grammar.Symbol,
	action Action,
) {
	symActs := m.states[state]
	if symActs == nil {
		panic("invalid state")
	}
	actionSet := symActs[sym]
	if actionSet == nil {
		panic("invalid symbol")
	}
	delete(actionSet, action)
}

func (m *ActionMap) ForEachActionSet(
	g *grammar.AugmentedGrammar,
	state *ItemSet,
	fn func(grammar.Symbol, ActionSet)) {
	symActs := m.states[state]
	if len(symActs) == 0 {
		panic("state has no actions")
	}
	syms := make([]grammar.Symbol, 0, len(symActs))
	for sym := range symActs {
		syms = append(syms, sym)
	}
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})
	for _, sym := range syms {
		fn(sym, symActs[sym])
	}
}

type ActionSet map[Action]*set.Set[*grammar.Prod]

func (s ActionSet) Add(a Action, p *grammar.Prod) {
	prodSet := s[a]
	if prodSet == nil {
		prodSet = new(set.Set[*grammar.Prod])
		s[a] = prodSet
	}
	prodSet.Add(p)
}

func (s ActionSet) Actions() []Action {
	actions := make([]Action, 0, len(s))
	for action := range s {
		actions = append(actions, action)
	}
	sort.Slice(actions, func(i, j int) bool {
		return actions[i].actionRank() < actions[j].actionRank()
	})
	return actions
}

func (s ActionSet) ProdsForAction(a Action) []*grammar.Prod {
	return s[a].Elements()
}
