package ast

type Rule struct {
	baseStatement

	Name  string
	Prods []*Prod
}

func (r *Rule) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		if !ctx.RegisterName(r.Name, r) {
			return
		}
	}

	RunPass(ctx, r.Prods, pass)
}
