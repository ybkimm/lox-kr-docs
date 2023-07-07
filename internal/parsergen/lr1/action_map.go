package lr1

import (
	"sort"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/set"
)

type symActions map[grammar.Symbol]*set.Set[Action]

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
) {
	symActs := m.states[state]
	if symActs == nil {
		symActs = make(symActions)
		m.states[state] = symActs
	}
	actionSet := symActs[sym]
	if actionSet == nil {
		actionSet = new(set.Set[Action])
		symActs[sym] = actionSet
	}
	if actionSet.Has(action) {
		return
	}
	actionSet.Add(action)
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
	actionSet.Remove(action)
}

func (m *ActionMap) ForEachActionSet(
	g *grammar.AugmentedGrammar,
	state *ItemSet,
	fn func(grammar.Symbol, []Action)) {
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
		actions := symActs[sym].Elements()
		sort.Slice(actions, func(i, j int) bool {
			switch {
			case actions[i].Type < actions[j].Type:
				return true
			case actions[i].Type > actions[j].Type:
				return false
			default:
				return g.ProdIndex(actions[i].Prod) < g.ProdIndex(actions[j].Prod)
			}
		})
		fn(sym, actions)
	}
}
