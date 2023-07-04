package state

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/logger"
	"github.com/dcaiafa/lox/internal/util/set"
)

type Item struct {
	Prod     uint32
	Dot      uint32
	Terminal uint32
}

func NewItem(g *grammar.AugmentedGrammar, prod *grammar.Prod, dot int, terminal *grammar.Terminal) Item {
	return Item{
		Prod:     uint32(g.ProdIndex(prod)),
		Dot:      uint32(dot),
		Terminal: uint32(g.TerminalIndex(terminal)),
	}
}

func (i Item) IsKernel() bool {
	// Assumes that [S' -> S] is Prod 0.
	return i.Dot != 0 || i.Prod == 0
}

func (i Item) ToString(g *grammar.AugmentedGrammar) string {
	var str strings.Builder
	prod := g.Prods[i.Prod]
	rule := g.ProdRule(prod)

	fmt.Fprintf(&str, "%v -> ", rule.Name)
	for j, term := range prod.Terms {
		if j != 0 {
			str.WriteString(" ")
		}
		if uint32(j) == i.Dot {
			str.WriteString(".")
		}
		str.WriteString(g.TermSymbol(term).SymName())
	}
	if i.Dot == uint32(len(prod.Terms)) {
		str.WriteString(".")
	}
	str.WriteString(", ")
	terminal := g.Terminals[i.Terminal]
	str.WriteString(terminal.Name)

	return str.String()
}

type State struct {
	Items []Item
	Key   string
	Index int
}

func (s *State) ItemSet(g *grammar.AugmentedGrammar) *ItemSet {
	itemSet := NewItemSet(g)
	for _, item := range s.Items {
		itemSet.Add(item)
	}
	itemSet.Closure()
	return itemSet
}

func (s *State) ToString(g *grammar.AugmentedGrammar) string {
	var str strings.Builder
	for i := range s.Items {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Items[i].ToString(g))
	}
	return str.String()
}

type ItemSet struct {
	g     *grammar.AugmentedGrammar
	items map[Item]struct{}
}

func NewItemSet(g *grammar.AugmentedGrammar) *ItemSet {
	return &ItemSet{
		g:     g,
		items: make(map[Item]struct{}),
	}
}

func (b *ItemSet) Add(item Item) bool {
	if _, ok := b.items[item]; ok {
		return false
	}
	b.items[item] = struct{}{}
	return true
}

func (b *ItemSet) Closure() {
	changed := true
	for changed {
		changed = false
		// For each item [A -> α.Bβ, a]:
		for item := range b.items {
			prod := b.g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				continue
			}
			B, ok := b.g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Rule)
			if !ok {
				continue
			}
			beta := b.g.TermSymbols(prod.Terms[item.Dot+1:])
			a := b.g.Terminals[item.Terminal]
			firstSet := b.g.First(append(beta, a))
			for _, prodB := range B.Prods {
				firstSet.ForEach(func(terminal *grammar.Terminal) {
					changed = b.Add(NewItem(b.g, prodB, 0, terminal)) || changed
				})
			}
		}
	}
}

func (s *ItemSet) FollowingSymbols() []grammar.Symbol {
	symSet := new(set.Set[grammar.Symbol])
	for item := range s.items {
		prod := s.g.Prods[item.Prod]
		if item.Dot >= uint32(len(prod.Terms)) {
			continue
		}
		symSet.Add(s.g.TermSymbol(prod.Terms[item.Dot]))
	}
	syms := symSet.Elements()

	// Symbol order determines state creation order.
	// Make the analysis deterministic by sorting.
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})

	return syms
}

func (b *ItemSet) ForEach(fn func(item Item)) {
	items := make([]Item, 0, len(b.items))
	for item := range b.items {
		items = append(items, item)
	}
	sortItems(items)
	for _, item := range items {
		fn(item)
	}
}

func (b *ItemSet) State() *State {
	items := make([]Item, 0, len(b.items))
	for item := range b.items {
		if item.IsKernel() {
			items = append(items, item)
		}
	}
	sortItems(items)

	itemKey := func(i Item) []byte {
		key := make([]byte, 0, binary.MaxVarintLen32*3)
		key = binary.AppendUvarint(key, uint64(i.Prod))
		key = binary.AppendUvarint(key, uint64(i.Dot))
		key = binary.AppendUvarint(key, uint64(i.Terminal))
		return key
	}

	keyLen := 0
	itemKeys := make([][]byte, len(items))
	for i, item := range items {
		itemKeys[i] = itemKey(item)
		keyLen += len(itemKeys[i])
	}

	key := make([]byte, 0, keyLen)
	for _, itemKey := range itemKeys {
		key = append(key, itemKey...)
	}

	return &State{
		Items: items,
		Key:   string(key),
	}
}

func sortItems(items []Item) {
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
}

type StateSet struct {
	stateMap map[string]*State
	states   []*State
	changed  bool
}

func NewStateSet() *StateSet {
	return &StateSet{
		stateMap: make(map[string]*State),
	}
}

func (c *StateSet) Changed() bool {
	return c.changed
}

func (c *StateSet) ResetChanged() {
	c.changed = false
}

func (c *StateSet) Add(s *State) *State {
	if existing, ok := c.stateMap[s.Key]; ok {
		return existing
	}
	c.changed = true
	s.Index = len(c.states)
	c.states = append(c.states, s)
	c.stateMap[s.Key] = s
	return s
}

func (c *StateSet) ForEach(fn func(s *State)) {
	for _, state := range c.states {
		fn(state)
	}
}

type transitionKey struct {
	From *State
	Sym  grammar.Symbol
}

type TransitionMap struct {
	transitions map[transitionKey]*State
}

func NewTransitionMap() *TransitionMap {
	return &TransitionMap{
		transitions: make(map[transitionKey]*State),
	}
}

func (m *TransitionMap) Add(from *State, to *State, sym grammar.Symbol) {
	key := transitionKey{from, sym}
	if existing, ok := m.transitions[key]; ok {
		if existing != to {
			panic("transition redefined")
		}
		return
	}
	m.transitions[key] = to
}

func (m *TransitionMap) Get(from *State, sym grammar.Symbol) *State {
	key := transitionKey{from, sym}
	toState := m.transitions[key]
	if toState == nil {
		panic("no transition")
	}
	return toState
}

func (m *TransitionMap) ForEach(fn func(from *State, to *State, sym grammar.Symbol)) {
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

type ActionType int

const (
	ActionShift ActionType = iota
	ActionReduce
	ActionAccept
)

type Action struct {
	Type   ActionType
	Reduce *grammar.Rule
	Shift  *State
}

func (a Action) String() string {
	switch a.Type {
	case ActionShift:
		return fmt.Sprintf("shift I%v", a.Shift.Index)
	case ActionReduce:
		return fmt.Sprintf("reduce %v", a.Reduce.SymName())
	case ActionAccept:
		return "accept"
	default:
		panic("not-reached")
	}
}

type actionKey struct {
	state *State
	sym   grammar.Symbol
}

type ActionMap struct {
	actions map[actionKey]Action
}

func NewActionMap() *ActionMap {
	return &ActionMap{
		actions: make(map[actionKey]Action),
	}
}

func (m *ActionMap) Add(
	state *State,
	sym grammar.Symbol,
	action Action,
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
		case action.Type == ActionShift && action2.Type == ActionReduce:
			logger.Logf("CONFLICT: shift/reduce")
		case action.Type == ActionReduce && action2.Type == ActionReduce:
			logger.Logf("CONFLICT: reduce/reduce")
		default:
			panic("invalid conflict")
		}
		return false
	}

	m.actions[key] = action
	return true
}

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
