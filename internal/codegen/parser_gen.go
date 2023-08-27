package codegen

import (
	"bytes"
	"fmt"
	gotoken "go/token"
	gotypes "go/types"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/codegen/table"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
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

func (s {{p}}Stack[T]) Slice(n int) []T {
	return s[len(s)-n:]
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
	{{- if has_on_reduce }}
	bounds {{p}}Stack[{{p}}Bounds]
	{{- end }}
}

func (p *{{parser}}) parse(lex {{p}}Lexer, errLogger {{p}}ErrorLogger) bool {
  const accept = {{ accept }}

	p.loxParser.state.Push(0)
	tok, tokType := lex.NextToken()

	for {
		topState := p.loxParser.state.Peek(0)
		action, ok := {{p}}Find({{p}}Actions, topState, int32(tokType))
		if !ok {
			errLogger.ParserError(&{{p}}UnexpectedTokenError{Token: tok})
			return false
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p.loxParser.state.Push(action)
			p.loxParser.sym.Push(tok)
			{{- if has_on_reduce }}
			p.loxParser.bounds.Push({{p}}Bounds{Begin: tok, End: tok})
			{{- end }}
			tok, tokType = lex.NextToken()
		} else { // reduce
			prod := -action
			termCount := {{p}}TermCounts[int(prod)]
			rule := {{p}}LHS[int(prod)]
			res := p.{{p}}Act(prod)
			{{- if has_on_reduce }}
			if termCount > 0 {
				bounds := {{p}}Bounds{
					Begin: p.loxParser.bounds.Peek(int(termCount-1)).Begin,
					End: p.loxParser.bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p.loxParser.bounds.Pop(int(termCount))
				p.loxParser.bounds.Push(bounds)
			} else {
				bounds := p.loxParser.bounds.Peek(0)
				bounds.Begin = bounds.End
				p.loxParser.bounds.Push(bounds)
			}
			{{- end }}
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
	implDir       string
	grammar       *grammar.AugmentedGrammar
	errs          *errlogger.ErrLogger
	fset          *gotoken.FileSet
	parser        gotypes.Object
	token         gotypes.Object
	parserTable   *lr1.ParserTable
	reduceMethods map[string][]*ReduceMethod
	reduceTypes   map[*grammar.Rule]gotypes.Type
	reduceMap     map[*grammar.Prod]*ReduceMethod
	packageName   string
	packagePath   string
	imports       *importBuilder
	hasOnReduce   bool
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
	errs *errlogger.ErrLogger,
) *ParserGenState {
	return &ParserGenState{
		implDir: implDir,
		grammar: g,
		errs:    errs,
	}
}

func (s *ParserGenState) ConstructParseTables() {
	s.parserTable = lr1.ConstructLALR(s.grammar)
}

func (s *ParserGenState) ParseGo() {
	var err error
	s.packageName, err = computePackageName(s.implDir)
	if err != nil {
		s.errs.Error(0, err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("p", prefix)
	vars.Set("package", s.packageName)
	loxGenGo := renderTemplate(parserPlaceholderTemplate, vars)
	loxGenGoPath, err := filepath.Abs(
		filepath.Join(s.implDir, parserGenGoName))
	if err != nil {
		s.errs.Error(0, fmt.Errorf("filepath.Abs failed: %w", err))
		return
	}

	fset := gotoken.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax,
		Dir:  filepath.Clean(s.implDir),
		Fset: fset,
		Overlay: map[string][]byte{
			loxGenGoPath: []byte(loxGenGo),
		},
	}

	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		s.errs.Error(0, err)
		return
	}

	pkg := pkgs[0]
	s.packagePath = pkg.PkgPath

	if len(pkg.Errors) != 0 {
		for _, err := range pkg.Errors {
			s.errs.Error(0, err)
		}
		return
	}

	scope := pkg.Types.Scope()
	parserObj, err := getParserObj(scope)
	if err != nil {
		s.errs.Error(0, err)
		return
	}

	tokenObj := scope.Lookup("Token")

	s.fset = fset
	s.parser = parserObj
	s.token = tokenObj
	s.reduceMethods = make(map[string][]*ReduceMethod)

	parserNamed := parserObj.Type().(*gotypes.Named)
	for i := 0; i < parserNamed.NumMethods(); i++ {
		method := parserNamed.Method(i)
		if method.Name() == "onReduce" {
			s.hasOnReduce = true
			continue
		}

		ruleName := ruleNameFromMethodName(method.Name())
		if ruleName == "" {
			continue
		}

		sig := method.Type().(*gotypes.Signature)
		if sig.Results().Len() != 1 {
			s.errs.Error(0, fmt.Errorf(
				"%v: reduce method must return exactly one result",
				method.Name()))
			return
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

		s.reduceMethods[ruleName] =
			append(s.reduceMethods[ruleName], reduceMethod)
	}
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

func (s *ParserGenState) assignActions() {
	s.reduceMap = make(map[*grammar.Prod]*ReduceMethod)
	s.reduceTypes = make(map[*grammar.Rule]gotypes.Type)

	// Determine the Go type of the reduce-artifact of each rule.
	// Non-generated (user-specified) rules first.
	for ruleName, methods := range s.reduceMethods {
		var reduceMethod *ReduceMethod
		for _, method := range methods {
			if reduceMethod == nil {
				reduceMethod = method
				continue
			}
			if !gotypes.Identical(method.ReturnType, reduceMethod.ReturnType) {
				s.errs.Errorf(
					s.fset.Position(method.Method.Pos()),
					"reduce method %v returns %v but another reduce "+
						"method for the same rule %v returns %v",
					method.MethodName, method.ReturnType, ruleName,
					reduceMethod.ReturnType)
				s.errs.Infof(
					s.fset.Position(reduceMethod.Method.Pos()),
					"method %v is the method that returns %v for rule %v",
					reduceMethod.MethodName, reduceMethod.ReturnType,
					ruleName)
			}
		}
		assert(reduceMethod != nil && reduceMethod.ReturnType != nil)
		rule, ok := s.grammar.GetSymbol(ruleName).(*grammar.Rule)
		if !ok {
			s.errs.Errorf(
				s.fset.Position(reduceMethod.Method.Pos()),
				"method %v does not match a rule",
				reduceMethod.MethodName)
		}
		s.reduceTypes[rule] = reduceMethod.ReturnType
	}

	if s.errs.HasError() {
		return
	}

	// Determine the reduce-artifact's Go-type for each rule.
	// Synthetic rules first.
	changed := true
	for changed {
		changed = false
		for _, prod := range s.grammar.Prods {
			rule := s.grammar.ProdRule(prod)
			typ := s.getReduceTypeForGeneratedRule(rule, prod)
			if typ == nil {
				continue
			}
			existing := s.reduceTypes[rule]
			if existing != nil {
				assert(gotypes.Identical(existing, typ))
				continue
			}
			s.reduceTypes[rule] = typ
			changed = true
		}
	}

	// User rules next.
	for _, rule := range s.grammar.Rules {
		if rule.Generated == grammar.GeneratedSPrime {
			continue
		}
		ruleReduceType := s.reduceTypes[rule]
		if ruleReduceType == nil {
			s.errs.Errorf(
				rule.Pos,
				"rule %v does not have a reduce method",
				rule.Name)
		}
	}

	if s.errs.HasError() {
		return
	}

	// Assign each method to a production.
	for _, prod := range s.grammar.Prods {
		rule := s.grammar.ProdRule(prod)
		if rule.Generated != grammar.NotGenerated {
			// Only user (non-synthetic) rules. The action for generated rules will
			// also be generated.
			continue
		}
		method := s.findMethodForProd(prod, s.reduceMethods[rule.Name])
		if method == nil {
			s.errs.Errorf(
				prod.Pos,
				"no matching reduce method for production of rule %v",
				rule.Name)
		}
		reduceType := method.ReturnType
		if existing := s.reduceTypes[rule]; existing == nil {
			s.reduceTypes[rule] = reduceType
		} else if !gotypes.Identical(existing, reduceType) {
			// We should have already caught this by now.
			panic(fmt.Errorf(
				"conflicting reduce types for %v: %v and %v",
				rule.Name, existing, reduceType))
		}
		s.reduceMap[prod] = method
	}
}

func (s *ParserGenState) termReduceType(term *grammar.Term) gotypes.Type {
	termSym := s.grammar.TermSymbol(term)
	termReduceType := s.token.Type()
	if cRule, ok := termSym.(*grammar.Rule); ok {
		termReduceType = s.reduceTypes[cRule]
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
			termSym := s.grammar.TermSymbol(prod.Terms[i])
			termReduceType := s.token.Type()
			if cRule, ok := termSym.(*grammar.Rule); ok {
				termReduceType = s.reduceTypes[cRule]
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
		cSym := s.grammar.TermSymbol(prod.Terms[0])
		if cRule, ok := cSym.(*grammar.Rule); ok {
			return s.reduceTypes[cRule]
		} else {
			return s.token.Type()
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
		cSym := s.grammar.TermSymbol(prod.Terms[0])
		cType := s.token.Type()
		if cRule, ok := cSym.(*grammar.Rule); ok {
			cType = s.reduceTypes[cRule]
		}
		return gotypes.NewSlice(cType)

	default:
		panic("unreachable")
	}
}

func (s *ParserGenState) Generate() error {
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
	vars.Set("grammar", s.grammar)
	vars.Set("methods", s.reduceMap)
	vars.Set("not_generated", grammar.NotGenerated)
	vars.Set("generated_zero_or_one", grammar.GeneratedZeroOrOne)
	vars.Set("generated_one_or_more", grammar.GeneratedOneOrMore)
	vars.Set("go_type", s.templGoType)
	vars.Set("term_reduce_type", s.termReduceType)
	vars.Set("rule_reduce_type", s.reduceTypes)
	vars.Set("has_on_reduce", s.hasOnReduce)

	body := renderTemplate(parserTemplate, vars)

	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", s.packageName)
	s.imports.WriteTo(out)
	out.WriteString(body)
	err := os.WriteFile(filepath.Join(s.implDir, parserGenGoName), out.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParserGenState) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.grammar.Terminals {
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
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		rule := s.grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.grammar.RuleIndex(rule))
	}
	return prods
}

func (s *ParserGenState) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *ParserGenState) templTermCounts() []int32 {
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *ParserGenState) templActions() []int32 {
	actionTable := table.New()
	for _, state := range s.parserTable.States.States() {
		var row []int32
		s.parserTable.Actions.ForEachActionSet(s.grammar, state,
			func(sym grammar.Symbol, actionSet lr1.ActionSet) {
				action := actionSet.Actions()[0]
				terminal := sym.(*grammar.Terminal)
				terminalIndex := s.grammar.TerminalIndex(terminal)
				row = append(row, int32(terminalIndex))

				switch action := action.(type) {
				case lr1.ActionShift:
					row = append(row, int32(action.State.Index))
				case lr1.ActionReduce:
					row = append(row, int32(s.grammar.ProdIndex(action.Prod)*-1))
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
	for stateIndex, state := range s.parserTable.States.States() {
		var row []int32
		s.parserTable.Transitions.ForEach(
			state,
			func(sym grammar.Symbol, to *lr1.ItemSet) {
				rule, ok := sym.(*grammar.Rule)
				if !ok {
					return
				}
				ruleIndex := s.grammar.RuleIndex(rule)
				row = append(row, int32(ruleIndex), int32(to.Index))
			})
		gotoTable.AddRow(stateIndex, row)
	}
	return gotoTable.Array()
}

func ruleNameFromMethodName(method string) string {
	const prefix = "on_"
	const sep = "__"
	if !strings.HasPrefix(method, prefix) {
		return ""
	}
	rule := method[len(prefix):]
	sepIdx := strings.Index(rule, sep)
	if sepIdx != -1 {
		rule = rule[:sepIdx]
	}
	return rule
}
