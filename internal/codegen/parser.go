package codegen

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

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
	Fset          *token.FileSet
	ParserDecl    types.Object
	ReduceMethods map[string]*ReduceMethod
}

type ReduceMethod struct {
	Method     *types.Func
	ProdName   string
	MethodName string
	Params     []*ReduceParam
	ReturnType types.Type
}

type ReduceParam struct {
	TermIndex int
	Type      types.Type
}

var reduceParamRegex = regexp.MustCompile(`^.+?(\d+)$`)

func Parse(path string) (*State, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("package contains no source files")
	}

	oneSource, err := parser.ParseFile(token.NewFileSet(), oneSourceName, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", oneSourceName, err)
	}

	packageName := oneSource.Name.Name
	loxGenGo := strings.Replace(loxGenGo, "{{package}}", packageName, 1)
	loxGenGoPath, err := filepath.Abs(
		filepath.Join(path, loxGenGoName))
	if err != nil {
		return nil, fmt.Errorf("filepath.Abs failed: %w", err)
	}

	fset := token.NewFileSet()
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
		return nil, err
	}

	pkg := pkgs[0]

	if len(pkg.Errors) != 0 {
		errs := multierror.MultiError{}
		for _, err := range pkg.Errors {
			errs.Add(err)
		}
		return nil, errs
	}

	scope := pkg.Types.Scope()

	loxStateTypeObj := scope.Lookup(loxParserTypeName)
	if loxStateTypeObj == nil {
		panic(fmt.Errorf("could not find type %q", loxParserTypeName))
	}
	loxStateType := loxStateTypeObj.Type()

	var parserObj types.Object
	for _, typeName := range scope.Names() {
		obj := scope.Lookup(typeName)
		namedType, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}
		structType, ok := namedType.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		for i := 0; i < structType.NumFields(); i++ {
			field := structType.Field(i)
			if field.Embedded() && field.Type() == loxStateType {
				if parserObj != nil {
					return nil, fmt.Errorf("multiple Parser objects")
				}
				parserObj = obj
				break
			}
		}
	}
	if parserObj == nil {
		return nil, fmt.Errorf("no parser found")
	}

	if parserObj.Type().(*types.Named).TypeParams().Len() != 0 {
		return nil, fmt.Errorf("%v: cannot have type parameters", parserObj.Name())
	}

	state := &State{
		Fset:          fset,
		ParserDecl:    parserObj,
		ReduceMethods: make(map[string]*ReduceMethod),
	}

	parserNamed := parserObj.Type().(*types.Named)
	for i := 0; i < parserNamed.NumMethods(); i++ {
		method := parserNamed.Method(i)
		if !strings.HasPrefix(method.Name(), "reduce") {
			continue
		}

		sig := method.Type().(*types.Signature)
		if sig.Results().Len() != 1 {
			return nil, fmt.Errorf(
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
			matches := reduceParamRegex.FindStringSubmatch(param.Name())
			if len(matches) == 0 {
				return nil, fmt.Errorf(
					"%v: invalid parameter name: %v",
					method.Name(), param.Name())
			}
			reduceParam := &ReduceParam{
				Type: param.Type(),
			}
			reduceParam.TermIndex, _ = strconv.Atoi(matches[1])
			reduceMethod.Params = append(reduceMethod.Params, reduceParam)
		}

		state.ReduceMethods[reduceMethod.ProdName] = reduceMethod
	}

	return nil, nil
}
