package codegen

import (
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

func (s *State) MapReduceActions(pt *lr1.ParserTable) error {
	s.ReduceMap = make(map[lr1.Action]*ReduceMethod)
	pt.States.ForEach(func(state *lr1.ItemSet) {
		pt.Actions.ForEachActionSet(pt.Grammar, state,
			func(_ grammar.Symbol, actions []lr1.Action) {
				action := actions[0]
				if action.Type != lr1.ActionReduce {
					return
				}
				rule := pt.Grammar.ProdRule(action.Prod)
				reduceName := rule.Name
				_ = reduceName
			})
	})
	return nil
}
