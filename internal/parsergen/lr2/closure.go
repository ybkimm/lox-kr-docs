package lr2

// Closure computes the closure of the ItemSet.
// The algorithm is summarized thusly:
//
//	For each item [A -> α.Bβ, a]:
//	 If there is a B -> γ:
//	   For each x in FIRST(βa):
//	     Add [B -> .γ, x]
func Closure(g *Grammar, i *ItemSet) *ItemSet {
	var result ItemSet
	result.AddSet(i)
	var pending ItemSet
	pending.AddSet(i)
	for !pending.Empty() {
		pendingItems := pending.Items()
		pending.Clear()
		for _, item := range pendingItems {
			prod := g.Prods[item.Prod]
			if item.Dot == len(prod.Terms) {
				continue
			}
			termB := prod.Terms[item.Dot]
			ruleB, ok := termB.(*Rule)
			if !ok {
				continue
			}
			beta := prod.Terms[item.Dot+1:]
			a := g.Terminals[item.Lookahead]
			first := First(g, append(beta, a))
			for _, prodB := range ruleB.Prods {
				first.ForEach(func(t *Terminal) {
					newItem := Item{Prod: prodB.Index, Dot: 0, Lookahead: t.Index}
					if result.Add(newItem) {
						pending.Add(newItem)
					}
				})
			}
		}
	}
	return &result
}
