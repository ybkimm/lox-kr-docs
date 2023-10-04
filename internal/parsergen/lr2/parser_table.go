package lr2

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/util/logger"
)

type ParserTable struct {
	Grammar *Grammar
	States  *StateSet
}

func NewParserTable(g *Grammar) *ParserTable {
	return &ParserTable{
		Grammar: g,
		States:  NewStateSet(),
	}
}

func (t *ParserTable) Print(w io.Writer) {
	l := logger.New(w)
	for stateIndex, state := range t.States.States() {
		l := l
		l.Logf("I%d:", stateIndex)
		l = l.WithIndent()
		l.Logf("%v", state.ToString(t.Grammar))
		transitions := t.States.GetTransitions(stateIndex)
		for _, input := range transitions.Inputs() {
			if IsRule(input) {
				to, _ := transitions.Get(input)
				l.Logf("on %v: goto I%v", t.Grammar.GetSymbolName(input), to)
			}
		}
	}
}

func (t *ParserTable) PrintGraph(w io.Writer) {
	l := logger.New(w)
	l.Logf("digraph G {")
	li := l.WithIndent()
	for stateIndex, state := range t.States.States() {
		li.Logf(
			"I%d [label=%q];",
			stateIndex,
			fmt.Sprintf("I%d\n%v", stateIndex, state.ToString(t.Grammar)))
	}
	for stateIndex := range t.States.States() {
		transitions := t.States.GetTransitions(stateIndex)
		for _, input := range transitions.Inputs() {
			toIndex, _ := transitions.Get(input)
			li.Logf(
				"I%d -> I%d [label=%q];",
				stateIndex,
				toIndex,
				t.Grammar.GetSymbolName(input))
		}
	}
	l.Logf("}")
}
