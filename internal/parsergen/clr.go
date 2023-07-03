package parsergen

import "github.com/dcaiafa/lox/internal/util/logger"

type clr struct {
	*ParserTable
	logger *logger.Logger
}

func constructCLR(g *AugmentedGrammar, logger *logger.Logger) *ParserTable {
	clr := &clr{
		ParserTable: NewParserTable(g),
		logger:      logger,
	}

	initialState := NewStateBuilder()
	initialState.Add(NewItem(g.Sprime.Prods[0].index, 0, g.EOF.index))
	initialState.Closure(g)

	clr.States.Add(initialState.Build())

	for clr.States.Changed() {
		clr.States.ResetChanged()

		clr.States.ForEach(func(fromState *State) {
			for _, sym := range fromState.DotSymbols(g) {
				toState := clr.gotoState(fromState, sym)
				clr.Transitions.Add(fromState, toState, sym)
			}
		})
	}

	clr.logger.Logf("STATES")
	clr.logger.Logf("======")
	clr.logger.Logf("")

	clr.States.ForEach(func(s *State) {
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
				act := Action{Type: ActionReduce, Reduce: prod.rule}
				if prod.rule == g.Sprime {
					act = Action{Type: ActionAccept}
				}
				terminal := g.Terminals[item.Terminal]
				if !clr.Actions.Add(s, terminal, act, logger) {
					clr.Ambiguous = true
				}
				continue
			}
			terminal, ok := prod.Terms[item.Dot].sym.(*Terminal)
			if !ok {
				continue
			}
			shiftState := clr.Transitions.Get(s, terminal)
			shiftAction := Action{Type: ActionShift, Shift: shiftState}
			if !clr.Actions.Add(s, terminal, shiftAction, logger) {
				clr.Ambiguous = true
			}
		}
	})

	return clr.ParserTable
}

func (clr *clr) gotoState(i *State, x Symbol) *State {
	j := NewStateBuilder()
	for _, item := range i.Items {
		prod := clr.Grammar.Prods[item.Prod]
		if item.Dot == len(prod.Terms) {
			continue
		}
		term := prod.Terms[item.Dot].sym
		if term != x {
			continue
		}
		j.Add(NewItem(item.Prod, item.Dot+1, item.Terminal))
	}
	j.Closure(clr.Grammar)
	return clr.States.Add(j.Build())
}
