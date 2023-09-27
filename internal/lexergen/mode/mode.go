package mode

import "github.com/dcaiafa/lox/internal/lexergen/nfadfa"

type Range struct {
	B rune
	E rune
}

type NFAComposite struct {
	B *nfadfa.NFAState
	E *nfadfa.NFAState
}

type Mode struct {
	Name string
	NFA  *nfadfa.NFA
}

func New(name string) *Mode {
	return &Mode{
		Name: name,
		NFA:  nfadfa.NewNFA(),
	}
}
