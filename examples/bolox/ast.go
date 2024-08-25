package main

import (
	"fmt"
	gotoken "go/token"
	"reflect"
	"strings"
)

// Bounds encapsulates the lexical beginning and end of an AST.
type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

// AST is the interface for ast nodes.
type AST interface {
	// Bounds returns the AST bounds.
	Bounds() Bounds

	// SetBounds set the AST bounds.
	SetBounds(b Bounds)

	// Discard returns whether a node should be discarded by the parser
	// when the *! cardinality is used.
	Discard() bool
}

// BaseAST can be embedded to provide a basic implementation of the AST
// interface.
type BaseAST struct {
	bounds Bounds
}

// Bounds returns the AST bounds.
func (b *BaseAST) Bounds() Bounds {
	return b.bounds
}

// SetBounds set the AST bounds.
func (b *BaseAST) SetBounds(v Bounds) {
	b.bounds = v
}

// Discard returns whether a node should be discarded by the parser
// when the *! cardinality is used. By default, no AST node should be discarded.
func (b *BaseAST) Discard() bool { return false }

// errorWithPos augments the provided error with the AST's location.
// N.B. this is not the best way to provide a real time error stack.
func errorWithPos(ctx *Context, ast AST, err error) error {
	return fmt.Errorf("%w\nfrom: %v", err, ctx.FileSet.Position(ast.Bounds().Begin))
}

// Control is returned by a statement to control program flow.
type Control int

const (
	Step     Control = iota // Move to the next step.
	Continue                // For/while continue.
)

// Statement is an AST node that can executed.
type Statement interface {
	AST
	Run(ctx *Context) (Control, error)
}

// Program is the top level AST for a program.
type Program struct {
	BaseAST
	Block *Block
}

func (p *Program) Run(ctx *Context) error {
	step, err := p.Block.Run(ctx)
	if err != nil {
		return err
	}

	switch step {
	case Continue:
		return fmt.Errorf("continue used outside of while")
	case Step:
	default:
		panic("unreachable")
	}

	return nil
}

// Block is a block of statements.
type Block struct {
	BaseAST
	Statements []Statement
}

func (b *Block) Run(ctx *Context) (Control, error) {
	for _, stmt := range b.Statements {
		ctrl, err := stmt.Run(ctx)
		if err != nil {
			return 0, err
		}
		if ctrl != Step {
			return ctrl, nil
		}
	}
	return Step, nil
}

// FuncCallStatement is a statement composed of a single function call.
type FuncCallStatement struct {
	BaseAST
	FuncCall *FuncCall
}

func (s *FuncCallStatement) Run(ctx *Context) (Control, error) {
	_, err := s.FuncCall.Eval(ctx)
	if err != nil {
		return 0, errorWithPos(ctx, s.FuncCall, err)
	}
	return Step, err
}

// VarAssign is a variable assignment.
type VarAssign struct {
	BaseAST

	VarName string
	Value   Expr
}

func (a *VarAssign) Run(ctx *Context) (Control, error) {
	v, err := a.Value.Eval(ctx)
	if err != nil {
		return Step, err
	}
	ctx.SetGlobal(a.VarName, v)
	return Step, nil
}

// Expr is an AST that can be evaluated for a value.
type Expr interface {
	AST
	Eval(ctx *Context) (any, error)
}

// FuncCall is an expression consisting of a function call.
type FuncCall struct {
	BaseAST

	FuncName string
	Args     []Expr
}

func (c *FuncCall) Eval(ctx *Context) (any, error) {
	var err error
	vals := make([]any, len(c.Args))
	for i, arg := range c.Args {
		vals[i], err = arg.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ctx.Call(c.FuncName, vals)
}

// VarRef is an expression for dereferencing a variable.
type VarRef struct {
	BaseAST

	VarName string
}

func (r *VarRef) Eval(ctx *Context) (any, error) {
	v, ok := ctx.GetGlobal(r.VarName)
	if !ok {
		return nil, errorWithPos(ctx, r, fmt.Errorf("undefined: %v", r.VarName))
	}
	return v, nil
}

// Literal is an expression for a literal: int, bool, string, etc.
type Literal struct {
	BaseAST

	Val any
}

func (l *Literal) Eval(ctx *Context) (any, error) {
	return l.Val, nil
}

// Op is an operator.
type Op int

const (
	OpPlus  Op = iota // +
	OpMinus           // -
	OpTimes           // *
	OpDiv             // /
	OpLT              // <
	OpLE              // <=
	OpGT              // >
	OpGE              // >=
	OpEq              // ==
	OpNE              // !=
	OpAnd             // and
	OpOr              // or
)

// BinaryExpr is an expression consisting of <expr> <operator> <expr>.
type BinaryExpr struct {
	BaseAST

	Left  Expr
	Right Expr
	Op    Op
}

func (e *BinaryExpr) Eval(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()
	switch e.Op {
	case OpAnd:
		return e.evalAnd(ctx)
	case OpOr:
		return e.evalOr(ctx)
	}

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}

	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	switch e.Op {
	case OpEq:
		return va == vb, nil
	case OpNE:
		return va != vb, nil
	}

	if a, b, ok := castBinaryExpr[int](va, vb); ok {
		switch e.Op {
		case OpPlus:
			return a + b, nil
		case OpMinus:
			return a - b, nil
		case OpTimes:
			return a * b, nil
		case OpDiv:
			if b == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			return a / b, nil
		case OpLT:
			return a < b, nil
		case OpLE:
			return a <= b, nil
		case OpGT:
			return a > b, nil
		case OpGE:
			return a >= b, nil
		default:
			panic("unreachable")
		}
	} else if a, b, ok := castBinaryExpr[string](va, vb); ok {
		switch e.Op {
		case OpPlus:
			return a + b, nil
		case OpLT:
			return a < b, nil
		case OpLE:
			return a <= b, nil
		case OpGT:
			return a > b, nil
		case OpGE:
			return a >= b, nil
		default:
			return nil, fmt.Errorf("operation not supported by string")
		}
	} else {
		return nil, fmt.Errorf(
			"operation not supported between %v and %v",
			reflect.TypeOf(va), reflect.TypeOf(vb))
	}
}

func (e *BinaryExpr) evalAnd(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	ba, ok := va.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(va))
	}
	if !ba {
		return false, nil
	}
	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	bb, ok := vb.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(vb))
	}
	return ba && bb, nil
}

func (e *BinaryExpr) evalOr(ctx *Context) (res any, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, e, err)
		}
	}()

	va, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	ba, ok := va.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(va))
	}
	if ba {
		return true, nil
	}
	vb, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	bb, ok := vb.(bool)
	if !ok {
		return nil, fmt.Errorf(
			"operation not supported for %v",
			reflect.TypeOf(vb))
	}
	return ba || bb, nil
}

func castBinaryExpr[T any](a, b any) (ca, cb T, ok bool) {
	ca, ok = a.(T)
	if !ok {
		return ca, cb, false
	}
	cb, ok = b.(T)
	if !ok {
		return ca, cb, false
	}
	return ca, cb, true
}

// While is a while loop.
//
//	while <expr> {
//	  <statements>
//	}
type While struct {
	BaseAST

	Pred  Expr
	Block *Block
}

func (w *While) Run(ctx *Context) (Control, error) {
	for {
		pred, err := evalPredicate(ctx, w.Pred)
		if err != nil {
			return 0, err
		}
		if !pred {
			break
		}
		ctrl, err := w.Block.Run(ctx)
		if err != nil {
			return 0, err
		}
		switch ctrl {
		case Continue:
			continue
		case Step:
		default:
			panic("unreachable")
		}
	}
	return Step, nil
}

// ContinueStatement is the 'continue' statement.
type ContinueStatement struct {
	BaseAST
}

func (s *ContinueStatement) Run(ctx *Context) (Control, error) {
	return Continue, nil
}

// IfStatement is the 'if' statement:
//
//	if <expr> {
//	  <statements>
//	} elif <expr> {
//	  <statements>
//	} else {
//	  <statements>
//	}
type IfStatement struct {
	BaseAST

	Pred  Expr
	Block *Block
	Elifs []*Elif
	Else  *Else
}

func (s *IfStatement) Run(ctx *Context) (Control, error) {
	v, err := evalPredicate(ctx, s.Pred)
	if err != nil {
		return 0, err
	}
	if v {
		return s.Block.Run(ctx)
	}

	for _, elif := range s.Elifs {
		v, err := evalPredicate(ctx, elif.Pred)
		if err != nil {
			return 0, err
		}
		if v {
			return elif.Block.Run(ctx)
		}
	}

	if s.Else != nil {
		return s.Else.Block.Run(ctx)
	}

	return Step, nil
}

// Elif is an 'elif' part of a 'IfStatement'
type Elif struct {
	BaseAST

	Pred  Expr
	Block *Block
}

// Else is a 'else' part of a 'IfStatement'
type Else struct {
	BaseAST

	Block *Block
}

// Noop is a no-op statement which will be discarded by the parser
// when the *! cardinality is used.
type Noop struct {
	BaseAST
}

func (n *Noop) Run(ctx *Context) (Control, error) { return Step, nil }
func (n *Noop) Discard() bool                     { return true }

func evalPredicate(ctx *Context, p Expr) (res bool, err error) {
	defer func() {
		if err != nil {
			err = errorWithPos(ctx, p, err)
		}
	}()
	v, err := p.Eval(ctx)
	if err != nil {
		return false, err
	}

	vb, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf(
			"predicate must be bool, not %T", v)
	}

	return vb, nil
}

// String is an expression for a string literal with string interpolation.
type String struct {
	BaseAST

	Parts []Expr
}

func (s *String) Eval(ctx *Context) (any, error) {
	var res strings.Builder
	for _, part := range s.Parts {
		v, err := part.Eval(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(&res, "%v", v)
	}
	return res.String(), nil
}

// StringCharSeq is a character sequence (an actual string) that is part of a
// String.
type StringCharSeq struct {
	BaseAST
	Seq string
}

func (s *StringCharSeq) Eval(ctx *Context) (any, error) {
	return s.Seq, nil
}
