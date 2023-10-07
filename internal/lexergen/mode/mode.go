package mode

import (
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type NFAComposite struct {
	B *nfa.State
	E *nfa.State
}

type Mode struct {
	Name  string
	NFA   *nfa.StateFactory
	Rules []NFAComposite
}

func (m *Mode) AddRule(r NFAComposite) {
	m.Rules = append(m.Rules, r)
}

func (m *Mode) ComputeDFA() *dfa.State {
	// Build a single NFA from all rules:
	//          ε
	//       +----> Rules[0].B ---> ...
	//      /   ε
	// start -----> Rules[1].B ---> ...
	//      \   ε
	//       +----> Rules[N].B ---> ...
	//
	start := m.NFA.NewState()
	for _, rule := range m.Rules {
		start.AddTransition(rule.B, nfa.Epsilon)
	}

	return dfa.NFAToDFA(start)
}

func New(name string) *Mode {
	return &Mode{
		Name: name,
		NFA:  nfa.NewStateFactory(),
	}
}
