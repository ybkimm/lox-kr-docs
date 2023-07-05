package lr1

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

type ParserTable struct {
	Grammar     *grammar.AugmentedGrammar
	States      *StateSet
	Transitions *TransitionMap
	Actions     *ActionMap
}

func NewParserTable(g *grammar.AugmentedGrammar) *ParserTable {
	return &ParserTable{
		Grammar:     g,
		States:      NewStateSet(),
		Transitions: NewTransitionMap(),
		Actions:     NewActionMap(),
	}
}

func (t *ParserTable) PrintStateGraph(w io.Writer) {
	fmt.Fprintln(w, `digraph G {`)
	fmt.Fprintln(w, `  rankdir="LR";`)
	t.States.ForEach(func(s *ItemSet) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			s.Index,
			fmt.Sprintf("I%d\n%v", s.Index, s.ToString(t.Grammar)),
		)
	})
	t.Transitions.ForEach(func(from, to *ItemSet, sym grammar.Symbol) {
		fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
			from.Index,
			to.Index,
			sym.SymName())
	})
	fmt.Fprintln(w, `}`)
}
