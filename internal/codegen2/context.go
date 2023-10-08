package codegen2

import (
	gotoken "go/token"
	gotypes "go/types"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

type context struct {
	Fset          *gotoken.FileSet
	Dir           string
	Errs          *errlogger.ErrLogger
	ParserGrammar *lr2.Grammar
	GoPackageName string
	GoPackagePath string
	TokenType     gotypes.Type
	ErrorType     gotypes.Type
	ParserType    *gotypes.Named
}
