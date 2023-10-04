package lr2

import "github.com/dcaiafa/lox/internal/util/set"

func ConstructLALR(g *Grammar) *ParserTable {
	t := NewParserTable(g)

	start := new(ItemSet)
	start.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
	start = Closure(g, start)
	startKey := start.LR0Key()
	t.States.Add(startKey, start)

	pendingSet := set.New[string](startKey)
	for !pendingSet.Empty() {
		pending := set.SortedElements(pendingSet)
		pendingSet.Clear()
		for _, fromKey := range pending {
			from, fromIndex := t.States.GetStateByKey(fromKey)
			for _, sym := range Next(g, *from) {
				changed := false
				to := Goto(g, from, sym)
				toKey := to.LR0Key()

				// The destination state might already exist in which case we might
				// need to complement its lookaheads.
				existingTo, existingToIndex := t.States.GetStateByKey(toKey)
				if existingTo != nil {
					for _, item := range to.Items() {
						changed = existingTo.Add(item) || changed
					}
					t.States.Transitions(fromIndex).Add(sym, existingToIndex)
				} else {
					toIndex := t.States.Add(toKey, to)
					t.States.Transitions(fromIndex).Add(sym, toIndex)
					changed = true
				}
				if changed {
					pendingSet.Add(toKey)
				}
			}
		}
	}

	createActions(t)

	return t
}

func createActions(t *ParserTable) {
	g := t.Grammar
	for stateIndex, state := range t.States.States() {
		for _, item := range state.Items() {
			prod := g.GetProd(item.Prod)
			if item.Dot == len(prod.Terms) {
				// A -> γ., x
				if item.Prod == sprimeProd {
					t.States.Actions(stateIndex).
						AddAccept(item.Lookahead)
				} else {
					t.States.Actions(stateIndex).
						AddReduce(item.Lookahead, item.Prod)
				}
			} else if terminal := prod.Terms[item.Dot]; IsTerminal(terminal) {
				// A -> α.xβ where x is a Terminal
				terminal := prod.Terms[item.Dot]
				shiftState := t.States.Transitions(stateIndex).Get(terminal)
				t.States.Actions(stateIndex).
					AddShift(terminal, shiftState, item.Prod)
			}
		}
	}
}
