package parsergen

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/dcaiafa/lox/internal/util/logger"
)

type item struct {
	Prod     int
	Dot      int
	Terminal int
}

func newItem(prod, dot, terminal int) item {
	return item{
		Prod:     prod,
		Dot:      dot,
		Terminal: terminal,
	}
}

func (i *item) Key() []byte {
	key := make([]byte, 0, binary.MaxVarintLen32*3)
	key = binary.AppendUvarint(key, uint64(i.Prod))
	key = binary.AppendUvarint(key, uint64(i.Dot))
	key = binary.AppendUvarint(key, uint64(i.Terminal))
	return key
}

func (i *item) ToString(g *AugmentedGrammar) string {
	var str strings.Builder
	prod := g.Prods[i.Prod]
	rule := prod.rule

	fmt.Fprintf(&str, "%v -> ", rule.Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if j == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(term.sym.SymName())
	}
	if i.Dot == len(prod.Terms) {
		str.WriteString(".")
	}
	str.WriteString(", ")
	terminal := g.Terminals[i.Terminal]
	str.WriteString(terminal.Name)

	return str.String()
}

type conflictType int

const (
	conflictNone conflictType = iota
	conflictShiftReduce
	conflictReduceReduce
)

func (c conflictType) String() string {
	switch c {
	case conflictNone:
		return "none"
	case conflictShiftReduce:
		return "shift/reduce"
	case conflictReduceReduce:
		return "reduce/reduce"
	default:
		return "???"
	}
}

type state struct {
	Items []item
	Key   string
	Index int
}

func (s *state) ToString(g *AugmentedGrammar) string {
	var str strings.Builder
	for i := range s.Items {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Items[i].ToString(g))
	}
	return str.String()
}

type stateBuilder struct {
	items map[string]item
}

func newStateBuilder() *stateBuilder {
	return &stateBuilder{
		items: make(map[string]item),
	}
}

func (b *stateBuilder) Add(item item) bool {
	itemKey := string(item.Key())
	if _, ok := b.items[itemKey]; ok {
		return false
	}
	b.items[itemKey] = item
	return true
}

func (b *stateBuilder) Build() *state {
	items := make([]item, 0, len(b.items))
	for _, item := range b.items {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		switch {
		case items[i].Prod < items[j].Prod:
			return true
		case items[i].Prod > items[j].Prod:
			return false
		case items[i].Dot < items[j].Dot:
			return true
		case items[i].Dot > items[j].Dot:
			return false
		default:
			return items[i].Terminal < items[j].Terminal
		}
	})

	keyLen := 0
	itemKeys := make([][]byte, len(items))
	for i, item := range items {
		itemKeys[i] = item.Key()
		keyLen += len(itemKeys[i])
	}

	key := make([]byte, 0, keyLen)
	for _, itemKey := range itemKeys {
		key = append(key, itemKey...)
	}

	return &state{
		Items: items,
		Key:   string(key),
	}
}

type stateSet struct {
	stateMap map[string]*state
	states   []*state
	changed  bool
}

func newStateSet() *stateSet {
	return &stateSet{
		stateMap: make(map[string]*state),
	}
}

func (c *stateSet) Changed() bool {
	return c.changed
}

func (c *stateSet) ResetChanged() {
	c.changed = false
}

func (c *stateSet) Add(s *state) *state {
	if existing, ok := c.stateMap[s.Key]; ok {
		return existing
	}
	c.changed = true
	s.Index = len(c.states)
	c.states = append(c.states, s)
	c.stateMap[s.Key] = s
	return s
}

func (c *stateSet) ForEach(fn func(s *state)) {
	for _, state := range c.states {
		fn(state)
	}
}

type transitionKey struct {
	From *state
	Sym  Symbol
}

type transitions struct {
	transitions map[transitionKey]*state
}

func newTransitions() *transitions {
	return &transitions{
		transitions: make(map[transitionKey]*state),
	}
}

func (m *transitions) Add(from *state, to *state, sym Symbol) {
	key := transitionKey{from, sym}
	if existing, ok := m.transitions[key]; ok {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	m.transitions[key] = to
}

func (m *transitions) Get(from *state, sym Symbol) *state {
	key := transitionKey{from, sym}
	toState := m.transitions[key]
	if toState == nil {
		panic("no transition")
	}
	return toState
}

func (m *transitions) ForEach(fn func(from *state, to *state, sym Symbol)) {
	keys := make([]transitionKey, 0, len(m.transitions))
	for key := range m.transitions {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		switch {
		case keys[i].From.Index < keys[j].From.Index:
			return true
		case keys[i].From.Index > keys[j].From.Index:
			return false
		default:
			return keys[i].Sym.SymName() < keys[j].Sym.SymName()
		}
	})
	for _, key := range keys {
		fn(key.From, m.transitions[key], key.Sym)
	}
}

type actionType int

const (
	actionShift actionType = iota
	actionReduce
	actionAccept
)

type action struct {
	Type   actionType
	Reduce *Rule
	Shift  *state
}

func (a action) String() string {
	switch a.Type {
	case actionShift:
		return fmt.Sprintf("shift I%v", a.Shift.Index)
	case actionReduce:
		return fmt.Sprintf("reduce %v", a.Reduce.SymName())
	case actionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
}

type actionKey struct {
	state *state
	sym   Symbol
}

type actionMap struct {
	actions map[actionKey]action
}

func newActionMap() *actionMap {
	return &actionMap{
		actions: make(map[actionKey]action),
	}
}

func (m *actionMap) Add(
	state *state,
	sym Symbol,
	action action,
	logger *logger.Logger,
) bool {
	key := actionKey{state, sym}
	action2, exists := m.actions[key]
	if exists && action == action2 {
		return true
	}

	logger.Logf(
		"state %v with %v: %v",
		state.Index,
		sym.SymName(),
		action.String())

	if exists {
		if action2.Type > action.Type {
			action, action2 = action2, action
		}
		switch {
		case action.Type == actionShift && action2.Type == actionReduce:
			logger.Logf("CONFLICT: shift/reduce")
		case action.Type == actionReduce && action2.Type == actionReduce:
			logger.Logf("CONFLICT: reduce/reduce")
		default:
			panic("invalid conflict")
		}
		return false
	}

	m.actions[key] = action
	return true
}

type parserTable struct {
	g           *AugmentedGrammar
	states      *stateSet
	transitions *transitions
	actions     *actionMap
	hasConflict bool
}

func newParserTable(g *AugmentedGrammar) *parserTable {
	return &parserTable{
		g:           g,
		states:      newStateSet(),
		transitions: newTransitions(),
		actions:     newActionMap(),
	}
}

func (t *parserTable) PrintStateGraph(w io.Writer) {
	fmt.Fprintf(w, "digraph G {\n")
	t.states.ForEach(func(s *state) {
		fmt.Fprintf(w, "  I%d [label=%q];\n",
			s.Index,
			fmt.Sprintf("I%d\n%v", s.Index, s.ToString(t.g)),
		)
	})
	t.transitions.ForEach(func(from, to *state, sym Symbol) {
		fmt.Fprintf(w, "  I%d -> I%d [label=%q];\n",
			from.Index,
			to.Index,
			sym.SymName())
	})
	fmt.Fprintf(w, "}\n")
}
