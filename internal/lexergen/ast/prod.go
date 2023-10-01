package ast

type Prod struct {
	baseAST
	Terms     []*Term
	Qualifier *ProdQualifier
}

func (p *Prod) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, p.Terms, pass)
}
