package codegen

import (
	gotoken "go/token"
	gotypes "go/types"
	"io"

	"github.com/dcaiafa/lox/internal/base/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

type context struct {
	UseParser2    bool
	Errs          *errlogger.ErrLogger
	Fset          *gotoken.FileSet
	Dir           string
	Report        io.Writer
	ParserGrammar *lr1.Grammar
	ParserTable   *lr1.ParserTable
	GoPackageName string
	GoPackagePath string
	TokenType     gotypes.Type
	ErrorType     gotypes.Type
	ParserType    *gotypes.Named
	RuleGoTypes   map[*lr1.Rule]gotypes.Type // rule => Go-type
	ActionMethods map[*lr1.Prod]*actionMethod
	HasOnReduce   bool
	LexerModes    map[string]*mode.Mode
}
