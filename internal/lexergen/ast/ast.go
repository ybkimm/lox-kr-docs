package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/nfadfa"
)

type Pass int

const (
	CreateNames Pass = iota
	Check
	BuildNFA
)

var passes = []Pass{
	CreateNames,
	Check,
	BuildNFA,
}

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	RunPass(ctx *Context, pass Pass)
	SetBounds(b Bounds)
	Bounds() Bounds
}

type baseAST struct {
	bounds Bounds
}

func (a *baseAST) SetBounds(b Bounds) {
	a.bounds = b
}

func (a *baseAST) Bounds() Bounds {
	return a.bounds
}

type Spec struct {
	baseAST

	Units []*Unit
}

func (s *Spec) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, s.Units, pass)
}

type Unit struct {
	baseAST

	Statements []Statement
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, u.Statements, pass)
}

type Statement interface {
	AST
	isStatement()
}

type baseStatement struct {
	baseAST
}

func (s *baseStatement) isStatement() {}

type Mode struct {
	baseStatement

	Name  string
	Rules []Statement
}

func (m *Mode) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(m.Name, m) {
			return
		}
	}
	RunPass(ctx, m.Rules, pass)
}

type TokenRule struct {
	baseStatement

	Name    string
	Expr    *Expr
	Actions []Action
}

func (r *TokenRule) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}

type FragRule struct {
	baseStatement
	Expr    *Expr
	Actions []Action
}

func (r *FragRule) RunPass(ctx *Context, pass Pass) {
	r.Expr.RunPass(ctx, pass)
	RunPass(ctx, r.Actions, pass)
}

type MacroRule struct {
	baseStatement
	Name string
	Expr *Expr

	cachedNFACons *mode.NFAComposite
}

var cycleDetect = &mode.NFAComposite{}

func (r *MacroRule) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}
	r.Expr.RunPass(ctx, pass)
}

func (r *MacroRule) NFACons(ctx *Context) *mode.NFAComposite {
	if r.cachedNFACons == cycleDetect {
		ctx.Errs.Errorf(ctx.Position(r), "macro cycle detected")
	} else if r.cachedNFACons == nil {
		r.cachedNFACons = cycleDetect
		r.cachedNFACons = r.Expr.NFACons(ctx)
	}
	return r.cachedNFACons
}

type Expr struct {
	baseAST

	Factors []*Factor
}

func (e *Expr) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, e.Factors, pass)
}

func (e *Expr) NFACons(ctx *Context) *mode.NFAComposite {
	// Build the following NFACons{b,e}:
	//        ε                      ε
	//      +----> F0b --> ... F0e -----+
	//      | ε                      ε  |
	//  b --+----> F1b --> ... F1e -----+--> e
	//      | ε                      ε  |
	//      +----> F2b --> ... F2e -----+
	//
	// For all {Fnb, Fne} where Fn ∈ Factors.
	nfa := ctx.Mode().NFA
	nfaCons := &mode.NFAComposite{}
	nfaCons.B = nfa.NewState()
	nfaCons.E = nfa.NewState()
	for _, f := range e.Factors {
		fc := f.NFACons(ctx)
		nfa.AddTransition(nfaCons.B, fc.B, nfadfa.Epsilon)
		nfa.AddTransition(fc.B, nfaCons.E, nfadfa.Epsilon)
	}
	return nfaCons
}

type Factor struct {
	baseAST
	Terms []*TermCard
}

func (f *Factor) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, f.Terms, pass)
}

func (f *Factor) NFACons(ctx *Context) *mode.NFAComposite {
	// Build the following NFACons{b,e}:
	//
	//    ε                   ε                  ε       ε
	//  b ->  T0b -...-> T0e  -> T1b -...-> T1e  -> ...  -> e
	//
	// For all {Tnb, Tne} where Tn ∈ Terms.
	nfa := ctx.Mode().NFA
	nfaCons := &mode.NFAComposite{}
	nfaCons.B = nfa.NewState()
	nfaCons.E = nfa.NewState()

	tip := nfaCons.B
	for _, t := range f.Terms {
		tc := t.NFACons(ctx)
		nfa.AddTransition(tip, tc.B, nfadfa.Epsilon)
		tip = tc.E
	}

	nfa.AddTransition(tip, nfaCons.E, nfadfa.Epsilon)
	return nfaCons
}

type TermCard struct {
	baseAST

	Term Term
	Card Card
}

func (t *TermCard) RunPass(ctx *Context, pass Pass) {
	t.Term.RunPass(ctx, pass)
}

func (t *TermCard) NFACons(ctx *Context) *mode.NFAComposite {
	if t.Card != One {
		panic("not implemented")
	}
	return t.Term.NFACons(ctx)
}

type Term interface {
	AST
	NFACons(ctx *Context) *mode.NFAComposite
}

type TermRef struct {
	baseAST

	Ref string

	refMacro *MacroRule
}

func (t *TermRef) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		ast := ctx.Lookup(t.Ref)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(t), "undefined: %v", t.Ref)
			return
		}
		macro, ok := ast.(*MacroRule)
		if !ok {
			ctx.Errs.Errorf(ctx.Position(t), "invalid term: %v", t.Ref)
			return
		}
		t.refMacro = macro
	}
}

func (t *TermRef) NFACons(ctx *Context) *mode.NFAComposite {
	return t.refMacro.NFACons(ctx)
}

type TermCharClass struct {
	baseAST

	Neg            bool
	CharClassItems []*CharClassItem
}

func (t *TermCharClass) RunPass(ctx *Context, pass Pass) {}

func (t *TermCharClass) NFACons(ctx *Context) *mode.NFAComposite {
	panic("not implemented")
}

type CharClassItem struct {
	baseAST
	From rune
	To   rune
}

func (i *CharClassItem) RunPass(ctx *Context, pass Pass) {}

type Card int

const (
	One Card = iota
	ZeroOrOne
	ZeroOrMore
	OneOrMore
)

type Action interface {
	AST
	isAction()
}

type baseAction struct {
	baseAST
}

func (a *baseAction) isAction() {}

type ActionSkip struct {
	baseAction
}

func (a *ActionSkip) RunPass(ctx *Context, pass Pass) {}

type ActionPushMode struct {
	baseAction
	Mode string

	modeAST *Mode
}

func (a *ActionPushMode) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case Check:
		ast := ctx.Lookup(a.Mode)
		if ast == nil {
			ctx.Errs.Errorf(ctx.Position(a), "undefined: %v", a.Mode)
			return
		}
		modeAST, ok := ast.(*Mode)
		if !ok {
			ctx.Errs.Errorf(ctx.Position(a), "not a mode: %v", a.Mode)
			return
		}
		a.modeAST = modeAST
	}
}

type ActionPopMode struct {
	baseAction
}

func (a *ActionPopMode) RunPass(ctx *Context, pass Pass) {
}
