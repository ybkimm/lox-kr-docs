package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
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

	return pt
}

func ConstructLALR(g *grammar.AugmentedGrammar) *ParserTable {
	pt := NewParserTable(g)

	start := NewItemSet()
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure(g)
	pt.States.Add(start.LR0Key(), start)

	changed := true
	for changed {
		changed = false
		pt.States.ForEach(func(from *ItemSet) {
			for _, sym := range from.Follow(g) {
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
			}
		})
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
				act := Action{
					Type:   ActionReduce,
					Reduce: rule,
					Prod:   prod,
				}
				if rule == g.Sprime {
					act = Action{Type: ActionAccept}
				}
				terminal := g.Terminals[item.Lookahead]
				pt.Actions.Add(s, terminal, act)
				continue
			}
			terminal, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Terminal)
			if !ok {
				continue
			}
			shiftState := pt.Transitions.Get(s, terminal)
			shiftAction := Action{
				Type:  ActionShift,
				Shift: shiftState,
				Prod:  prod,
			}
			pt.Actions.Add(s, terminal, shiftAction)
		}
	})
}

func resolveConflicts(pt *ParserTable) {
	pt.States.ForEach(
		func(state *ItemSet) {
			pt.Actions.ForEachActionSet(
				state,
				func(sym grammar.Symbol, actions []Action) {
					// We can only resolve shift/reduce conflicts.
					if len(actions) != 2 {
						return
					}
					shift, reduce := actions[0], actions[1]
					if shift.Type == ActionReduce {
						shift, reduce = reduce, shift
					}
					if shift.Type != ActionShift || reduce.Type != ActionReduce {
						return
					}

					// Both Prods involved must belong to the same Rule, and must have
					// explicit precedences.
					haveCommonRule :=
						pt.Grammar.ProdRule(shift.Prod) == pt.Grammar.ProdRule(reduce.Prod)
					if !haveCommonRule ||
						shift.Prod.Precence <= 0 ||
						reduce.Prod.Precence <= 0 {
						return
					}

					switch {
					case shift.Prod.Precence < reduce.Prod.Precence:
						pt.Actions.Remove(state, sym, shift)
					case shift.Prod.Precence > reduce.Prod.Precence:
						pt.Actions.Remove(state, sym, reduce)
					case shift.Prod == reduce.Prod && shift.Prod.Associativity == grammar.Right:
						pt.Actions.Remove(state, sym, reduce)
					default:
						pt.Actions.Remove(state, sym, shift)
					}
				})
		})
}
