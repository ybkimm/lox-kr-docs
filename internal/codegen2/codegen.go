package codegen2

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/errlogger"
)

type Config struct {
	Fset *gotoken.FileSet
	Errs *errlogger.ErrLogger
	Dir  string
}

func Generate(cfg *Config) bool {
	ctx := &context{
		Fset: cfg.Fset,
		Errs: cfg.Errs,
		Dir:  cfg.Dir,
	}

	return ctx.ParseLox() &&
		ctx.PreParseGo() &&
		ctx.EmitBase() &&
		ctx.ParseGo() &&
		ctx.AssignActions()
}
