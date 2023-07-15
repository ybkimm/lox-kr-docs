package codegen

import (
	gotoken "go/token"
	gotypes "go/types"
	"math"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

const accept = math.MaxInt32

const loxGenGo = `
package {{package}}

type _lxLexer interface {
	NextToken() (int, Token)
}

type loxParser struct {}

func (p *loxParser) parse(l _lxLexer) {}

`
const loxGenGoName = "lox.gen.go"
const loxParserTypeName = "loxParser"

type State struct {
	LoxDir        string
	ImplDir       string
	Grammar       *grammar.AugmentedGrammar
	PackageName   string
	Fset          *gotoken.FileSet
	Parser        gotypes.Object
	Token         gotypes.Object
	ParserTable   *lr1.ParserTable
	ProdLabels    map[*grammar.Prod]string
	ReduceMethods map[string][]*ReduceMethod
	ReduceTypes   map[*grammar.Rule]gotypes.Type
	ReduceMap     map[*grammar.Prod]*ReduceMethod
	imports       *importBuilder
}

type ReduceMethod struct {
	Method     *gotypes.Func
	MethodName string
	Params     []*ReduceParam
	ReturnType gotypes.Type
}

type ReduceParam struct {
	Type gotypes.Type
}

func NewState(loxDir, implDir string) *State {
	return &State{
		LoxDir:  loxDir,
		ImplDir: implDir,
	}
}
