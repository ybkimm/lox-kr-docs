package codegen2

import (
	gotoken "go/token"
	gotypes "go/types"

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

type context struct {
	Errs          *errlogger.ErrLogger
	Fset          *gotoken.FileSet
	Dir           string
	ParserGrammar *lr2.Grammar
	ParserTable   *lr2.ParserTable
	GoPackageName string
	GoPackagePath string
	TokenType     gotypes.Type
	ErrorType     gotypes.Type
	ParserType    *gotypes.Named
	RuleGoTypes   map[*lr2.Rule]gotypes.Type // rule => Go-type
	ActionMethods map[*lr2.Prod]*actionMethod
	HasOnReduce   bool
}
