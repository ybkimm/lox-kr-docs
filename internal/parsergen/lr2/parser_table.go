package lr2

import (
	"fmt"
	"io"

	"github.com/dcaiafa/lox/internal/util/logger"
)

type ParserTable struct {
	Grammar     *Grammar
	states      []*ItemSet             // all states
	stateMap    map[string]int         // state-key => state
	transitions map[int]*TransitionMap // state => transitions
	actions     map[int]*ActionMap     // state => actions
}

func NewParserTable(g *Grammar) *ParserTable {
	return &ParserTable{
		Grammar:     g,
		stateMap:    make(map[string]int),
		transitions: make(map[int]*TransitionMap),
		actions:     make(map[int]*ActionMap),
	}
}

func (t *ParserTable) States() []*ItemSet {
	return t.states
}

func (c *ParserTable) GetStateByKey(key string) (*ItemSet, int) {
	stateIndex, ok := c.stateMap[key]
	if !ok {
		return nil, 0
	}
	return c.states[stateIndex], stateIndex
}

func (c *ParserTable) GetStateByIndex(stateIndex int) *ItemSet {
	return c.states[stateIndex]
}

func (c *ParserTable) AddState(key string, s *ItemSet) int {
	if _, ok := c.stateMap[key]; ok {
		panic("state already exists")
	}
	c.states = append(c.states, s)
	c.stateMap[key] = len(c.states) - 1
	return len(c.states) - 1
}

func (c *ParserTable) Transitions(from int) *TransitionMap {
	ts := c.transitions[from]
	if ts == nil {
		ts = new(TransitionMap)
		c.transitions[from] = ts
	}
	return ts
}

func (c *ParserTable) Actions(state int) *ActionMap {
	am := c.actions[state]
	if am == nil {
		am = new(ActionMap)
		c.actions[state] = am
	}
	return am
}

func (t *ParserTable) Print(w io.Writer) {
	l := logger.New(w)
	for stateIndex, state := range t.States() {
		l := l
		l.Logf("I%d:", stateIndex)
		l = l.WithIndent()
		l.Logf("%v", state.ToString(t.Grammar))
		l = l.WithIndent()
		actionMap := t.Actions(stateIndex)
		for _, sym := range actionMap.Terminals() {
			actions := actionMap.Get(sym)
			conflict := ""
			if actions.Len() > 1 {
				conflict = " <== CONFLICT"
			}
			for _, action := range actions.Elements() {
				l.Logf(
					"on %v %v%v",
					t.Grammar.GetSymbolName(sym),
					action.ToString(t.Grammar),
					conflict)
			}
		}
		transitions := t.Transitions(stateIndex)
		for _, input := range transitions.Inputs() {
			if IsRule(input) {
				to := transitions.Get(input)
				l.Logf("on %v goto I%v", t.Grammar.GetSymbolName(input), to)
			}
		}
	}
}

func (t *ParserTable) PrintGraph(w io.Writer) {
	l := logger.New(w)
	l.Logf("digraph G {")
	li := l.WithIndent()
	for stateIndex, state := range t.States() {
		li.Logf(
			"I%d [label=%q];",
			stateIndex,
			fmt.Sprintf("I%d\n%v", stateIndex, state.ToString(t.Grammar)))
	}
	for stateIndex := range t.States() {
		transitions := t.Transitions(stateIndex)
		for _, input := range transitions.Inputs() {
			toIndex := transitions.Get(input)
			li.Logf(
				"I%d -> I%d [label=%q];",
				stateIndex,
				toIndex,
				t.Grammar.GetSymbolName(input))
		}
	}
	l.Logf("}")
}
