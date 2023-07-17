package codegen

import (
	"fmt"
	goparser "go/parser"
	gotoken "go/token"
	gotypes "go/types"
	"os"
	"path/filepath"
	"regexp"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
	"github.com/dcaiafa/lox/internal/util/multierror"
	"golang.org/x/tools/go/packages"
)

var reduceMethodNameRegex = regexp.MustCompile(`^reduce([A-Za-z][A-Za-z0-9]*).*$`)

func (s *State) ParseGrammar() error {
	s.ProdLabels = make(map[*grammar.Prod]string)

	loxFiles, err := filepath.Glob(filepath.Join(s.LoxDir, "*.lox"))
	if err != nil {
		return err
	}

	if len(loxFiles) == 0 {
		return fmt.Errorf("%v contains no .lox files", s.LoxDir)
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
		s.addSpecToGrammar(spec, grammar)
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
	fmt.Println(parserGenGo)
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

func (s *State) addSpecToGrammar(spec *ast.Spec, g *grammar.Grammar) {
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
						if astProd.Label != nil {
							s.ProdLabels[prod] = astProd.Label.Label
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
