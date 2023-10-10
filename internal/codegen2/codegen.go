package codegen2

import (
	gotoken "go/token"
	gotypes "go/types"
	"strings"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

const (
	baseGenGo       = "base.gen.go"
	parserGenGo     = "parser.gen.go"
	lexerGenGo      = "lexer.gen.go"
	parserStateName = "lox"
	onReduce        = "onReduce"
)

type actionMethod struct {
	Method *gotypes.Func
	Params []gotypes.Type
	Return gotypes.Type
}

func (m *actionMethod) Name() string {
	return m.Method.Name()
}

type Generated int

const (
	NotGenerated Generated = iota
	GeneratedSPrime
	GeneratedZeroOrMore
	GeneratedOneOrMore
	GeneratedZeroOrOne
	GeneratedList
)

func RuleGenerated(r *lr2.Rule) Generated {
	switch {
	case r.Name == "S'":
		return GeneratedSPrime
	case strings.HasSuffix(r.Name, "*"):
		return GeneratedZeroOrMore
	case strings.HasSuffix(r.Name, "+"):
		return GeneratedOneOrMore
	case strings.HasSuffix(r.Name, "?"):
		return GeneratedZeroOrOne
	case strings.HasPrefix(r.Name, "@list"):
		return GeneratedList
	default:
		return NotGenerated
	}
}

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
