package ast

type Spec struct {
	baseAST

	Units []*Unit
}

func (s *Spec) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, s.Units, pass)
}
