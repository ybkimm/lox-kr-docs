package lr1

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
)

func ConstructCLR(
	g *grammar.AugmentedGrammar,
	logger *logger.Logger,
) *ParserTable {
	pt := NewParserTable(g)

	initialState := NewItemSet(g)
	initialState.Add(NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	initialState.Closure()

	pt.States.Add(initialState.State())

	for pt.States.Changed() {
		pt.States.ResetChanged()
		pt.States.ForEach(func(fromState *State) {
			fromItemSet := fromState.ItemSet(g)
			for _, sym := range fromItemSet.FollowingSymbols() {
				toItemSet := Goto(g, fromItemSet, sym)
				toState := pt.States.Add(toItemSet.State())
				pt.Transitions.Add(fromState, toState, sym)
			}
		})
	}

	pt.States.ForEach(func(s *State) {
		logger := logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		s.ItemSet(g).ForEach(func(item Item) {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				rule := g.ProdRule(prod)
				act := Action{
					Type:   ActionReduce,
					Reduce: rule,
				}
				if rule == g.Sprime {
					act = Action{Type: ActionAccept}
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
			shiftAction := Action{
				Type:  ActionShift,
				Shift: shiftState,
			}
			if !pt.Actions.Add(s, terminal, shiftAction, logger) {
				pt.Ambiguous = true
			}
		})
	})

	return pt
}
