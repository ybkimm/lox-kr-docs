package codegen

import (
	"bytes"
	gotoken "go/token"
	gotypes "go/types"
	"math"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

const accept = math.MaxInt32

const loxGenGo = `
package {{package}}

type loxParser struct {}
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
	ReduceMethods map[string]*ReduceMethod
	ReduceTypes   map[*grammar.Rule]gotypes.Type
	ReduceMap     map[*grammar.Prod]*ReduceMethod
	imports       *importBuilder
	body          bytes.Buffer
}

type ReduceMethod struct {
	Method     *gotypes.Func
	ProdName   string
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
