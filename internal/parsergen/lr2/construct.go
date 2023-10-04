package lr2

import "github.com/dcaiafa/lox/internal/util/set"

func ConstructLALR(g *Grammar) *ParserTable {
	pt := NewParserTable(g)

	start := new(ItemSet)
	start.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
	start = Closure(g, start)
	startKey := start.LR0Key()
	pt.States.Add(startKey, start)

	pendingSet := set.New[string](startKey)
	for !pendingSet.Empty() {
		pending := set.SortedElements(pendingSet)
		pendingSet.Clear()
		for _, fromKey := range pending {
			from, fromIndex := pt.States.GetStateByKey(fromKey)
			for _, sym := range Next(g, *from) {
				changed := false
				to := Goto(g, from, sym)
				toKey := to.LR0Key()

				// The destination state might already exist in which case we might
				// need to complement its lookaheads.
				existingTo, existingToIndex := pt.States.GetStateByKey(toKey)
				if existingTo != nil {
					for _, item := range to.Items() {
						changed = existingTo.Add(item) || changed
					}
					pt.States.AddTransition(fromIndex, sym, existingToIndex)
				} else {
					toIndex := pt.States.Add(toKey, to)
					pt.States.AddTransition(fromIndex, sym, toIndex)
					changed = true
				}
				if changed {
					pendingSet.Add(toKey)
				}
			}
		}
	}

	return pt
}
