package codegen

import (
	gotoken "go/token"
	gotypes "go/types"
	"io"
	"strings"

	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
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

type generated string

const (
	notGenerated        generated = "not_generated"
	generatedSPrime     generated = "sprime"
	generatedZeroOrMore generated = "zero_or_more"
	generatedOneOrMore  generated = "one_or_more"
	generatedZeroOrOne  generated = "zero_or_one"
	generatedList       generated = "list"
)

func RuleGenerated(r *lr1.Rule) generated {
	switch {
	case r.Name == lr1.SPrime:
		return generatedSPrime
	case strings.HasSuffix(r.Name, "*"):
		return generatedZeroOrMore
	case strings.HasSuffix(r.Name, "+"):
		return generatedOneOrMore
	case strings.HasSuffix(r.Name, "?"):
		return generatedZeroOrOne
	case strings.HasPrefix(r.Name, "@list"):
		return generatedList
	default:
		return notGenerated
	}
}

type Config struct {
	Fset   *gotoken.FileSet
	Errs   *errlogger.ErrLogger
	Dir    string
	Report io.Writer
}

func Generate(cfg *Config) bool {
	ctx := &context{
		Fset:   cfg.Fset,
		Errs:   cfg.Errs,
		Dir:    cfg.Dir,
		Report: cfg.Report,
	}
	return ctx.ParseLox() &&
		ctx.PreParseGo() &&
		ctx.EmitBase() &&
		ctx.ParseGo() &&
		ctx.AssignActions() &&
		ctx.EmitParser() &&
		ctx.EmitLexer()
}
