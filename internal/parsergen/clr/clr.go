package clr

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/state/lr1"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func ConstructParserTable(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *lr1.ParserTable {
	pt := lr1.NewParserTable(g)

	initialState := lr1.NewItemSet(g)
	initialState.Add(lr1.NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	initialState.Closure()

	pt.States.Add(initialState.State())

	for pt.States.Changed() {
		pt.States.ResetChanged()
		pt.States.ForEach(func(fromState *lr1.State) {
			fromItemSet := fromState.ItemSet(g)
			for _, sym := range fromItemSet.FollowingSymbols() {
				toItemSet := lr1.Goto(g, fromItemSet, sym)
				toState := pt.States.Add(toItemSet.State())
				pt.Transitions.Add(fromState, toState, sym)
			}
		})
	}

	pt.States.ForEach(func(s *lr1.State) {
		logger := logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		s.ItemSet(g).ForEach(func(item lr1.Item) {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				rule := g.ProdRule(prod)
				act := lr1.Action{
					Type:   lr1.ActionReduce,
					Reduce: rule,
				}
				if rule == g.Sprime {
					act = lr1.Action{Type: lr1.ActionAccept}
				}
				terminal := g.Terminals[item.Lookahead]
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
			shiftAction := lr1.Action{
				Type:  lr1.ActionShift,
				Shift: shiftState,
			}
			if !pt.Actions.Add(s, terminal, shiftAction, logger) {
				pt.Ambiguous = true
			}
		})
	})

	return pt
}
