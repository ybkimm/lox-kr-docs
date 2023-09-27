package dfa

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/util/set"
)

func requireEqual[T any](t *testing.T, a, b T) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("not equal:\na:\n%v\nb:\n%v", a, b)
	}
}

// nfa1 produces the following NFA:
//
//	                  ε
//	   +---------------------------------+
//	   |                                 |
//		 | ε     ε     a         ε         |
//		s0 -> s1 -> s2 -> s3 --------+     |
//		      ^ \                    |     |
//		      |  \   ε      b     ε  v  ε  v  a     b     b
//		      |   \----> s4 -> s5 -> s6 -> s7 -> s8 -> s9 -> ((s10))
//	        |          ε           |
//	        +----------------------+
//
// Recognizes (a|b)*abb
// Based on Dragon Book Fig. 3.27.
func nfa1() *nfa.State {
	n := nfa.NewStateFactory()

	s0 := n.NewState()
	s1 := n.NewState()
	s2 := n.NewState()
	s3 := n.NewState()
	s4 := n.NewState()
	s5 := n.NewState()
	s6 := n.NewState()
	s7 := n.NewState()
	s8 := n.NewState()
	s9 := n.NewState()
	s10 := n.NewState()

	s0.AddTransition(s1, nfa.Epsilon)
	s1.AddTransition(s2, nfa.Epsilon)
	s2.AddTransition(s3, "a")
	s3.AddTransition(s6, nfa.Epsilon)
	s1.AddTransition(s4, nfa.Epsilon)
	s4.AddTransition(s5, "b")
	s5.AddTransition(s6, nfa.Epsilon)
	s6.AddTransition(s7, nfa.Epsilon)
	s6.AddTransition(s1, nfa.Epsilon)
	s7.AddTransition(s8, "a")
	s8.AddTransition(s9, "b")
	s9.AddTransition(s10, "b")
	s0.AddTransition(s7, nfa.Epsilon)
	s10.Accept = true

	return s0
}

func TestEClosure(t *testing.T) {
	ids := func(nfaStates []*nfa.State) []uint32 {
		ids := make([]uint32, len(nfaStates))
		for i, s := range nfaStates {
			ids[i] = s.ID
		}
		return ids
	}

	c := eClosure(set.New[*nfa.State](nfa1()))
	requireEqual(t, []uint32{0, 1, 2, 4, 7}, ids(c.NFAStates))
}

func TestNFAToDFA(t *testing.T) {
	n := nfa1()
	d := NFAToDFA(n)

	str := &strings.Builder{}
	d.Print(str)

	requireEqual(t, strings.TrimSpace(str.String()), strings.TrimSpace(`
digraph G {
  rankdir="LR";
  0 -> 1 [label="b"];
  0 -> 2 [label="a"];
  1 -> 1 [label="b"];
  1 -> 2 [label="a"];
  2 -> 2 [label="a"];
  2 -> 3 [label="b"];
  3 -> 2 [label="a"];
  3 -> 4 [label="b"];
  4 -> 1 [label="b"];
  4 -> 2 [label="a"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="doublecircle"];
}
`))
}
