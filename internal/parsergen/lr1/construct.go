package lr1

import (
	"sort"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/set"
)

func ConstructLR(g *grammar.AugmentedGrammar) *ParserTable {
	pt := NewParserTable(g)

	start := NewItemSet()
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure(g)
	pt.States.Add(start.LR1Key(), start)

	changed := true
	for changed {
		changed = false
		pt.States.ForEach(func(from *ItemSet) {
			for _, sym := range from.Follow(g) {
				to := from.Goto(g, sym)
				toKey := to.LR1Key()
				existing := pt.States.Get(toKey)
				if existing != nil {
					to = existing
				} else {
					pt.States.Add(toKey, to)
					changed = true
				}
				pt.Transitions.Add(from, to, sym)
			}
		})
	}

	createActions(pt)
	resolveConflicts(pt)

	return pt
}

func ConstructLALR(g *grammar.AugmentedGrammar) *ParserTable {
	pt := NewParserTable(g)

	var pending set.Set[string]

	start := NewItemSet()
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure(g)

	startKey := start.LR0Key()
	pt.States.Add(startKey, start)
	pending.Add(startKey)

	for pending.Len() > 0 {
		pendingSorted := pending.Elements()
		sort.Strings(pendingSorted)
		pending.Clear()
		for _, fromKey := range pendingSorted {
			from := pt.States.Get(fromKey)
			for _, sym := range from.Follow(g) {
				changed := false
				to := from.Goto(g, sym)
				toKey := to.LR0Key()
				existing := pt.States.Get(toKey)
				if existing != nil {
					for _, item := range to.GetItems() {
						changed = existing.Add(item) || changed
					}
					to = existing
				} else {
					pt.States.Add(toKey, to)
					changed = true
				}
				pt.Transitions.Add(from, to, sym)
				if changed {
					pending.Add(toKey)
				}
			}
		}
	}

	createActions(pt)
	resolveConflicts(pt)

	return pt
}

func createActions(pt *ParserTable) {
	g := pt.Grammar
	pt.States.ForEach(func(s *ItemSet) {
		for _, item := range s.GetItems() {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				rule := g.ProdRule(prod)
				var act Action = ActionReduce{
					Prod: prod,
				}
				if rule == g.Sprime {
					act = ActionAccept{}
				}
				terminal := g.Terminals[item.Lookahead]
				pt.Actions.Add(s, terminal, act, prod)
				continue
			}
			terminal, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Terminal)
			if !ok {
				continue
			}
			shiftState := pt.Transitions.Get(s, terminal)
			shiftAction := ActionShift{
				State: shiftState,
			}
			pt.Actions.Add(s, terminal, shiftAction, prod)
		}
	})
}

func resolveConflicts(pt *ParserTable) {
	resolveConflict := func(state *ItemSet, sym grammar.Symbol, actionSet ActionSet, actions []Action) bool {
		// We can only resolve shift/reduce conflicts.
		if len(actions) != 2 {
			return false
		}

		shift, reduce, ok := ShiftReduce(actions[0], actions[1])
		if !ok {
			shift, reduce, ok = ShiftReduce(actions[1], actions[0])
			if !ok {
				return false
			}
		}

		shiftProds := actionSet.ProdsForAction(shift)
		var shiftRule *grammar.Rule
		var shiftPrec int
		for i, shiftProd := range shiftProds {
			rule := pt.Grammar.ProdRule(shiftProd)
			if i == 0 {
				shiftRule = rule
				shiftPrec = shiftProd.Precence
			} else if shiftRule != rule || shiftPrec != shiftProd.Precence {
				return false
			}
		}

		reduceProds := actionSet.ProdsForAction(reduce)
		assert.True(len(reduceProds) == 1)
		reduceProd := reduceProds[0]
		assert.True(reduceProd == reduce.Prod)
		reducePrec := reduceProd.Precence

		// Both Prods involved must belong to the same Rule, and must have
		// explicit precedences.
		haveCommonRule := shiftRule == pt.Grammar.ProdRule(reduceProd)
		if !haveCommonRule ||
			shiftPrec <= 0 ||
			reduceProd.Precence <= 0 {
			return false
		}

		switch {
		case shiftPrec < reducePrec:
			pt.Actions.Remove(state, sym, shift)
		case shiftPrec > reducePrec:
			pt.Actions.Remove(state, sym, reduce)
		case len(shiftProds) == 1 &&
			shiftProds[0] == reduce.Prod &&
			shiftProds[0].Associativity == grammar.Right:
			pt.Actions.Remove(state, sym, reduce)
		default:
			pt.Actions.Remove(state, sym, shift)
		}

		return true
	}

	pt.States.ForEach(
		func(state *ItemSet) {
			pt.Actions.ForEachActionSet(
				pt.Grammar, state,
				func(sym grammar.Symbol, actionSet ActionSet) {
					actions := actionSet.Actions()
					if len(actions) == 1 {
						return
					}
					if !resolveConflict(state, sym, actionSet, actions) {
						pt.HasConflicts = true
					}
				},
			)
		})
}
