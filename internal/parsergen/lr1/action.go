package lr1

import (
	"fmt"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type Action interface {
	actionRank() int
}

type ActionShift struct {
	State *ItemSet
}

func (a ActionShift) actionRank() int { return 1 }

type ActionReduce struct {
	Prod *grammar.Prod
}

func (a ActionReduce) actionRank() int { return 2 }

type ActionAccept struct {
}

func (a ActionAccept) actionRank() int { return 3 }

func ActionString(a Action, g *grammar.AugmentedGrammar) string {
	switch a := a.(type) {
	case ActionShift:
		return fmt.Sprintf("shift I%v", a.State.Index)
	case ActionReduce:
		return fmt.Sprintf("reduce %v", g.ProdRule(a.Prod).SymName())
	case ActionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
}

func ShiftReduce(s, r Action) (ActionShift, ActionReduce, bool) {
	shift, okShift := s.(ActionShift)
	reduce, okReduce := r.(ActionReduce)
	return shift, reduce, okShift && okReduce
}
