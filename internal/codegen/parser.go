package codegen

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	_ "golang.org/x/tools/go/loader"
)

const loxGenGo = `
package {{package}}

type loxParser struct {}
`

type State struct {
	Files map[string]*File
}

type File struct {
	Filename string
}

func Parse(path string) (*State, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var astFiles []*ast.File
	fset := token.NewFileSet()
	for _, entry := range entries {
		if entry.IsDir() ||
			filepath.Ext(entry.Name()) != ".go" ||
			entry.Name() == "lox.gen.go" {
			continue
		}
		filename := filepath.Join(path, entry.Name())
		astFile, err := parseFile(fset, filename)
		if err != nil {
			return nil, err
		}
		astFiles = append(astFiles, astFile)
	}

	if len(astFiles) == 0 {
		return nil, fmt.Errorf("no files")
	}

	packageName := astFiles[0].Name.Name
	loxGenGo := strings.Replace(loxGenGo, "{{package}}", packageName, 1)

	loxGenGoAST, err := parser.ParseFile(
		fset, "lox.gen.go", loxGenGo, parser.DeclarationErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lox.gen.go: %w", err)
	}

	astFiles = append(astFiles, loxGenGoAST)

	conf := types.Config{
		Importer:                 importer.Default(),
		IgnoreFuncBodies:         true,
		Error:                    func(err error) { fmt.Println(err) },
		DisableUnusedImportCheck: true,
	}

	pkg, err := conf.Check(packageName, fset, astFiles, nil)
	if err != nil {
		return nil, err
	}

	_ = pkg

	return nil, nil
}

func parseFile(fset *token.FileSet, filename string) (*ast.File, error) {
	astFile, err := parser.ParseFile(
		fset, filename, nil, parser.ParseComments|parser.DeclarationErrors)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", filename, err)
	}
	return astFile, nil
}

/*
func parseFile_(filename string) (*File, error) {
	astFile, err := parser.ParseFile(fset, filename, nil, parser.ParseComments|parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}

	conf := types.Config{
		Importer:                 importer.Default(),
		IgnoreFuncBodies:         true,
		Error:                    func(err error) { fmt.Println(err) },
		DisableUnusedImportCheck: true,
	}

	pkg, _ := conf.Check("foobar", fset, []*ast.File{astFile}, nil)

	if pkg == nil {
		return nil, errors.New("Check() failed completely")
	}

	scope := pkg.Scope()

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		fmt.Println(obj)
	}

	return nil, nil
}

func findParserStruct(astFile *ast.File) *ast.TypeSpec {
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
      	continue
			}
			for _, field := range structType.Fields.List {
				if field.Names != nil {
					continue
				}


			}
		}
	}
}
*/
