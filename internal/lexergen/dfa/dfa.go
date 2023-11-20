package dfa

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"io"
	"slices"
	"sort"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/util/array"
	"github.com/dcaiafa/lox/internal/util/set"
	"github.com/dcaiafa/lox/internal/util/stablemap"
	"github.com/dcaiafa/lox/internal/util/stack"
)

type DFA struct {
	States []*State
}

func (d *DFA) Print(out io.Writer) {
	d.States[0].Print(out)
}

type State struct {
	ID          uint32
	Transitions stablemap.Map[any, *State]
	Accept      bool
	NFAStates   []*nfa.State
	Data        any
}

func (s *State) AddTransition(toState *State, input any) {
	s.Transitions.Put(input, toState)
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
		state.Transitions.ForEach(func(input any, destState *State) {
			edges = append(edges, Edge{
				from:  state,
				to:    destState,
				input: input,
			})
			stack.Push(destState)
		})
	}

	slices.SortStableFunc(edges, func(a, b Edge) int {
		c := cmp.Compare(a.from.ID, b.from.ID)
		if c == 0 {
			c = cmp.Compare(a.to.ID, b.to.ID)
		}
		return c
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

func NFAToDFA(n *nfa.State) *DFA {
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
				for _, toNFA := range fromNFA.Transitions.GetOrZero(input).Elements() {
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

	dfa := &DFA{
		States: assignIDs(start),
	}

	optimize(dfa)

	return dfa
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
		eTransitions := state.Transitions.GetOrZero(nfa.Epsilon)
		for _, to := range eTransitions.Elements() {
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
		nfaState.Transitions.ForEach(func(input any, _ *array.Array[*nfa.State]) {
			if input != nfa.Epsilon {
				inputs.Add(input)
			}
		})
	}
	return inputs
}

func assignIDs(s *State) []*State {
	var visited set.Set[*State]
	var pending stack.Stack[*State]
	var states []*State

	visited.Add(s)
	pending.Push(s)
	for !pending.Empty() {
		s = pending.Pop()
		s.ID = uint32(len(states))
		states = append(states, s)
		dests := s.Transitions.Values()
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
	return states
}

type partitions struct {
	stateToGroup stablemap.Map[*State, int]
	groupToState stablemap.Map[int, *set.Set[*State]]
}

func newPartitions(states []*State) *partitions {
	p := new(partitions)
	for _, s := range states {
		if s.Accept {
			p.add(s, 1)
		} else {
			p.add(s, 0)
		}
	}
	return p
}

func (p *partitions) add(s *State, group int) {
	p.stateToGroup.Put(s, group)
	states, ok := p.groupToState.Get(group)
	if !ok {
		states = new(set.Set[*State])
		p.groupToState.Put(group, states)
	}
	states.Add(s)
}

func (p *partitions) remove(s *State) {
	group, ok := p.stateToGroup.Get(s)
	assert.True(ok)
	states, ok := p.groupToState.Get(group)
	assert.True(ok)
	states.Remove(s)
}

func (p *partitions) Move(s *State, group int) {
	p.remove(s)
	p.add(s, group)
}

func (p *partitions) Count() int {
	return p.groupToState.Len()
}

func (p *partitions) GetGroup(group int) *set.Set[*State] {
	states, ok := p.groupToState.Get(group)
	assert.True(ok)
	return states
}

func (p *partitions) GetStateGroup(s *State) int {
	group, ok := p.stateToGroup.Get(s)
	assert.True(ok)
	return group
}

func optimize(d *DFA) {
	p := newPartitions(d.States)

	partitionCount := p.Count()
	for {
		for i := 0; i < partitionCount; i++ {
			subPartition(p, i)
		}
		if p.Count() == partitionCount {
			break
		}
		partitionCount = p.Count()
	}

	newStates := make([]*State, p.Count())
	for i := range newStates {
		newStates[i] = new(State)
	}

	var startGroup int
	for group := 0; group < p.Count(); group++ {
		groupStates := p.GetGroup(group)
		groupStates.ForEach(func(s *State) {
			if s.ID == 0 {
				startGroup = group
			}
			newStates[group].Accept = newStates[group].Accept || s.Accept
			newStates[group].NFAStates = append(
				newStates[group].NFAStates,
				s.NFAStates...)
		})
	}

	for _, s := range d.States {
		fromGroup := p.GetStateGroup(s)
		fromState := newStates[fromGroup]

		s.Transitions.ForEach(func(input any, ts *State) {
			toGroup := p.GetStateGroup(ts)
			toState := newStates[toGroup]
			fromState.AddTransition(toState, input)
		})
	}

	// Place group with the start event at index 0 so that it becomes the starting
	// event of the new DFA.
	newStates[0], newStates[startGroup] = newStates[startGroup], newStates[0]

	for i := range newStates {
		newStates[i].ID = uint32(i)
	}

	d.States = newStates
}

func subPartition(p *partitions, group int) {
	transitionGroup := func(s *State, input any) int {
		toState, ok := s.Transitions.Get(input)
		if !ok {
			return -1
		}
		return p.GetStateGroup(toState)
	}

	newGroup := p.Count()
	states := p.GetGroup(group)

	var inputs set.Set[any]
	states.ForEach(func(s *State) {
		s.Transitions.ForEach(func(input any, toState *State) {
			inputs.Add(input)
		})
	})

	var first *State
	var move set.Set[*State]
	states.ForEach(func(s *State) {
		if first == nil {
			first = s
			return
		}
		if first.Accept {
			assert.True(s.Accept)
			var nfaFirst, nfaS set.Set[*nfa.State]
			for _, ns := range first.NFAStates {
				if ns.Accept {
					nfaFirst.Add(ns)
				}
			}
			for _, ns := range s.NFAStates {
				if ns.Accept {
					nfaS.Add(ns)
				}
			}
			if !nfaFirst.Equal(nfaS) {
				move.Add(s)
				return
			}
		}
		inputs.ForEach(func(input any) {
			toGroupFirst := transitionGroup(first, input)
			toGroupS := transitionGroup(s, input)
			if toGroupFirst != toGroupS {
				move.Add(s)
			}
		})
	})
	move.ForEach(func(s *State) {
		p.Move(s, newGroup)
	})
}
