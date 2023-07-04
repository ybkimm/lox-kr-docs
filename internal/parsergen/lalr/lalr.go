package lalr

/*
func ConstructParserTable(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *state.ParserTable {
}

func lr0Items(g *grammar.AugmentedGrammar) *state.StateSet {
	stateSet := state.NewStateSet()

	first := state.NewItemSet(g)
	first.Add(state.NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	first.Closure()
	stateSet.Add(first.State())

	for stateSet.Changed() {
		stateSet.ResetChanged()
		stateSet.ForEach(func (state *state.State) {
			itemSet := state.ItemSet(g)
		})
	}
}
*/
