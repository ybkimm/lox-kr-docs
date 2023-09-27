package dfa

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/util/set"
	"github.com/dcaiafa/lox/internal/util/stack"
)

type State struct {
	ID          uint32
	Transitions map[any]*State
	Accept      bool
	NFAStates   []*nfa.State
	Data        any
}

func (s *State) AddTransition(toState *State, input any) {
	if s.Transitions == nil {
		s.Transitions = make(map[any]*State)
	}
	s.Transitions[input] = toState
}

func (s *State) sig() string {
	sig := make([]byte, len(s.NFAStates)*4)
	for i, nfaState := range s.NFAStates {
		binary.BigEndian.PutUint32(sig[i*4:], nfaState.ID)
	}
	return string(sig)
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
		for input, destState := range state.Transitions {
			edges = append(edges, Edge{
				from:  state,
				to:    destState,
				input: input,
			})
			stack.Push(destState)
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

	states := visited.Elements()
	sort.Slice(states, func(i, j int) bool {
		return states[i].ID < states[j].ID
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

func NFAToDFA(n *nfa.State) *State {
	// DFA states already created, indexed by their signature:
	// the unique combination of NFA states that define a DFA state.
	states := make(map[string]*State)

	// Create the start state.
	start := eClosure(set.New[*nfa.State](n))
	states[start.sig()] = start

	// Starting from the newly created 'start' state, build all DFA states, one by
	// one.
	var stack stack.Stack[*State]
	stack.Push(start)
	for !stack.Empty() {
		from := stack.Pop()

		inputs := getInputs(from.NFAStates)
		inputs.ForEach(func(input any) {
			// Create a subset of all NFA states reachable from the the NFA states
			// composing 'from', using 'input'.
			var subset set.Set[*nfa.State]
			for _, fromNFA := range from.NFAStates {
				for _, toNFA := range fromNFA.Transitions[input] {
					subset.Add(toNFA)
				}
			}

			// Expand the subset to include states reachable via ε.
			to := eClosure(subset)
			toSig := to.sig()

			// Look for an existing DFA state with the same set of NFA states
			// (uniquely represented by its signature).
			if existing := states[toSig]; existing != nil {
				// Reuse state.
				from.AddTransition(existing, input)
			} else {
				// Create a new state.
				states[toSig] = to
				from.AddTransition(to, input)

				// Add it to the stack for processing.
				stack.Push(to)
			}
		})
	}

	assignIDs(start)

	return start
}

func eClosure(nfaStates set.Set[*nfa.State]) *State {
	dfaState := new(State)

	closure := make(map[uint32]*nfa.State)
	var stack stack.Stack[*nfa.State]

	nfaStates.ForEach(func(s *nfa.State) {
		closure[s.ID] = s
		stack.Push(s)
		dfaState.Accept = dfaState.Accept || s.Accept
	})

	for !stack.Empty() {
		state := stack.Pop()
		eTransitions := state.Transitions[nfa.Epsilon]
		for _, to := range eTransitions {
			if _, ok := closure[to.ID]; !ok {
				closure[to.ID] = to
				stack.Push(to)
				dfaState.Accept = dfaState.Accept || to.Accept
			}
		}
	}

	dfaState.NFAStates = make([]*nfa.State, 0, len(closure))
	for _, nfaState := range closure {
		dfaState.NFAStates = append(dfaState.NFAStates, nfaState)
	}
	sort.Slice(dfaState.NFAStates, func(i, j int) bool {
		return dfaState.NFAStates[i].ID < dfaState.NFAStates[j].ID
	})

	return dfaState
}

func getInputs(nfaStates []*nfa.State) set.Set[any] {
	var inputs set.Set[any]
	for _, nfaState := range nfaStates {
		for input := range nfaState.Transitions {
			if input != nfa.Epsilon {
				inputs.Add(input)
			}
		}
	}
	return inputs
}

func assignIDs(s *State) {
	var visited set.Set[*State]
	var pending stack.Stack[*State]

	visited.Add(s)
	pending.Push(s)
	var nextID uint32
	for !pending.Empty() {
		s = pending.Pop()
		s.ID = nextID
		nextID++
		dests := make([]*State, 0, len(s.Transitions))
		for _, to := range s.Transitions {
			dests = append(dests, to)
		}
		sort.Slice(dests, func(i, j int) bool {
			return dests[i].ID < dests[j].ID
		})
		for _, dest := range dests {
			if !visited.Has(dest) {
				visited.Add(dest)
				pending.Push(dest)
			}
		}
	}
}
