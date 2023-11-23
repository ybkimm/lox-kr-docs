package lr2

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/base/logger"
)

type ParserTable struct {
	Grammar      *Grammar
	HasConflicts bool
	States       []*ItemSet

	stateMap    map[string]int              // state-key => state
	transitions map[*ItemSet]*TransitionMap // state => transitions
	actions     map[*ItemSet]*ActionMap     // state => actions
}

func NewParserTable(g *Grammar) *ParserTable {
	return &ParserTable{
		Grammar:     g,
		stateMap:    make(map[string]int),
		transitions: make(map[*ItemSet]*TransitionMap),
		actions:     make(map[*ItemSet]*ActionMap),
	}
}

func (c *ParserTable) GetStateByKey(key string) *ItemSet {
	stateIndex, ok := c.stateMap[key]
	if !ok {
		return nil
	}
	return c.States[stateIndex]
}

func (c *ParserTable) GetStateByIndex(stateIndex int) *ItemSet {
	return c.States[stateIndex]
}

func (c *ParserTable) AddState(key string, s *ItemSet) {
	if _, ok := c.stateMap[key]; ok {
		panic("state already exists")
	}
	c.States = append(c.States, s)
	c.stateMap[key] = len(c.States) - 1
	s.Index = len(c.States) - 1
}

func (c *ParserTable) Transitions(from *ItemSet) *TransitionMap {
	ts := c.transitions[from]
	if ts == nil {
		ts = new(TransitionMap)
		c.transitions[from] = ts
	}
	return ts
}

func (c *ParserTable) Actions(state *ItemSet) *ActionMap {
	am := c.actions[state]
	if am == nil {
		am = new(ActionMap)
		c.actions[state] = am
	}
	return am
}

func (t *ParserTable) Print(w io.Writer) {
	l := logger.New(w)
	for stateIndex, state := range t.States {
		l := l
		l.Logf("I%d:", stateIndex)
		l = l.WithIndent()
		l.Logf("%v", state.ToString(t.Grammar))
		l = l.WithIndent()
		actionMap := t.Actions(state)
		for _, sym := range actionMap.Terminals() {
			actions := actionMap.Get(sym)
			conflict := ""
			if actions.Len() > 1 {
				conflict = " <== CONFLICT"
			}
			for _, action := range actions.Elements() {
				l.Logf(
					"on %v %v%v",
					sym.TermName(),
					action.ToString(t.Grammar),
					conflict)
			}
		}
		transitions := t.Transitions(state)
		for _, input := range transitions.Inputs() {
			if rule, ok := input.(*Rule); ok {
				to := transitions.Get(input)
				l.Logf("on %v goto I%v", rule.Name, to.Index)
			}
		}
	}
}

func (t *ParserTable) PrintGraph(w io.Writer) {
	l := logger.New(w)
	l.Logf("digraph G {")
	li := l.WithIndent()
	for _, state := range t.States {
		li.Logf(
			"I%d [label=%q];",
			state.Index,
			fmt.Sprintf("I%d\n%v", state.Index, state.ToString(t.Grammar)))
	}
	for _, state := range t.States {
		transitions := t.Transitions(state)
		for _, input := range transitions.Inputs() {
			to := transitions.Get(input)
			li.Logf(
				"I%d -> I%d [label=%q];",
				state.Index,
				to.Index,
				input.TermName())
		}
	}
	l.Logf("}")
}
