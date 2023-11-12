package mode

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
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
		if actionSet[i].Pos > winner.Pos {
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
