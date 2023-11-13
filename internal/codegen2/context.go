package codegen2

import (
	gotoken "go/token"
	gotypes "go/types"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

type context struct {
	Errs          *errlogger.ErrLogger
	Fset          *gotoken.FileSet
	Dir           string
	Report        io.Writer
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
	LexerModes    map[string]*mode.Mode
}
