package mode

import (
	"fmt"
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
	"github.com/dcaiafa/lox/internal/util/array"
	"github.com/dcaiafa/lox/internal/util/set"
	"github.com/dcaiafa/lox/internal/util/stack"
)

type Mode struct {
	Index int
	Name  string
	DFA   *dfa.DFA
}

type NFAComposite struct {
	B *nfa.State
	E *nfa.State
}

type ModeBuilder struct {
	Name         string
	StateFactory *nfa.StateFactory
	Rules        []NFAComposite
}

func (m *ModeBuilder) AddRule(r NFAComposite) {
	m.Rules = append(m.Rules, r)
}

func (m *ModeBuilder) Build(errs *errlogger.ErrLogger, fset *gotoken.FileSet) *Mode {
	// Build a single NFA from all rules:
	//          ε
	//       +----> Rules[0].B ---> ...
	//      /   ε
	// start -----> Rules[1].B ---> ...
	//      \   ε
	//       +----> Rules[N].B ---> ...
	//
	start := m.StateFactory.NewState()
	for _, rule := range m.Rules {
		start.AddTransition(rule.B, nfa.Epsilon)
	}

	normalizeInputs(start)

	d := dfa.NFAToDFA(start)

	for _, state := range d.States {
		state.Data = m.pickAction(errs, fset, state)
	}

	if errs.HasError() {
		return nil
	}

	return &Mode{
		Name: m.Name,
		DFA:  d,
	}
}

func (m *ModeBuilder) pickAction(
	errs *errlogger.ErrLogger,
	fset *gotoken.FileSet,
	state *dfa.State,
) *Actions {
	var actionSet []*Actions
	for _, nstate := range state.NFAStates {
		actions, ok := nstate.Data.(*Actions)
		if !ok {
			continue
		}
		assert.True(len(actions.Actions) > 0)
		actionSet = append(actionSet, actions)
	}

	if len(actionSet) == 0 {
		return nil
	}

	conflict := func(a1, a2 *Actions) {
		errs.Errorf(fset.Position(a1.Pos), "Conflicting lexer actions: %v", a1)
		errs.Infof(fset.Position(a2.Pos), "Conflicts with other action: %v", a2)
	}

	winner := actionSet[0]
	for i := 1; i < len(actionSet); i++ {
		actions := actionSet[i]
		if fset.File(actions.Pos) != fset.File(winner.Pos) {
			conflict(winner, actionSet[i])
			return nil
		}
		if actionSet[i].Pos < winner.Pos {
			winner = actionSet[i]
		}
	}

	return winner
}

func New(name string) *ModeBuilder {
	return &ModeBuilder{
		Name:         name,
		StateFactory: nfa.NewStateFactory(),
	}
}

func normalizeInputs(s *nfa.State) {
	graph := make(map[rang3.Range][]*nfa.State)
	visited := set.Set[*nfa.State]{}
	pending := stack.Stack[*nfa.State]{}
	pending.Push(s)

	for !pending.Empty() {
		s = pending.Pop()
		if visited.Has(s) {
			continue
		}
		visited.Add(s)
		s.Transitions.ForEach(func(input any, toStates *array.Array[*nfa.State]) {
			for _, toState := range toStates.Elements() {
				pending.Push(toState)
			}
			inputRange, ok := input.(rang3.Range)
			if !ok {
				// Probably an ε.
				return
			}
			graph[inputRange] = append(graph[inputRange], s)
		})
	}

	for r, states := range graph {
		fmt.Println("Input", r)
		for _, state := range states {
			fmt.Println(" ", state.ID)
		}
	}

	ranges := make([]rang3.Range, 0, len(graph))
	for r := range graph {
		ranges = append(ranges, r)
	}

	rang3.Normalize(ranges, func(o, a, b, c rang3.Range) {
		states := graph[o]
		assert.True(len(states) > 0)
		for _, s := range states {
			toStates := s.Transitions.GetOrZero(o)
			s.Transitions.Remove(o)
			for _, toState := range toStates.Elements() {
				s.AddTransition(toState, a)
				s.AddTransition(toState, b)
				if c != b {
					s.AddTransition(toState, c)
				}
			}
		}
		delete(graph, o)
		graph[a] = append(graph[a], states...)
		graph[b] = append(graph[b], states...)
		if c != b {
			graph[c] = append(graph[c], states...)
		}
	})
}
