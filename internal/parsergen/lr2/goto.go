package lr2

func Goto(g *Grammar, from *ItemSet, sym int) *ItemSet {
	to := new(ItemSet)
	from.ForEach(func(i Item) {
		prod := g.GetProd(i.Prod)
		if i.Dot == len(prod.Terms) {
			return
		}
		if prod.Terms[i.Dot] != sym {
			return
		}
		toItem := i
		toItem.Dot++
		to.Add(toItem)
	})
	return Closure(g, to)
}
