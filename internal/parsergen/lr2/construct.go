package lr2

import (
	"sort"

	"github.com/dcaiafa/lox/internal/util/set"
)

func ConstructLALR(g *Grammar) *ParserTable {
	pt := NewParserTable(g)

	var pending set.Set[string]
	start := new(ItemSet)
	start.Add(Item{Prod: sprimeProd, Dot: 0, Lookahead: EOF})
	*start = Closure(g, *start)

	for !pending.Empty() {
		pendingSorted := pending.Elements()
		sort.Strings(pendingSorted)
		pending.Clear()
		for _, fromKey := range pendingSorted {
			from, fromIndex := pt.States.GetStateByKey(fromKey)
			for _, sym := range Next(g, *from) {
				changed := false
				to := Goto(g, *from, sym)
				toKey := to.LR0Key()
				existing, _ := pt.States.GetStateByKey(toKey)
				if existing != nil {
					for _, item := range to.Items() {
						changed = existing.Add(item) || changed
					}
					to = existing
				}

			}
		}
	}
}
