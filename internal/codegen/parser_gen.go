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
	"sort"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/codegen/table"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
	"github.com/dcaiafa/lox/internal/util/set"
	"golang.org/x/tools/go/packages"
)

const accept = math.MaxInt32

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

type parserGen struct {
	implDir       string
	grammar       *grammar.AugmentedGrammar
	errs          *errlogger.ErrLogger
	fset          *gotoken.FileSet
	tokenType     gotypes.Type
	errorType     gotypes.Type
	parserType    *gotypes.Named
	parserTable   *lr1.ParserTable
	actionMethods map[string][]*actionMethod
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

func newParserGen(
	implDir string,
	g *grammar.AugmentedGrammar,
	parserTable *lr1.ParserTable,
	errs *errlogger.ErrLogger,
) *parserGen {
	return &parserGen{
		implDir:     implDir,
		grammar:     g,
		parserTable: parserTable,
		errs:        errs,
	}
}

func (s *parserGen) Generate() bool {
	s.parseGo()
	if s.errs.HasError() {
		return false
	}
	s.catalogActionMethods()
	if s.errs.HasError() {
		return false
	}
	s.assignActions()
	if s.errs.HasError() {
		return false
	}
	err := s.emit()
	if err != nil {
		s.errs.Errorf(gotoken.Position{}, "failed to emit parser.gen.go: %v", err)
		return false
	}
	return true
}

// parseGo parses all Go files in the project
func (s *parserGen) parseGo() {
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

	s.fset = gotoken.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax,
		Dir:  filepath.Clean(s.implDir),
		Fset: s.fset,
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

	tokenObj := scope.Lookup("Token")
	if tokenObj == nil {
		s.errs.Errorf(
			gotoken.Position{},
			"Token undefined: you must define a type named Token")
		return
	}
	s.tokenType = tokenObj.Type()

	errorObj := gotypes.Universe.Lookup("error")
	if errorObj == nil {
		panic("error is undefined")
	}

	s.errorType = errorObj.Type()

	s.parserType = s.lookupParserType(scope)
	if s.parserType == nil {
		// Error was already logged.
		return
	}
}

func (s *parserGen) catalogActionMethods() {
	s.actionMethods = make(map[string][]*actionMethod)
	for i := 0; i < s.parserType.NumMethods(); i++ {
		method := s.parserType.Method(i)
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
				s.fset.Position(method.Pos()),
				"%v: reduce method must return exactly one result",
				method.Name())
			return
		}

		action := &actionMethod{
			Method:     method,
			Name:       method.Name(),
			ReturnType: sig.Results().At(0).Type(),
		}

		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			param := params.At(i)
			actionParam := &actionParam{
				Type: param.Type(),
			}
			action.Params = append(action.Params, actionParam)
		}

		s.actionMethods[ruleName] =
			append(s.actionMethods[ruleName], action)
	}
}

// lookupParserType finds the parser struct Named type that has the form:
//
//	// Must be a top level package object (can't be embedded).
//	// The name does not matter.
//	// But it can't have type parameters (non-generic).
//	type myParser struct {
//	  // Must embed the "lox" generated type.
//	  // This contains the parser state.
//	  lox
//
//	  // Can have other fields
//	  whatever int
//	}
func (s *parserGen) lookupParserType(scope *gotypes.Scope) *gotypes.Named {
	loxObj := scope.Lookup(parserStateTypeName)
	if loxObj == nil {
		// This type is generated so this should always succeed.
		panic(fmt.Errorf("could not find type %q", parserStateTypeName))
	}
	loxType := loxObj.Type()

	// Iterate through all objects in this scope.
	var parserObj *gotypes.Named
	names := scope.Names()
	for _, name := range names {
		obj := scope.Lookup(name)

		namedType, ok := obj.Type().(*gotypes.Named)
		if !ok {
			continue
		}

		// It must be a struct.
		structType, ok := namedType.Underlying().(*gotypes.Struct)
		if !ok {
			continue
		}

		// It must embed the "lox" type.
		foundLox := false
		for i := 0; i < structType.NumFields(); i++ {
			field := structType.Field(i)
			if field.Embedded() && field.Type() == loxType {
				foundLox = true
				break
			}
		}
		if !foundLox {
			continue
		}

		// Can't have type parameters (non-generic).
		if obj.Type().(*gotypes.Named).TypeParams().Len() != 0 {
			s.errs.Errorf(
				s.fset.Position(obj.Pos()),
				"parser %v cannot have type parameters",
				namedType.Obj().Name())
			return nil
		}

		// There can be only one.
		if parserObj != nil {
			s.errs.Errorf(
				s.fset.Position(obj.Pos()),
				"there can only be one parser struct")
			s.errs.Infof(
				s.fset.Position(parserObj.Obj().Pos()),
				"here is the other one")
			return nil
		}

		parserObj = namedType
	}

	if parserObj == nil {
		s.errs.Errorf(
			gotoken.Position{},
			"Parser struct undefined")
		s.errs.Infof(
			gotoken.Position{},
			"You must define a struct for the parser that embeds \"lox\".\n"+
				"Example:\n"+
				"type myParser struct {\n"+
				"  lox  // <== must embed this\n"+
				"}")

		return nil
	}

	return parserObj
}

func (s *parserGen) assignActions() {
	s.reduceMap = make(map[*grammar.Prod]*actionMethod)
	s.reduceTypes = make(map[*grammar.Rule]gotypes.Type)

	// Determine the Go type of reduce-artifacts: user-specified rules first.
	for ruleName, methods := range s.actionMethods {
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

	// Determine the Go type of reduce-artifacts: generated rules.
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

	// Check that every rule has been covered.
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

	var unassignedMethods set.Set[*actionMethod]
	for _, methods := range s.actionMethods {
		unassignedMethods.AddSlice(methods)
	}

	// Assign each method to a production.
	for _, prod := range s.grammar.Prods {
		rule := s.grammar.ProdRule(prod)
		if rule.Generated != grammar.NotGenerated {
			// Only user (non-synthetic) rules. The action for generated rules will
			// also be generated.
			continue
		}
		method := s.findMethodForProd(prod, s.actionMethods[rule.Name])
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
		unassignedMethods.Remove(method)
	}

	if s.errs.HasError() {
		return
	}

	if unassignedMethods.Len() > 0 {
		methods := unassignedMethods.Elements()
		sort.Slice(methods, func(i, j int) bool {
			return methods[i].Name < methods[j].Name
		})
		for _, method := range methods {
			s.errs.Errorf(
				s.fset.Position(method.Method.Pos()),
				"method %v does not match a production",
				method.Name)
		}
		return
	}
}

func (s *parserGen) actionTypeForTerm(term *grammar.Term) gotypes.Type {
	termSym := s.grammar.TermSymbol(term)
	termReduceType := s.tokenType
	if cRule, ok := termSym.(*grammar.Rule); ok {
		termReduceType = s.reduceTypes[cRule]
	}
	return termReduceType
}

func (s *parserGen) findMethodForProd(
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
					termReduceType = s.errorType
				} else {
					termReduceType = s.tokenType
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

func (s *parserGen) actionTypeForGeneratedRule(
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
			return s.tokenType
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
		cType := s.tokenType
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
		cType := s.tokenType
		if cRule, ok := cSym.(*grammar.Rule); ok {
			cType = s.reduceTypes[cRule]
		}
		return gotypes.NewSlice(cType)

	default:
		panic("unreachable")
	}
}

func (s *parserGen) emit() error {
	s.imports = newImportBuilder()

	vars := jet.VarMap{}
	vars.Set("accept", accept)
	vars.Set("imp", s.templImport)
	vars.Set("actions", s.templActions)
	vars.Set("array", s.templArray)
	vars.Set("goto", s.templGoto)
	vars.Set("lhs", s.templLHS)
	vars.Set("term_counts", s.templTermCounts)
	vars.Set("parser", s.parserType.Obj().Name())
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

func (s *parserGen) terminals() map[int]string {
	terminals := make(map[int]string)
	for i, terminal := range s.grammar.Terminals {
		terminals[i] = terminal.Name
	}
	return terminals
}

func (s *parserGen) templGoType(t gotypes.Type) string {
	return gotypes.TypeString(t, func(pkg *gotypes.Package) string {
		if pkg.Path() == s.packagePath {
			return ""
		}
		return s.imports.Import(pkg.Path())
	})
}

func (s *parserGen) templImport(path string) string {
	return s.imports.Import(path)
}

func (s *parserGen) templLHS() []int32 {
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		rule := s.grammar.ProdRule(prod)
		prods[prodIndex] = int32(s.grammar.RuleIndex(rule))
	}
	return prods
}

func (s *parserGen) templArray(arr []int32) string {
	var str strings.Builder
	table.WriteArray(&str, arr)
	return str.String()
}

func (s *parserGen) templTermCounts() []int32 {
	prods := make([]int32, len(s.grammar.Prods))
	for prodIndex, prod := range s.grammar.Prods {
		prods[prodIndex] = int32(len(prod.Terms))
	}
	return prods
}

func (s *parserGen) templActions() []int32 {
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

func (s *parserGen) templGoto() []int32 {
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
