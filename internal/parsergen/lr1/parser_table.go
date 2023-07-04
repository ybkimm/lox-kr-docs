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
	Ambiguous   bool
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
	t.States.ForEach(func(s *State) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			s.Index,
			fmt.Sprintf("I%d\n%v", s.Index, s.ToString(t.Grammar)),
		)
	})
	t.Transitions.ForEach(func(from, to *State, sym grammar.Symbol) {
		fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
			from.Index,
			to.Index,
			sym.SymName())
	})
	fmt.Fprintln(w, `}`)
}

func Goto(
	g *grammar.AugmentedGrammar,
	from *ItemSet,
	sym grammar.Symbol,
) *ItemSet {
	toState := NewItemSet(g)
	from.ForEach(func(item Item) {
		prod := g.Prods[item.Prod]
		if item.Dot == uint32(len(prod.Terms)) {
			return
		}
		term := g.TermSymbol(prod.Terms[item.Dot])
		if term != sym {
			return
		}
		toItem := item
		toItem.Dot++
		toState.Add(toItem)
	})
	toState.Closure()
	return toState
}
