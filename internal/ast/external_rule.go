package ast

type ExternalRule struct {
	baseStatement
	Names []*ExternalName
}

func (r *ExternalRule) RunPass(ctx *Context, pass Pass) {
	RunPass(ctx, r.Names, pass)
}

type ExternalName struct {
	baseAST
	Name string
}

func (n *ExternalName) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		if err := validateTokenName(n.Name); err != nil {
			ctx.Errs.Errorf(n.bounds.Begin, "%s", err)
			return
		}
		if !ctx.RegisterName(n.Name, n) {
			return
		}
		ctx.Grammar.AddTerminal(n.Name)
	}
}
