package codegen

import (
	"bytes"
	"fmt"
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

const parserTypeName = "parser"

const parserPlaceholderTemplate = `
package {{package}}

type loxParser struct {}

func (p *loxParser) parse(l {{p}}Lexer, errLogger {{p}}ErrorLogger) bool {
	panic("not-implemented")
}
`

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

type loxParser struct {
	state {{p}}Stack[int32]
	sym   {{p}}Stack[any]
}

func (p *{{parser}}) parse(lex {{p}}Lexer, errLogger {{p}}ErrorLogger) bool {
  const accept = {{ accept }}

	p.loxParser.state.Push(0)
	tok := lex.NextToken()

	for {
		lookahead := int32(tok.Type)
		topState := p.loxParser.state.Peek(0)
		action, ok := {{p}}Find({{p}}Actions, topState, lookahead)
		if !ok {
			errLogger.Error(tok.Pos, &{{p}}UnexpectedTokenError{Token: tok})
			return false
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			tok = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := {{p}}TermCounts[int(prod)]
			rule := {{p}}LHS[int(prod)]
			res := p.{{p}}Act(prod)
			p.loxParser.state.Pop(int(termCount))
			p.loxParser.sym.Pop(int(termCount))
			topState = p.loxParser.state.Peek(0)
			nextState, _ := {{p}}Find({{p}}Goto, topState, rule)
			p.loxParser.state.Push(nextState)
			p.loxParser.sym.Push(res)
		}
	}

	return true
}

func (p *{{parser}}) {{p}}Act(prod int32) any {
	switch prod {
{{- range prod_index, prod := grammar.Prods }}
	{{- rule := grammar.ProdRule(prod) }}
	{{- if rule.Generated == not_generated }}
		{{- method := methods[prod] }}
			case {{ prod_index }}:
				return p.{{ method.MethodName}}(
				{{- range param_index, param := method.Params }}
					p.sym.Peek({{ len(method.Params) - param_index - 1 }}).({{ go_type(param.Type) }}),
				{{- end }}
		    )
	{{- else if rule.Generated == generated_one_or_more }}
  case {{ prod_index }}:  // OneOrMore
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				p.sym.Peek(0).({{ term_go_type }}),
			}
		{{- else }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[1])) }}
			return append(
				p.sym.Peek(1).([]{{term_go_type}}),
				p.sym.Peek(0).({{term_go_type}}),
			)
		{{- end }}
	{{- else if rule.Generated == generated_zero_or_one }}
  case {{ prod_index }}:  // ZeroOrOne
		{{- rule_go_type := go_type(rule_reduce_type[rule]) }}
		{{- if len(prod.Terms) == 1 }}
			return p.sym.Peek(0).({{rule_go_type}})
		{{- else }}
			{
				var zero {{rule_go_type}}
				return zero
			}
		{{- end }}
	{{- end }}
{{- end }}
	default:
		panic("unreachable")
	}
}
`

const parserGenGoName = "parser.gen.go"
const parserStateTypeName = "loxParser"

type ParserGenState struct {
	ImplDir       string
	Grammar       *grammar.AugmentedGrammar
	Fset          *gotoken.FileSet
	Parser        gotypes.Object
	Token         gotypes.Object
	ParserTable   *lr1.ParserTable
	ReduceMethods map[string][]*ReduceMethod
	ReduceTypes   map[*grammar.Rule]gotypes.Type
	ReduceMap     map[*grammar.Prod]*ReduceMethod
	packageName   string
	packagePath   string
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

func NewParserGenState(
	implDir string,
	g *grammar.AugmentedGrammar,
) *ParserGenState {
	return &ParserGenState{
		ImplDir: implDir,
		Grammar: g,
	}
}

var reduceMethodNameRegex = regexp.MustCompile(
	`^reduce([A-Za-z][A-Za-z0-9]*).*$`)

func (s *ParserGenState) ConstructParseTables() {
	s.ParserTable = lr1.ConstructLALR(s.Grammar)
}

func (s *ParserGenState) ParseGo() error {
	var err error
	s.packageName, err = computePackageName(s.ImplDir)
	if err != nil {
		return err
	}

	vars := make(jet.VarMap)
	vars.Set("p", prefix)
	vars.Set("package", s.packageName)
	loxGenGo := renderTemplate(parserPlaceholderTemplate, vars)
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
	s.packagePath = pkg.PkgPath

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
	loxStateTypeObj := scope.Lookup(parserStateTypeName)
	if loxStateTypeObj == nil {
		panic(fmt.Errorf("could not find type %q", parserStateTypeName))
	}
	loxStateType := loxStateTypeObj.Type()

	obj := scope.Lookup(parserTypeName)
	if obj == nil {
		return nil, fmt.Errorf("no type named %v", parserTypeName)
	}
	namedType, ok := obj.Type().(*gotypes.Named)
	if !ok {
		return nil, fmt.Errorf("%v is not a struct", parserTypeName)
	}
	structType, ok := namedType.Underlying().(*gotypes.Struct)
	if !ok {
		return nil, fmt.Errorf("%v is not a struct", parserTypeName)
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
		return nil, fmt.Errorf(
			"%v does not embed %v", parserTypeName, parserStateTypeName)
	}
	if obj.Type().(*gotypes.Named).TypeParams().Len() != 0 {
		return nil, fmt.Errorf(
			"%v cannot have type parameters",
			parserTypeName)
	}
	return obj, nil
}

func (s *ParserGenState) MapReduceActions() error {
	s.ReduceMap = make(map[*grammar.Prod]*ReduceMethod)
	s.ReduceTypes = make(map[*grammar.Rule]gotypes.Type)

	// Determine the Go type of the reduce-artifact of each rule.
	// Non-generated (user-specified) rules first.
	for ruleName, methods := range s.ReduceMethods {
		var reduceMethod *ReduceMethod
		for _, method := range methods {
			if reduceMethod == nil {
				reduceMethod = method
				continue
			}
			if !gotypes.Identical(method.ReturnType, reduceMethod.ReturnType) {
				return fmt.Errorf(
					"reduce methods %v and %v differ return types",
					method.MethodName, reduceMethod.MethodName)
			}
		}
		assert(reduceMethod != nil && reduceMethod.ReturnType != nil)
		rule, ok := s.Grammar.GetSymbol(ruleName).(*grammar.Rule)
		if !ok {
			return fmt.Errorf(
				"method %v has no corresponding rule",
				reduceMethod.MethodName)
		}
		s.ReduceTypes[rule] = reduceMethod.ReturnType
		fmt.Println(rule.Name, reduceMethod.ReturnType)
	}

	// Determine the Go type of the reduce-artifact of each rule.
	// Process generated rules this time.
	changed := true
	for changed {
		changed = false
		for _, prod := range s.Grammar.Prods {
			rule := s.Grammar.ProdRule(prod)
			typ := s.getReduceTypeForGeneratedRule(rule, prod)
			if typ == nil {
				continue
			}
			existing := s.ReduceTypes[rule]
			if existing != nil {
				assert(gotypes.Identical(existing, typ))
				continue
			}
			s.ReduceTypes[rule] = typ
			changed = true
		}
	}

	for _, rule := range s.Grammar.Rules {
		if rule.Generated == grammar.GeneratedSPrime {
			continue
		}
		ruleReduceType := s.ReduceTypes[rule]
		if ruleReduceType == nil {
			return fmt.Errorf(
				"rule %v does not have a ruduce method",
				rule.Name)
		}
	}

	// Assign each method to a production.
	// Only non-generated methods at this time.
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		if rule.Generated != grammar.NotGenerated {
			continue
		}
		method := s.findMethodForProd(prod, s.ReduceMethods[rule.Name])
		if method == nil {
			return fmt.Errorf(
				"there is no reduce method for %v prod #%v",
				rule.Name, prodIndex+1)
		}
		reduceType := method.ReturnType
		if existing := s.ReduceTypes[rule]; existing == nil {
			s.ReduceTypes[rule] = reduceType
		} else if !gotypes.Identical(existing, reduceType) {
			return fmt.Errorf(
				"conflicting reduce types for %v: %v and %v",
				rule.Name, existing, reduceType)
		}
		s.ReduceMap[prod] = method
	}

	for prod, method := range s.ReduceMap {
		if len(method.Params) != len(prod.Terms) {
			return fmt.Errorf(
				"%v: prod has %v terms but reduce method has %v parameters",
				method.MethodName,
				len(prod.Terms),
				len(method.Params))
		}
		for i, param := range method.Params {
			term := prod.Terms[i]
			termReduceType := s.termReduceType(term)
			if !gotypes.AssignableTo(termReduceType, param.Type) {
				return fmt.Errorf(
					"%v: param %v has type %v but term symbol %v has reduce type %v",
					method.MethodName,
					i,
					param.Type,
					term.Name,
					termReduceType.String())
			}
		}
	}

	return nil
}

func (s *ParserGenState) termReduceType(term *grammar.Term) gotypes.Type {
	termSym := s.Grammar.TermSymbol(term)
	termReduceType := s.Token.Type()
	if cRule, ok := termSym.(*grammar.Rule); ok {
		termReduceType = s.ReduceTypes[cRule]
	}
	return termReduceType
}

func (s *ParserGenState) findMethodForProd(
	prod *grammar.Prod,
	methods []*ReduceMethod,
) *ReduceMethod {

	isMatch := func(method *ReduceMethod) bool {
		if len(method.Params) != len(prod.Terms) {
			return false
		}
		for i, param := range method.Params {
			termSym := s.Grammar.TermSymbol(prod.Terms[i])
			termReduceType := s.Token.Type()
			if cRule, ok := termSym.(*grammar.Rule); ok {
				termReduceType = s.ReduceTypes[cRule]
			}
			if !gotypes.AssignableTo(termReduceType, param.Type) {
				return false
			}
		}
		return true
	}

	for _, method := range methods {
		if isMatch(method) {
			return method
		}
	}

	return nil
}

func (s *ParserGenState) getReduceTypeForGeneratedRule(
	rule *grammar.Rule,
	prod *grammar.Prod,
) gotypes.Type {
	switch rule.Generated {
	case grammar.NotGenerated,
		grammar.GeneratedSPrime:
		// S' is never reduced. Ignore.
		return nil
	case grammar.GeneratedZeroOrOne:
		// a = b c?
		//  =>
		// a = b a'
		// a' = c | e
		if prod != rule.Prods[0] {
			return nil
		}
		cSym := s.Grammar.TermSymbol(prod.Terms[0])
		if cRule, ok := cSym.(*grammar.Rule); ok {
			return s.ReduceTypes[cRule]
		} else {
			return s.Token.Type()
		}

	case grammar.GeneratedOneOrMore:
		// a = b c+
		//  =>
		// a = b a'
		// a' = a' c
		//    | c
		if prod != rule.Prods[1] {
			return nil
		}
		cSym := s.Grammar.TermSymbol(prod.Terms[0])
		cType := s.Token.Type()
		if cRule, ok := cSym.(*grammar.Rule); ok {
			cType = s.ReduceTypes[cRule]
		}
		return gotypes.NewSlice(cType)

	default:
		panic("unreachable")
	}
}

func (s *ParserGenState) Generate2() error {
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
	vars.Set("parser", parserTypeName)
	vars.Set("grammar", s.Grammar)
	vars.Set("methods", s.ReduceMap)
	vars.Set("not_generated", grammar.NotGenerated)
	vars.Set("generated_zero_or_one", grammar.GeneratedZeroOrOne)
	vars.Set("generated_one_or_more", grammar.GeneratedOneOrMore)
	vars.Set("go_type", s.templGoType)
	vars.Set("term_reduce_type", s.termReduceType)
	vars.Set("rule_reduce_type", s.ReduceTypes)

	body := renderTemplate(parserTemplate, vars)

	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", s.packageName)
	s.imports.WriteTo(out)
	out.WriteString(body)
	err := os.WriteFile(filepath.Join(s.ImplDir, parserGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParserGenState) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.Grammar.Terminals {
		terminals[i] = terminal.Name
	}
	return terminals
}

func (s *ParserGenState) templGoType(t gotypes.Type) string {
	return gotypes.TypeString(t, func(pkg *gotypes.Package) string {
		if pkg.Path() == s.packagePath {
			return ""
		}
		return s.imports.Import(pkg.Path())
	})
}

func (s *ParserGenState) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *ParserGenState) templLHS() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.Grammar.RuleIndex(rule))
	}
	return prods
}

func (s *ParserGenState) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *ParserGenState) templTermCounts() []int32 {
	prods := make([]int32, len(s.Grammar.Prods))
	for prodIndex, prod := range s.Grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *ParserGenState) templActions() []int32 {
	actionTable := table.New()
	for _, state := range s.ParserTable.States.States() {
		var row []int32
		s.ParserTable.Actions.ForEachActionSet(s.Grammar, state,
			func(sym grammar.Symbol, actionSet lr1.ActionSet) {
				action := actionSet.Actions()[0]
				terminal := sym.(*grammar.Terminal)
				terminalIndex := s.Grammar.TerminalIndex(terminal)
				row = append(row, int32(terminalIndex))

				switch action := action.(type) {
				case lr1.ActionShift:
					row = append(row, int32(action.State.Index))
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

func (s *ParserGenState) templGoto() []int32 {
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
