package mode

import (
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type NFAComposite struct {
	B *nfa.State
	E *nfa.State
}

type Mode struct {
	Name string
	NFA  *nfa.StateFactory
}

func New(name string) *Mode {
	return &Mode{
		Name: name,
		NFA:  nfa.NewStateFactory(),
	}
}
