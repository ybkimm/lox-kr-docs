package nfa

import (
	"reflect"
	"strings"
	"testing"
)

func requireEqual[T any](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("not equal:\nexpected:\n%v\nactual:\n%v", expected, actual)
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
func nfa1() *State {
	n := NewStateFactory()

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

	s0.AddTransition(s1, Epsilon)
	s1.AddTransition(s2, Epsilon)
	s2.AddTransition(s3, "a")
	s3.AddTransition(s6, Epsilon)
	s1.AddTransition(s4, Epsilon)
	s4.AddTransition(s5, "b")
	s5.AddTransition(s6, Epsilon)
	s6.AddTransition(s7, Epsilon)
	s6.AddTransition(s1, Epsilon)
	s7.AddTransition(s8, "a")
	s8.AddTransition(s9, "b")
	s9.AddTransition(s10, "b")
	s0.AddTransition(s7, Epsilon)
	s10.Accept = true

	return s0
}

func TestPrint(t *testing.T) {
	n := nfa1()

	str := &strings.Builder{}
	n.Print(str)
	requireEqual(t, strings.TrimSpace(str.String()), strings.TrimSpace(`
digraph G {
  rankdir="LR";
  0 -> 1 [label="ε"];
  0 -> 7 [label="ε"];
  1 -> 2 [label="ε"];
  1 -> 4 [label="ε"];
  2 -> 3 [label="a"];
  3 -> 6 [label="ε"];
  4 -> 5 [label="b"];
  5 -> 6 [label="ε"];
  6 -> 7 [label="ε"];
  7 -> 8 [label="a"];
  8 -> 9 [label="b"];
  9 -> 10 [label="b"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="circle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="circle"];
  7 [label="7", shape="circle"];
  8 [label="8", shape="circle"];
  9 [label="9", shape="circle"];
  10 [label="10", shape="doublecircle"];
}
`))
}
