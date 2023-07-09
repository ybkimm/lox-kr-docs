package codegen

import (
	"fmt"
	goparser "go/parser"
	gotoken "go/token"
	gotypes "go/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
	"github.com/dcaiafa/lox/internal/util/multierror"
	"golang.org/x/tools/go/packages"
)

const loxGenGo = `
package {{package}}

type loxParser struct {}
`
const loxGenGoName = "lox.gen.go"
const loxParserTypeName = "loxParser"

type State struct {
	Grammar       *grammar.AugmentedGrammar
	Fset          *gotoken.FileSet
	Parser        gotypes.Object
	Token         gotypes.Object
	ParserTable   *lr1.ParserTable
	ReduceMethods map[string]*ReduceMethod
	ReduceMap     map[lr1.Action]*ReduceMethod
	ProdLabels    map[*grammar.Prod]string
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

func NewState() *State {
	return &State{}
}

func (s *State) ParseGrammar(dir string) error {
	loxFiles, err := filepath.Glob(filepath.Join(dir, "*.lox"))
	if err != nil {
		return err
	}

	if len(loxFiles) == 0 {
		return fmt.Errorf("%v contains no .lox files", dir)
	}

	grammar := new(grammar.Grammar)
	for _, loxFile := range loxFiles {
		loxFileData, err := os.ReadFile(loxFile)
		if err != nil {
			return err
		}
		errs := errs.New()
		spec := parser.Parse(loxFile, loxFileData, errs)
		if errs.HasErrors() {
			errs.Dump(os.Stderr)
			return fmt.Errorf("parsing lox files")
		}
		addSpecToGrammar(spec, grammar)
	}

	s.Grammar, err = grammar.ToAugmentedGrammar()
	if err != nil {
		return err
	}

	return nil
}

func (s *State) ConstructParseTables() {
	s.ParserTable = lr1.ConstructLR(s.Grammar)
}

func (s *State) ParseGo(path string) error {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	var oneSourceName string
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() &&
			filepath.Ext(dirEntry.Name()) == ".go" &&
			dirEntry.Name() != loxGenGoName {
			oneSourceName = filepath.Join(path, dirEntry.Name())
		}
	}
	if oneSourceName == "" {
		return fmt.Errorf("package contains no source files")
	}

	oneSource, err := goparser.ParseFile(gotoken.NewFileSet(), oneSourceName, nil, 0)
	if err != nil {
		return fmt.Errorf("%v: %w", oneSourceName, err)
	}

	packageName := oneSource.Name.Name
	loxGenGo := strings.Replace(loxGenGo, "{{package}}", packageName, 1)
	loxGenGoPath, err := filepath.Abs(
		filepath.Join(path, loxGenGoName))
	if err != nil {
		return fmt.Errorf("filepath.Abs failed: %w", err)
	}

	fset := gotoken.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax,
		Dir:  filepath.Clean(path),
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
	s.ReduceMethods = make(map[string]*ReduceMethod)

	parserNamed := parserObj.Type().(*gotypes.Named)
	for i := 0; i < parserNamed.NumMethods(); i++ {
		method := parserNamed.Method(i)
		if !strings.HasPrefix(method.Name(), "reduce") {
			continue
		}

		sig := method.Type().(*gotypes.Signature)
		if sig.Results().Len() != 1 {
			return fmt.Errorf(
				"%v: reduce method must return exactly one result",
				method.Name())
		}

		reduceMethod := &ReduceMethod{
			Method:     method,
			ProdName:   strings.TrimPrefix(method.Name(), "reduce"),
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

		s.ReduceMethods[reduceMethod.ProdName] = reduceMethod
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

func addSpecToGrammar(spec *ast.Spec, g *grammar.Grammar) {
	for _, section := range spec.Sections {
		switch section := section.(type) {
		case *ast.Lexer:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.CustomTokenDecl:
					for _, token := range decl.CustomTokens {
						terminal := &grammar.Terminal{
							Name: token.Name,
						}
						g.Terminals = append(g.Terminals, terminal)
					}
				default:
					panic("not-reached")
				}
			}
		case *ast.Parser:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.Rule:
					rule := &grammar.Rule{
						Name: decl.Name,
					}
					for _, astProd := range decl.Prods {
						prod := &grammar.Prod{}
						for _, astTerm := range astProd.Terms {
							term := &grammar.Term{
								Name: astTerm.Name,
							}
							switch astTerm.Qualifier {
							case ast.NoQualifier:
								term.Cardinality = grammar.One
							case ast.ZeroOrMore:
								term.Cardinality = grammar.ZeroOrMore
							case ast.OneOrMore:
								term.Cardinality = grammar.OneOrMore
							case ast.ZeroOrOne:
								term.Cardinality = grammar.ZeroOrOne
							default:
								panic("not-reached")
							}
							prod.Terms = append(prod.Terms, term)
						}
						rule.Prods = append(rule.Prods, prod)
					}
					g.Rules = append(g.Rules, rule)

				default:
					panic("not-reached")
				}
			}
		default:
			panic("not-reached")
		}
	}
}
