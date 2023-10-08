package codegen2

import (
	gotoken "go/token"
	gotypes "go/types"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
	"github.com/dcaiafa/lox/internal/util/array"
)

const (
	baseGenGo       = "base.gen.go"
	parserGenGo     = "parser.gen.go"
	lexerGenGo      = "lexer.gen.go"
	parserStateName = "lox"
)

type actionMethod struct {
	Name   string
	Method *gotypes.Func
	Params []gotypes.Type
	Return gotypes.Type
}

type context struct {
	Errs          *errlogger.ErrLogger
	Fset          *gotoken.FileSet
	Dir           string
	ParserGrammar *lr2.Grammar
	GoPackageName string
	GoPackagePath string
	TokenType     gotypes.Type
	ErrorType     gotypes.Type
	ParserType    *gotypes.Named
	ActionMethods map[string]*array.Array[*actionMethod]
}
