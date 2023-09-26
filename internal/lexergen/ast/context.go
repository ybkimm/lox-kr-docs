package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/errlogger"
)

type context struct {
	fset  *gotoken.FileSet
	Errs  *errlogger.ErrLogger
	Names map[string]AST
}

func newContext(fset *gotoken.FileSet, errs *errlogger.ErrLogger) *context {
	return &context{
		fset:  fset,
		Errs:  errs,
		Names: make(map[string]AST),
	}
}

func (c *context) RegisterName(name string, ast AST) bool {
	otherAST, alreadyExists := c.Names[name]
	if alreadyExists {
		c.Errs.Errorf(c.Position(ast), "%v redefined", name)
		c.Errs.Infof(c.Position(otherAST), "other %v defined here", name)
		return false
	}
	c.Names[name] = ast
	return true
}

func (c *context) Lookup(name string) AST {
	return c.Names[name]
}

func (c *context) Position(ast AST) gotoken.Position {
	return c.fset.Position(ast.Bounds().Begin)
}

func RunPass[T AST](ctx *context, asts []T, pass Pass) {
	for _, ast := range asts {
		ast.RunPass(ctx, pass)
	}
}

func Analyze(spec *Spec, fset *gotoken.FileSet, errs *errlogger.ErrLogger) {
	ctx := newContext(fset, errs)
	for _, pass := range passes {
		spec.RunPass(ctx, pass)
	}
}
