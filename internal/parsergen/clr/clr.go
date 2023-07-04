package clr

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/state"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func ConstructParserTable(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *state.ParserTable {
	pt := state.NewParserTable(g)

	initialState := state.NewItemSet(g)
	initialState.Add(state.NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	initialState.Closure()

	pt.States.Add(initialState.State())

	for pt.States.Changed() {
		pt.States.ResetChanged()
		pt.States.ForEach(func(fromState *state.State) {
			fromItemSet := fromState.ItemSet(g)
			for _, sym := range fromItemSet.FollowingSymbols() {
				toItemSet := state.Goto(g, fromItemSet, sym)
				toState := pt.States.Add(toItemSet.State())
				pt.Transitions.Add(fromState, toState, sym)
			}
		})
	}

	pt.States.ForEach(func(s *state.State) {
		logger := logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		s.ItemSet(g).ForEach(func(item state.Item) {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				rule := g.ProdRule(prod)
				act := state.Action{
					Type:   state.ActionReduce,
					Reduce: rule,
				}
				if rule == g.Sprime {
					act = state.Action{Type: state.ActionAccept}
				}
				terminal := g.Terminals[item.Terminal]
				if !pt.Actions.Add(s, terminal, act, logger) {
					pt.Ambiguous = true
				}
				return
			}
			terminal, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Terminal)
			if !ok {
				return
			}
			shiftState := pt.Transitions.Get(s, terminal)
			shiftAction := state.Action{
				Type:  state.ActionShift,
				Shift: shiftState,
			}
			if !pt.Actions.Add(s, terminal, shiftAction, logger) {
				pt.Ambiguous = true
			}
		})
	})

	return pt
}
