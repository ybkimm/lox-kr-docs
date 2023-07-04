package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func ConstructLALR(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *ParserTable {

	// Construct sets of LR(0) items.
	lr0States := lr0.NewStateSet()
	first := NewItemSet(g)
	first.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	first.Closure()
	stateSet.Add(first.State())

	/*
		for stateSet.Changed() {
			stateSet.ResetChanged()
			stateSet.ForEach(func(state *state.State) {
				itemSet := state.ItemSet(g)
			})
		}
	*/

}
