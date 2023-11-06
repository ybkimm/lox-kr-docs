package mode

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/nfa"
	"github.com/dcaiafa/lox/internal/util/set"
)

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

func (m *ModeBuilder) Build(errs *errlogger.ErrLogger, fset *gotoken.FileSet) *dfa.State {
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

	var visited set.Set[*dfa.State]
	m.fixState(errs, fset, d, &visited)
	if errs.HasError() {
		return nil
	}

	return d
}

func (m *ModeBuilder) fixState(
	errs *errlogger.ErrLogger,
	fset *gotoken.FileSet,
	state *dfa.State,
	visited *set.Set[*dfa.State],
) {
	if visited.Has(state) {
		return
	}
	visited.Add(state)
	state.Data = m.pickAction(errs, fset, state)
}

func (m *ModeBuilder) pickAction(
	errs *errlogger.ErrLogger,
	fset *gotoken.FileSet,
	state *dfa.State,
) *Action {
	var actions []*Action
	for _, nstate := range state.NFAStates {
		action, ok := nstate.Data.(*Action)
		if !ok {
			continue
		}
		actions = append(actions, action)
	}

	if len(actions) == 0 {
		return nil
	}

	conflict := func(a1, a2 *Action) {
		errs.Errorf(fset.Position(a1.Pos), "Conflicting lexer actions: %v", a1)
		errs.Infof(fset.Position(a2.Pos), "Conflicts with other action: %v", a2)
	}

	if len(actions) > 1 {
		if actions[0].Type != ActionEmit {
			conflict(actions[0], actions[1])
			return nil
		}
	}

	winner := actions[0]
	for i := 1; i < len(actions); i++ {
		if fset.File(actions[i].Pos) != fset.File(winner.Pos) {
			conflict(winner, actions[i])
			return nil
		}
		if actions[i].Pos > winner.Pos {
			winner = actions[i]
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
