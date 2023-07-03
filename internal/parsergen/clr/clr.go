package clr

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/state"
	"github.com/dcaiafa/lox/internal/util/logger"
)

type clr struct {
	*state.ParserTable
	logger *logger.Logger
}

func constructCLR(g *grammar.AugmentedGrammar, logger *logger.Logger) *state.ParserTable {
	clr := &clr{
		ParserTable: state.NewParserTable(g),
		logger:      logger,
	}

	initialState := state.NewStateBuilder()
	initialState.Add(state.NewItem(g, g.Sprime.Prods[0], 0, g.EOF))
	initialState.Closure(g)

	clr.States.Add(initialState.Build())

	for clr.States.Changed() {
		clr.States.ResetChanged()

		clr.States.ForEach(func(fromState *state.State) {
			for _, sym := range fromState.DotSymbols(g) {
				toState := clr.gotoState(g, fromState, sym)
				clr.Transitions.Add(fromState, toState, sym)
			}
		})
	}

	clr.logger.Logf("STATES")
	clr.logger.Logf("======")
	clr.logger.Logf("")

	clr.States.ForEach(func(s *state.State) {
		logger := clr.logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		for _, item := range s.Items {
			prod := g.Prods[item.Prod]
			if item.Dot == len(prod.Terms) {
				rule := g.ProdRule(prod)
				act := state.Action{
					Type:   state.ActionReduce,
					Reduce: rule,
				}
				if rule == g.Sprime {
					act = state.Action{Type: state.ActionAccept}
				}
				terminal := g.Terminals[item.Terminal]
				if !clr.Actions.Add(s, terminal, act, logger) {
					clr.Ambiguous = true
				}
				continue
			}
			terminal, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Terminal)
			if !ok {
				continue
			}
			shiftState := clr.Transitions.Get(s, terminal)
			shiftAction := state.Action{
				Type:  state.ActionShift,
				Shift: shiftState,
			}
			if !clr.Actions.Add(s, terminal, shiftAction, logger) {
				clr.Ambiguous = true
			}
		}
	})

	return clr.ParserTable
}

func (clr *clr) gotoState(g *grammar.AugmentedGrammar, i *state.State, x grammar.Symbol) *state.State {
	j := state.NewStateBuilder()
	for _, item := range i.Items {
		prod := clr.Grammar.Prods[item.Prod]
		if item.Dot == len(prod.Terms) {
			continue
		}
		term := g.TermSymbol(prod.Terms[item.Dot])
		if term != x {
			continue
		}

		toItem := item
		toItem.Dot++
		j.Add(toItem)
	}
	j.Closure(clr.Grammar)
	return clr.States.Add(j.Build())
}
