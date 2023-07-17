package codegen

import (
	"bytes"
	"fmt"
	goparser "go/parser"
	gotoken "go/token"
	gotypes "go/types"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/codegen/table"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
	"github.com/dcaiafa/lox/internal/util/multierror"
	"golang.org/x/tools/go/packages"
)

const accept = math.MaxInt32

const parserGenGo = `
package {{package}}

type {{p}}Lexer interface {
	NextToken() (int, Token)
}

type loxParser struct {}

func (p *loxParser) parse(l {{p}}Lexer) {}
`

const parserGenGoName = "parser.gen.go"
const loxParserTypeName = "loxParser"

type State struct {
	ImplDir       string
	Grammar       *grammar.AugmentedGrammar
	PackageName   string
	Fset          *gotoken.FileSet
	Parser        gotypes.Object
	Token         gotypes.Object
	ParserTable   *lr1.ParserTable
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

func NewState(g *grammar.AugmentedGrammar, implDir string) *State {
	return &State{
		Grammar: g,
		ImplDir: implDir,
	}
}

var reduceMethodNameRegex = regexp.MustCompile(`^reduce([A-Za-z][A-Za-z0-9]*).*$`)

func (s *State) ConstructParseTables() {
	s.ParserTable = lr1.ConstructLR(s.Grammar)
}

func (s *State) ParseGo() error {
	dirEntries, err := os.ReadDir(s.ImplDir)
	if err != nil {
		return err
	}
	var oneSourceName string
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() &&
			filepath.Ext(dirEntry.Name()) == ".go" &&
			dirEntry.Name() != parserGenGoName {
			oneSourceName = filepath.Join(s.ImplDir, dirEntry.Name())
		}
	}
	if oneSourceName == "" {
		return fmt.Errorf("package contains no source files")
	}

	oneSource, err := goparser.ParseFile(gotoken.NewFileSet(), oneSourceName, nil, 0)
	if err != nil {
		return fmt.Errorf("%v: %w", oneSourceName, err)
	}

	s.PackageName = oneSource.Name.Name

	vars := make(jet.VarMap)
	vars.Set("p", prefix)
	vars.Set("package", s.PackageName)
	loxGenGo := renderTemplate(parserGenGo, vars)
	loxGenGoPath, err := filepath.Abs(
		filepath.Join(s.ImplDir, parserGenGoName))
	if err != nil {
		return fmt.Errorf("filepath.Abs failed: %w", err)
	}

	fset := gotoken.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax,
		Dir:  filepath.Clean(s.ImplDir),
		Fset: fset,
		Overlay: map[string][]byte{
			loxGenGoPath: []byte(loxGenGo),
		},
	}

	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return err
	}

	pkg := pkgs[0]

	if len(pkg.Errors) != 0 {
		errs := multierror.MultiError{}
		for _, err := range pkg.Errors {
			errs.Add(err)
		}
		return errs
	}

	scope := pkg.Types.Scope()
	parserObj, err := getParserObj(scope)
	if err != nil {
		return err
	}

	tokenObj := scope.Lookup("Token")

	s.Fset = fset
	s.Parser = parserObj
	s.Token = tokenObj
	s.ReduceMethods = make(map[string][]*ReduceMethod)

	parserNamed := parserObj.Type().(*gotypes.Named)
	for i := 0; i < parserNamed.NumMethods(); i++ {
		method := parserNamed.Method(i)
		matches := reduceMethodNameRegex.FindStringSubmatch(method.Name())
		if matches == nil {
			continue
		}

		ruleName := matches[1]

		sig := method.Type().(*gotypes.Signature)
		if sig.Results().Len() != 1 {
			return fmt.Errorf(
				"%v: reduce method must return exactly one result",
				method.Name())
		}

		reduceMethod := &ReduceMethod{
			Method:     method,
			MethodName: method.Name(),
			ReturnType: sig.Results().At(0).Type(),
		}

		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			param := params.At(i)
			reduceParam := &ReduceParam{
				Type: param.Type(),
			}
			reduceMethod.Params = append(reduceMethod.Params, reduceParam)
		}

		s.ReduceMethods[ruleName] =
			append(s.ReduceMethods[ruleName], reduceMethod)
	}

	return nil
}

func getParserObj(scope *gotypes.Scope) (gotypes.Object, error) {
	loxStateTypeObj := scope.Lookup(loxParserTypeName)
	if loxStateTypeObj == nil {
		panic(fmt.Errorf("could not find type %q", loxParserTypeName))
	}
	loxStateType := loxStateTypeObj.Type()

	obj := scope.Lookup("Parser")
	if obj == nil {
		return nil, fmt.Errorf("no type named Parser")
	}
	namedType, ok := obj.Type().(*gotypes.Named)
	if !ok {
		return nil, fmt.Errorf("Parser is not a struct")
	}
	structType, ok := namedType.Underlying().(*gotypes.Struct)
	if !ok {
		return nil, fmt.Errorf("Parser is not a struct")
	}
	foundLoxState := false
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		if field.Embedded() && field.Type() == loxStateType {
			foundLoxState = true
			break
		}
	}
	if !foundLoxState {
		return nil, fmt.Errorf("Parser does not embed %v", loxParserTypeName)
	}
	if obj.Type().(*gotypes.Named).TypeParams().Len() != 0 {
		return nil, fmt.Errorf("Parser cannot have type parameters")
	}
	return obj, nil
}

func (s *State) Generate2() error {
	s.imports = newImportBuilder()

	vars := jet.VarMap{}
	vars.Set("accept", accept)
	vars.Set("imp", s.templImport)
	vars.Set("p", "_lx")
	vars.Set("actions", s.templActions)
	vars.Set("array", s.templArray)
	vars.Set("goto", s.templGoto)
	vars.Set("lhs", s.templLHS)
	vars.Set("term_counts", s.templTermCounts)

	body := renderTemplate(parserTemplate, vars)

	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", s.PackageName)
	s.imports.WriteTo(out)
	out.WriteString(body)
	err := os.WriteFile(filepath.Join(s.ImplDir, parserGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *State) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.Grammar.Terminals {
		terminals[i] = terminal.Name
	}
	return terminals
}

func (s *State) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *State) templLHS() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	return prods
}

func (s *State) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *State) templTermCounts() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *State) templActions() []int32 {
	actionTable := table.New()
	for _, state := range s.ParserTable.States.States() {
		var row []int32
		s.ParserTable.Actions.ForEachActionSet(s.Grammar, state,
			func(sym grammar.Symbol, actions []lr1.Action) {
				action := actions[0]
				terminal := sym.(*grammar.Terminal)
				terminalIndex := s.Grammar.TerminalIndex(terminal)
				row = append(row, int32(terminalIndex))

				switch action.Type {
				case lr1.ActionShift:
					row = append(row, int32(action.Shift.Index))
				case lr1.ActionReduce:
					row = append(row, int32(s.Grammar.ProdIndex(action.Prod)*-1))
				case lr1.ActionAccept:
					row = append(row, accept)
				default:
					panic("unreachable")
				}
			})
		actionTable.AddRow(state.Index, row)
	}
	return actionTable.Array()
}

func (s *State) templGoto() []int32 {
	gotoTable := table.New()
	for stateIndex, state := range s.ParserTable.States.States() {
		var row []int32
		s.ParserTable.Transitions.ForEach(
			state,
			func(sym grammar.Symbol, to *lr1.ItemSet) {
				rule, ok := sym.(*grammar.Rule)
				if !ok {
					return
				}
				ruleIndex := s.Grammar.RuleIndex(rule)
				row = append(row, int32(ruleIndex), int32(to.Index))
			})
		gotoTable.AddRow(stateIndex, row)
	}
	return gotoTable.Array()
}

const parserTemplate = `
var {{p}}LHS = []int32 {
	{{ lhs() | array }}
}

var {{p}}TermCounts = []int32 {
	{{ term_counts() | array }}	
}

var {{p}}Actions = []int32 {
	{{ actions() | array }}
}

var {{p}}Goto = []int32 {
	{{ goto() | array }}
}

type {{p}}Stack[T any] []T

func (s *{{p}}Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *{{p}}Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s {{p}}Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func {{p}}Find(table []int32, y, x int32) (int32, bool) {
	i := int(table[int(y)])
	count := int(table[i])
	i++
	end := i + count
	for ; i < end; i+=2 {
		if table[i] == x {
			return table[i+1], true
		}
	}
	return 0, false
}

type {{p}}Lexer interface {
	NextToken() (int, Token)
}

type loxParser struct {
	state {{p}}Stack[int32]
	sym   {{p}}Stack[any]
}

func (p *Parser) parse(lex {{p}}Lexer) {
  const accept = {{ accept }}

	p.loxParser.state.Push(0)
	lookahead, tok := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := {{p}}Find({{p}}Actions, topState, int32(lookahead))
		if !ok {
			p.onError(tok, "boom")
			return
		}
		if action == accept {
    	break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			lookahead, tok = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := {{p}}TermCounts[int(prod)]
			rule := {{p}}LHS[int(prod)]
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := {{p}}Find({{p}}Goto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(nil)
		}
	}
}

func (p *Parser) onError(tok Token, err string) {
	{{ imp("fmt") }}.Println("ERROR:", err)
	{{ imp("os") }}.Exit(1)
}
`
