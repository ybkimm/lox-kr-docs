package lr2

import (
	"slices"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/util/array"
	"github.com/dcaiafa/lox/internal/util/set"
)

func ConstructLALR(g *Grammar) *ParserTable {
	t := NewParserTable(g)

	start := new(ItemSet)
	start.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
	start = Closure(g, start)
	startKey := start.LR0Key()
	t.AddState(startKey, start)

	pendingSet := set.New[string](startKey)
	for !pendingSet.Empty() {
		pending := pendingSet.Elements()
		slices.Sort(pending)
		pendingSet.Clear()
		for _, fromKey := range pending {
			from, fromIndex := t.GetStateByKey(fromKey)
			for _, sym := range Next(g, *from) {
				changed := false
				to := Goto(g, from, sym)
				toKey := to.LR0Key()

				// The destination state might already exist in which case we might
				// need to complement its lookaheads.
				existingTo, existingToIndex := t.GetStateByKey(toKey)
				if existingTo != nil {
					for _, item := range to.Items() {
						changed = existingTo.Add(item) || changed
					}
					t.Transitions(fromIndex).Add(sym, existingToIndex)
				} else {
					toIndex := t.AddState(toKey, to)
					t.Transitions(fromIndex).Add(sym, toIndex)
					changed = true
				}
				if changed {
					pendingSet.Add(toKey)
				}
			}
		}
	}

	createActions(t)
	resolveConflicts(t)

	return t
}

func createActions(t *ParserTable) {
	g := t.Grammar
	for stateIndex, state := range t.States {
		for _, item := range state.Items() {
			prod := g.GetProd(item.Prod)
			if item.Dot == len(prod.Terms) {
				// A -> γ., x
				if item.Prod == sprimeProd {
					t.Actions(stateIndex).
						AddAccept(item.Lookahead)
				} else {
					t.Actions(stateIndex).
						AddReduce(item.Lookahead, item.Prod)
				}
			} else if terminal := prod.Terms[item.Dot]; IsTerminal(terminal) {
				// A -> α.xβ where x is a Terminal
				terminal := prod.Terms[item.Dot]
				shiftState := t.Transitions(stateIndex).Get(terminal)
				t.Actions(stateIndex).
					AddShift(terminal, shiftState, item.Prod)
			}
		}
	}
}

func resolveConflicts(t *ParserTable) {
	resolveConflict := func(
		state *ItemSet,
		terminal int,
		actions *array.Array[*Action],
	) bool {
		// We can only resolve shift/reduce conflicts.
		if actions.Len() != 2 {
			return false
		}
		shift, reduce := actions.Get(0), actions.Get(1)
		if shift.Type != ActionShift || reduce.Type != ActionReduce {
			shift, reduce = reduce, shift
			if shift.Type != ActionShift || reduce.Type != ActionReduce {
				return false
			}
		}

		// A shift action can be associated with multiple productions. For example,
		// we could be shifting '+' for the following two productions:
		//  A = .'+' '-'
		//    | .'+' '*'
		// But we can only proceed with conflict resolution iff all the involved
		// productions belong to the same Rule and they all have the same
		// precendence value.
		var shiftRule int
		var shiftPrec int
		for i, prodIndex := range shift.Prods {
			prod := t.Grammar.GetProd(prodIndex)
			if i == 0 {
				shiftRule = prod.Rule
				shiftPrec = prod.Precendence
			} else if shiftRule != prod.Rule || shiftPrec != prod.Precendence {
				return false
			}
		}

		assert.True(len(reduce.Prods) == 1)
		reduceProd := t.Grammar.GetProd(reduce.Prods[0])
		reducePrec := reduceProd.Precendence

		// The production(s) associated with each action must belong to the same
		// Rule, and ust have explicit precendences.
		haveCommonRule := shiftRule == reduceProd.Rule
		if !haveCommonRule || shiftPrec <= 0 || reduceProd.Precendence <= 0 {
			return false
		}

		remove := func(action *Action) {
			actions.DeleteFunc(func(a *Action) bool {
				return a == action
			})
		}

		switch {
		case shiftPrec < reducePrec:
			remove(shift)
		case shiftPrec > reducePrec:
			remove(reduce)
		case len(shift.Prods) == 1 &&
			shift.Prods[0] == reduce.Prods[0] &&
			t.Grammar.GetProd(shift.Prods[0]).Associativity == Right:
			remove(reduce)
		default:
			remove(shift)
		}

		return true
	}

	for stateIndex, state := range t.States {
		actionMap := t.Actions(stateIndex)
		for _, terminal := range actionMap.Terminals() {
			actions := actionMap.Get(terminal)
			assert.True(!actions.Empty())
			if actions.Len() != 1 {
				if !resolveConflict(state, terminal, actions) {
					t.HasConflicts = true
				}
			}
		}
	}
}
