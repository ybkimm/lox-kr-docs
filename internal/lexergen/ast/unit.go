package ast

type Unit struct {
	baseAST

	Statements []Statement
}

func (u *Unit) RunPass(ctx *Context, pass Pass) {
	ctx.CurrentUnit.Push(u)
	defer ctx.CurrentUnit.Pop()

	RunPass(ctx, u.Statements, pass)
}
