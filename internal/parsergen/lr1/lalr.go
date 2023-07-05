package lr1

/*
func ConstructLALR(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *ParserTable {
	pt := NewParserTable(g)

	start := NewItemSet(g)
	start.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	start.Closure()

	return nil
}
*/

/*
func lr0Items(g *grammar.AugmentedGrammar) *lr0.StateSet {
	lr0States := lr0.NewStateSet()
	first := lr0.NewItemSet(g)
	first.Add(lr0.NewItem(g, g.Sprime.Prods[0], 0))
	first.Closure()
	lr0States.Add(first.State())

	for lr0States.Changed() {
		lr0States.ResetChanged()
		lr0States.ForEach(func(fromState *lr0.State) {
			fromItemSet := fromState.ItemSet(g)
			for _, sym := range fromItemSet.FollowingSymbols() {
				toItemSet := fromItemSet.Goto(sym)
				lr0States.Add(toItemSet.State())
			}
		})
	}
	return lr0States
}

func determineLookaheads(pt *ParserTable, `
*/
