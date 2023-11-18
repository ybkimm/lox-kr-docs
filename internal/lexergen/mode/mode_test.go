package mode

import (
	"math"
	"strings"
	"testing"

	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
	"github.com/dcaiafa/lox/internal/testutil"
)

func TestNormalizeInputs(t *testing.T) {
	n := nfa.NewStateFactory()
	s := make([]*nfa.State, 12)
	for i := range s {
		s[i] = n.NewState()
	}

	s[0].AddTransition(s[1], nfa.Epsilon)
	s[1].AddTransition(s[2], rang3.Range{B: '5', E: '5'})
	s[2].AddTransition(s[7], nfa.Epsilon)
	s[0].AddTransition(s[3], nfa.Epsilon)
	s[3].AddTransition(s[4], rang3.Range{B: '0', E: '9'})
	s[4].AddTransition(s[7], nfa.Epsilon)
	s[0].AddTransition(s[5], nfa.Epsilon)
	s[5].AddTransition(s[6], rang3.Range{B: 0, E: rune(math.MaxInt32)})
	s[6].AddTransition(s[7], nfa.Epsilon)

	s[2].Accept = true
	s[4].Accept = true
	s[6].Accept = true

	normalizeInputs(s[0])

	var str strings.Builder
	s[0].Print(&str)

	testutil.RequireEqualStr(t, str.String(), `
digraph G {
  rankdir="LR";
  0 -> 1 [label="ε"];
  0 -> 3 [label="ε"];
  0 -> 5 [label="ε"];
  1 -> 2 [label="5"];
  2 -> 7 [label="ε"];
  3 -> 4 [label="0-4"];
  3 -> 4 [label="5"];
  3 -> 4 [label="6-9"];
  4 -> 7 [label="ε"];
  5 -> 6 [label="\\u0000-/"];
  5 -> 6 [label=":-\\u7fffffff"];
  5 -> 6 [label="0-4"];
  5 -> 6 [label="5"];
  5 -> 6 [label="6-9"];
  6 -> 7 [label="ε"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="doublecircle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="doublecircle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="doublecircle"];
  7 [label="7", shape="circle"];
}
`)
}
