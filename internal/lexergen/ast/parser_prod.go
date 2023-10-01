package ast

type ParserProd struct {
	baseAST
	Terms     []*ParserTerm
	Qualifier *ProdQualifier
}

func (p *ParserProd) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, p.Terms, pass)
}
