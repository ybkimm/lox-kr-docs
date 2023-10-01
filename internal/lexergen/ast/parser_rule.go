package ast

type ParserRule struct {
	baseStatement

	Name  string
	Prods []*ParserProd
}

func (r *ParserRule) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}

	RunPass(ctx, r.Prods, pass)
}
