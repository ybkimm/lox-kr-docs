package nfa

import (
	"cmp"
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/lox/internal/util/set"
	"github.com/dcaiafa/lox/internal/util/stack"
)

type epsilon struct {
}

func (e epsilon) String() string {
	return "ε"
}

// Epsilon is the ε (empty string) input.
var Epsilon = epsilon{}

// State is a state in a NFA.
// NFAStates should only created using NFA.NewState().
type State struct {
	// ID of the state.
	// It is assigned by the NFA and should be read-only.
	ID uint32

	// Transitions of the state.
	// Transitions should only be added using NFA.AddTransitions.
	Transitions map[any][]*State

	// Accept indicates that the state machine should accept/recognize the input.
	// You set this yourself.
	Accept bool

	// Data is some user-data associated with this state.
	// Your set this yourself (or don't, I don't care).
	Data any
}

// AddTransition adds a transition between two states on a given input.
// The input can be any comparable data, including Epsilon.
func (s *State) AddTransition(to *State, input any) {
	if s.Transitions == nil {
		s.Transitions = make(map[any][]*State)
	}
	s.Transitions[input] = append(s.Transitions[input], to)
}

// NFA represents a Nondeterministic Finite Automaton.
// https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
// The NFA is state machine where:
// - A state is allowed to have multiple transitions for the same input.
// - The input ε (empty string) is allowed.
type StateFactory struct {
	nextID uint32
}

func NewStateFactory() *StateFactory {
	return &StateFactory{}
}

// NewState creates a new state in the NFA.
func (n *StateFactory) NewState() *State {
	s := &State{
		ID: n.nextID,
	}
	n.nextID++
	return s
}

func (n *State) Print(out io.Writer) {
	fmt.Fprintf(out, "digraph G {\n")
	fmt.Fprintf(out, "  rankdir=\"LR\";\n")

	type Edge struct {
		from  *State
		to    *State
		input any
	}

	var edges []Edge
	var visited set.Set[*State]
	var stack stack.Stack[*State]
	stack.Push(n)
	for !stack.Empty() {
		state := stack.Pop()
		if visited.Has(state) {
			continue
		}
		visited.Add(state)
		for input, destStates := range state.Transitions {
			for _, destState := range destStates {
				edges = append(edges, Edge{
					from:  state,
					to:    destState,
					input: input,
				})
				stack.Push(destState)
			}
		}
	}

	sort.SliceStable(edges, func(i, j int) bool {
		if edges[i].from.ID == edges[j].from.ID {
			return edges[i].to.ID < edges[j].to.ID
		} else {
			return edges[i].from.ID < edges[j].from.ID
		}
	})

	for _, e := range edges {
		inputStr := fmt.Sprintf("%v", e.input)
		fmt.Fprintf(out, "  %v -> %v [label=%q];\n", e.from.ID, e.to.ID, inputStr)
	}

	states := set.SortedElementsFunc(
		visited,
		func(a, b *State) int {
			return cmp.Compare(a.ID, b.ID)
		})

	for _, state := range states {
		shape := "circle"
		if state.Accept {
			shape = "doublecircle"
		}
		fmt.Fprintf(out, "  %v [label=\"%v\", shape=%q];\n", state.ID, state.ID, shape)
	}
	fmt.Fprintf(out, "}\n")
}

func getInputs(nfaStates []*State) map[any]bool {
	inputs := map[any]bool{}
	for _, nfaState := range nfaStates {
		for input := range nfaState.Transitions {
			if input != Epsilon {
				inputs[input] = true
			}
		}
	}
	return inputs
}
