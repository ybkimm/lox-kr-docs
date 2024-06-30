package codegen

import (
	"bytes"
	"fmt"
	goformat "go/format"
	gotypes "go/types"

	"github.com/CloudyKit/jet/v6"
)

func renderTemplate(
	packageName string,
	packagePath string,
	templ string,
	vars jet.VarMap,
) string {
	imports := newImports()

	loader := jet.NewInMemLoader()
	loader.Set("lox", templ)

	set := jet.NewSet(loader, jet.WithSafeWriter(nil))
	t, err := set.GetTemplate("lox")
	if err != nil {
		panic(err)
	}

	// Functions available to all templates:

	// imp(importPath string): replaces import_path with an alias, and adds the
	// import to top import prologue.
	vars.Set("imp", imports.Import)

	// go_type(t gotypes.Type): returns the Go type name including package prefix
	// if applicable.
	vars.Set("go_type", func(t gotypes.Type) string {
		return gotypes.TypeString(t, func(pkg *gotypes.Package) string {
			if pkg.Path() == packagePath {
				return ""
			}
			return imports.Import(pkg.Path())
		})
	})

	body := &bytes.Buffer{}
	err = t.Execute(body, vars, nil)
	if err != nil {
		panic(err)
	}

	full := bytes.NewBuffer(make([]byte, 0, body.Len()+2048))
	fmt.Fprintf(full, "package %v\n\n", packageName)
	imports.WriteTo(full)
	body.WriteTo(full)

	fullFormatted, err := goformat.Source(full.Bytes())
	if err != nil {
		panic(fmt.Errorf("failed to format %v: %w", lexerGenGo, err))
	}

	return string(fullFormatted)
}
