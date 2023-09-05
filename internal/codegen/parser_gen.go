package codegen

import (
	"bytes"
	"fmt"
	goformat "go/format"
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

type lox struct {}

func (p *lox) parse(l lexer) bool {
	panic("not-implemented")
}

func (p *lox) errorToken() Token {
	panic("not-implemented")
}

`

const parserTemplate = `
var _LHS = []int32 {
	{{ lhs() | array }}
}

var _TermCounts = []int32 {
	{{ term_counts() | array }}	
}

var _Actions = []int32 {
	{{ actions() | array }}
}

var _Goto = []int32 {
	{{ goto() | array }}
}

type _Stack[T any] []T

func (s *_Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func _cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

var _errorPlaceholder = {{imp("errors")}}.New("error placeholder")

func _Find(table []int32, y, x int32) (int32, bool) {
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

type lox struct {
	_lex   lexer
	_state _Stack[int32]
	_sym   _Stack[any]
	{{- if has_on_reduce }}
	_bounds _Stack[_Bounds]
	{{- end }}

	_lookahead     Token
	_lookaheadType TokenType
	_errorToken    Token
}

func (p *{{parser}}) parse(lex lexer) bool {
  const accept = {{ accept }}

	p._lex = lex

	p._state.Push(0)
	p._ReadToken()

	for {
		if p._lookaheadType == ERROR {
			_, ok := p._Recover()
			if !ok {
				return false
			}
		}
		topState := p._state.Peek(0)
		action, ok := _Find(
			_Actions, topState, int32(p._lookaheadType))
		if !ok {
			action, ok = p._Recover()
			if !ok {
				return false
			}
		}
		if action == accept {
			break
		} else if action >= 0 { // shift
			p._state.Push(action)
			p._sym.Push(p._lookahead)
			{{- if has_on_reduce }}
			p._bounds.Push(
				_Bounds{Begin: p._lookahead,
				End: p._lookahead})
			{{- end }}
			p._ReadToken()
		} else { // reduce
			prod := -action
			termCount := _TermCounts[int(prod)]
			rule := _LHS[int(prod)]
			res := p._Act(prod)
			{{- if has_on_reduce }}
			if termCount > 0 {
				bounds := _Bounds{
					Begin: p._bounds.Peek(int(termCount-1)).Begin,
					End: p._bounds.Peek(0).End,
				}
				p.onReduce(res, bounds.Begin, bounds.End)
				p._bounds.Pop(int(termCount))
				p._bounds.Push(bounds)
			} else {
				bounds := p._bounds.Peek(0)
				bounds.Begin = bounds.End
				p._bounds.Push(bounds)
			}
			{{- end }}
			p._state.Pop(int(termCount))
			p._sym.Pop(int(termCount))
			topState = p._state.Peek(0)
			nextState, _ := _Find(_Goto, topState, rule)
			p._state.Push(nextState)
			p._sym.Push(res)
		}
	}

	return true
}

func (p *{{parser}}) errorToken() Token {
	return p._errorToken
}

func (p *{{parser}}) _ReadToken() {
	p._lookahead, p._lookaheadType = p._lex.ReadToken()
}

func (p *{{parser}}) _Recover() (int32, bool) {
	p._errorToken = p._lookahead

	for {
		for p._lookaheadType == ERROR {
			p._ReadToken()
		}

		saveState := p._state
		saveSym := p._sym
		{{- if has_on_reduce }}
			saveBounds := p._bounds
		{{- end }}

		for len(p._state) > 1 {
			topState := p._state.Peek(0)
			action, ok := _Find(_Actions, topState, int32(ERROR))
			if ok {
				action2, ok := _Find(
					_Actions, action, int32(p._lookaheadType))
				if ok {
					p._state.Push(action)
					p._sym.Push(_errorPlaceholder)
					{{- if has_on_reduce }}
					  p._bounds.Push(_Bounds{})
					{{- end }}
					return action2, true
				}
			}
			p._state.Pop(1)
			p._sym.Pop(1)
			{{- if has_on_reduce }}
				p._bounds.Pop(1)
			{{- end }}
		}

		if p._lookaheadType == EOF {
			p.onError()
			return 0, false
		}

		p._ReadToken()
		p._state = saveState
		p._sym = saveSym
		{{- if has_on_reduce }}
		p._bounds = saveBounds
		{{- end }}
	}
}

func (p *{{parser}}) _Act(prod int32) any {
	switch prod {
{{- range prod_index, prod := grammar.Prods }}
	{{- rule := grammar.ProdRule(prod) }}
	{{- if rule.Generated == not_generated }}
		{{- method := methods[prod] }}
			case {{ prod_index }}:
				return p.{{ method.Name}}(
				{{- range param_index, param := method.Params }}
				  _cast[{{ go_type(param.Type) }}](p._sym.Peek({{ len(method.Params) - param_index - 1 }})),
				{{- end }}
		    )
	{{- else if rule.Generated == generated_one_or_more }}
	case {{ prod_index }}:  // OneOrMore
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			}
		{{- else }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[1])) }}
			return append(
				_cast[[]{{term_go_type}}](p._sym.Peek(1)),
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			)
		{{- end }}
	{{- else if rule.Generated == generated_list }}
	case {{ prod_index }}:  // List
		{{- if len(prod.Terms) == 1 }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[0])) }}
		  return []{{ term_go_type }}{
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			}
		{{- else }}
			{{- term_go_type := go_type(term_reduce_type(prod.Terms[2])) }}
			return append(
				_cast[[]{{ term_go_type }}](p._sym.Peek(2)),
				_cast[{{ term_go_type }}](p._sym.Peek(0)),
			)
		{{- end }}
	{{- else if rule.Generated == generated_zero_or_one }}
  case {{ prod_index }}:  // ZeroOrOne
		{{- term_go_type := go_type(rule_reduce_type[rule]) }}
		{{- if len(prod.Terms) == 1 }}
			return _cast[{{ term_go_type }}](p._sym.Peek(0))
		{{- else }}
			{
				var zero {{term_go_type}}
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
const parserStateTypeName = "lox"

type parserGenState struct {
	implDir       string
	grammar       *grammar.AugmentedGrammar
	errs          *errlogger.ErrLogger
	fset          *gotoken.FileSet
	parser        gotypes.Object
	token         gotypes.Object
	errorMark     gotypes.Object
	parserTable   *lr1.ParserTable
	reduceMethods map[string][]*actionMethod
	reduceTypes   map[*grammar.Rule]gotypes.Type
	reduceMap     map[*grammar.Prod]*actionMethod
	packageName   string
	packagePath   string
	imports       *importBuilder
	hasOnReduce   bool
}

type actionMethod struct {
	Method     *gotypes.Func
	Name       string
	Params     []*actionParam
	ReturnType gotypes.Type
}

type actionParam struct {
	Type gotypes.Type
}

func newParserGenState(
	implDir string,
	g *grammar.AugmentedGrammar,
	errs *errlogger.ErrLogger,
) *parserGenState {
	return &parserGenState{
		implDir: implDir,
		grammar: g,
		errs:    errs,
	}
}

func (s *parserGenState) ConstructParseTables() {
	s.parserTable = lr1.ConstructLALR(s.grammar)
}

func (s *parserGenState) ParseGo() {
	var err error
	s.packageName, err = computePackageName(s.implDir)
	if err != nil {
		s.errs.Errorf(gotoken.Position{}, "%v", err)
		return
	}

	vars := make(jet.VarMap)
	vars.Set("p", prefix)
	vars.Set("package", s.packageName)
	loxGenGo := renderTemplate(parserPlaceholderTemplate, vars)
	loxGenGoPath, err := filepath.Abs(
		filepath.Join(s.implDir, parserGenGoName))
	if err != nil {
		s.errs.Errorf(gotoken.Position{}, "filepath.Abs failed: %w", err)
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
		s.errs.Errorf(gotoken.Position{}, "%v", err)
		return
	}

	pkg := pkgs[0]
	s.packagePath = pkg.PkgPath

	if len(pkg.Errors) != 0 {
		for _, err := range pkg.Errors {
			s.errs.Errorf(gotoken.Position{}, "%v", err)
		}
		return
	}

	scope := pkg.Types.Scope()
	parserObj, err := getParserObj(scope)
	if err != nil {
		s.errs.Errorf(gotoken.Position{}, "%v", err)
		return
	}

	s.token = scope.Lookup("Token")
	if s.token == nil {
		s.errs.Errorf(gotoken.Position{}, "type Token is undefined")
		return
	}

	s.errorMark = gotypes.Universe.Lookup("error")
	if s.errorMark == nil {
		panic("error is undefined")
	}

	s.fset = fset
	s.parser = parserObj
	s.reduceMethods = make(map[string][]*actionMethod)

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
			s.errs.Errorf(
				gotoken.Position{},
				"%v: reduce method must return exactly one result",
				method.Name())
			return
		}

		reduceMethod := &actionMethod{
			Method:     method,
			Name:       method.Name(),
			ReturnType: sig.Results().At(0).Type(),
		}

		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			param := params.At(i)
			reduceParam := &actionParam{
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

func (s *parserGenState) assignActions() {
	s.reduceMap = make(map[*grammar.Prod]*actionMethod)
	s.reduceTypes = make(map[*grammar.Rule]gotypes.Type)

	// Determine the Go type of the reduce-artifact of each rule.
	// Non-generated (user-specified) rules first.
	for ruleName, methods := range s.reduceMethods {
		var reduceMethod *actionMethod
		for _, method := range methods {
			if reduceMethod == nil {
				reduceMethod = method
				continue
			}
			if !gotypes.Identical(method.ReturnType, reduceMethod.ReturnType) {
				s.errs.Errorf(
					s.fset.Position(method.Method.Pos()),
					"%v returns %v but %v returns %v",
					method.Name, method.ReturnType,
					reduceMethod.Name, reduceMethod.ReturnType)
				s.errs.Infof(
					s.fset.Position(method.Method.Pos()),
					"all actions for the same rule must return the same type")
				s.errs.Infof(
					s.fset.Position(reduceMethod.Method.Pos()),
					"%v is defined here",
					reduceMethod.Name)
			}
		}
		assert(reduceMethod != nil && reduceMethod.ReturnType != nil)
		rule, ok := s.grammar.GetSymbol(ruleName).(*grammar.Rule)
		if !ok {
			s.errs.Errorf(
				s.fset.Position(reduceMethod.Method.Pos()),
				"method %v does not match a rule",
				reduceMethod.Name)
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
			typ := s.actionTypeForGeneratedRule(rule, prod)
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

	// User-defined rules next.
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
				"this production does not have an action method")
			continue
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

	if s.errs.HasError() {
		return
	}
}

func (s *parserGenState) actionTypeForTerm(term *grammar.Term) gotypes.Type {
	termSym := s.grammar.TermSymbol(term)
	termReduceType := s.token.Type()
	if cRule, ok := termSym.(*grammar.Rule); ok {
		termReduceType = s.reduceTypes[cRule]
	}
	return termReduceType
}

func (s *parserGenState) findMethodForProd(
	prod *grammar.Prod,
	methods []*actionMethod,
) *actionMethod {

	isMatch := func(method *actionMethod) bool {
		if len(method.Params) != len(prod.Terms) {
			return false
		}
		for i, param := range method.Params {
			termSym := s.grammar.TermSymbol(prod.Terms[i])

			var termReduceType gotypes.Type
			switch termSym := termSym.(type) {
			case *grammar.Rule:
				termReduceType = s.reduceTypes[termSym]
			case *grammar.Terminal:
				if prod.Terms[i].Type == grammar.Error {
					termReduceType = s.errorMark.Type()
				} else {
					termReduceType = s.token.Type()
				}
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

func (s *parserGenState) actionTypeForGeneratedRule(
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

	case grammar.GeneratedList:
		// a = b @list(c, sep)
		//  =>
		// a = b a'
		// a' = a' sep c
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

func (s *parserGenState) Generate() error {
	s.imports = newImportBuilder()

	vars := jet.VarMap{}
	vars.Set("accept", accept)
	vars.Set("imp", s.templImport)
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
	vars.Set("generated_list", grammar.GeneratedList)
	vars.Set("go_type", s.templGoType)
	vars.Set("term_reduce_type", s.actionTypeForTerm)
	vars.Set("rule_reduce_type", s.reduceTypes)
	vars.Set("has_on_reduce", s.hasOnReduce)

	body := renderTemplate(parserTemplate, vars)

	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", s.packageName)
	s.imports.WriteTo(out)
	out.WriteString(body)

	outFinal, err := goformat.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format %v: %w", parserGenGoName, err)
	}

	err = os.WriteFile(filepath.Join(s.implDir, parserGenGoName), outFinal, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *parserGenState) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.grammar.Terminals {
		terminals[i] = terminal.Name
	}
	return terminals
}

func (s *parserGenState) templGoType(t gotypes.Type) string {
	return gotypes.TypeString(t, func(pkg *gotypes.Package) string {
		if pkg.Path() == s.packagePath {
			return ""
		}
		return s.imports.Import(pkg.Path())
	})
}

func (s *parserGenState) templIsErrorParam(t gotypes.Type) bool {
	return gotypes.Identical(t, s.errorMark.Type())
}

func (s *parserGenState) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *parserGenState) templLHS() []int32 {
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		rule := s.grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.grammar.RuleIndex(rule))
	}
	return prods
}

func (s *parserGenState) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *parserGenState) templTermCounts() []int32 {
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *parserGenState) templActions() []int32 {
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

func (s *parserGenState) templGoto() []int32 {
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
