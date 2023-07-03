package parsergen

import "github.com/dcaiafa/lox/internal/util/logger"

type clr struct {
	*parserTable
	logger *logger.Logger
}

func constructCLR(g *Grammar, logger *logger.Logger) *parserTable {
	clr := &clr{
		parserTable: newParserTable(g),
		logger:      logger,
	}

	initialState := newStateBuilder()
	initialState.Add(newItem(g.sp.Prods[0].index, 0, g.eof.index))
	g.closure(initialState)

	clr.states.Add(initialState.Build())

	for clr.states.Changed() {
		clr.states.ResetChanged()

		clr.states.ForEach(func(fromState *state) {
			for _, sym := range g.transitionSymbols(fromState) {
				toState := clr.gotoState(fromState, sym)
				clr.transitions.Add(fromState, toState, sym)
			}
		})
	}

	clr.logger.Logf("STATES")
	clr.logger.Logf("======")
	clr.logger.Logf("")

	clr.states.ForEach(func(s *state) {
		logger := clr.logger
		if s.Index > 0 {
			logger.Logf("")
		}
		logger.Logf("I%d:", s.Index)
		logger = logger.WithIndent()
		logger.Logf("%v", s.ToString(g))
		logger.Logf("")

		for _, item := range s.Items {
			prod := g.prods[item.Prod]
			if item.Dot == len(prod.Terms) {
				act := action{Type: actionReduce, Reduce: prod.rule}
				if prod.rule == g.sp {
					act = action{Type: actionAccept}
				}
				terminal := g.Terminals[item.Terminal]
				if !clr.actions.Add(s, terminal, act, logger) {
					clr.hasConflict = true
				}
				continue
			}
			terminal, ok := prod.Terms[item.Dot].sym.(*Terminal)
			if !ok {
				continue
			}
			shiftState := clr.transitions.Get(s, terminal)
			shiftAction := action{Type: actionShift, Shift: shiftState}
			if !clr.actions.Add(s, terminal, shiftAction, logger) {
				clr.hasConflict = true
			}
		}
	})

	return clr.parserTable
}

func (clr *clr) gotoState(i *state, x Symbol) *state {
	j := newStateBuilder()
	for _, item := range i.Items {
		prod := clr.g.prods[item.Prod]
		if item.Dot == len(prod.Terms) {
			continue
		}
		term := prod.Terms[item.Dot].sym
		if term != x {
			continue
		}
		j.Add(newItem(item.Prod, item.Dot+1, item.Terminal))
	}
	clr.g.closure(j)
	return clr.states.Add(j.Build())
}
