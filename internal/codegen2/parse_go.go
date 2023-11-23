package codegen2

import (
	"fmt"
	gotoken "go/token"
	gotypes "go/types"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/base/assert"
	"golang.org/x/tools/go/packages"
)

func (c *context) ParseGo() bool {
	assert.True(c.GoPackageName != "")

	// parser.gen.go does not exist yet (or will be excluded), but there are
	// parts of that code that the user code is allowed to reference. To allow
	// Go parsing/analysis to succeed we need to generate a placeholder for
	// parser.gen.go with the bare minimum API.
	parserGenPlaceholder := renderParserTemplate(&parserTemplateInputs{
		Placeholder: true,
		Package:     c.GoPackageName,
	})
	lexerGenPlaceholder := renderLexerTemplatePlaceholder(c.GoPackageName)

	absDir, err := filepath.Abs(c.Dir)
	if err != nil {
		panic(err)
	}

	lexerGenPath := filepath.Join(absDir, lexerGenGo)
	parserGenPath := filepath.Join(absDir, parserGenGo)

	// Parse and analyze Go sources in the project directory.
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax,
		Dir:  filepath.Clean(c.Dir),
		Fset: c.Fset,
		Overlay: map[string][]byte{
			// Inject the placeholder implementations.
			parserGenPath: []byte(parserGenPlaceholder),
			lexerGenPath:  []byte(lexerGenPlaceholder),
		},
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	c.GoPackagePath = pkgs[0].PkgPath

	if len(pkgs[0].Errors) != 0 {
		for _, err := range pkgs[0].Errors {
			c.Errs.GeneralError(err)
		}
		return false
	}

	// Get a hold of some important types:

	// The "Token" type must be provided by the programmer.
	scope := pkgs[0].Types.Scope()
	tokenObj := scope.Lookup("Token")
	if tokenObj == nil {
		c.Errs.GeneralErrorf("Token type is undefined")
		return false
	}
	c.TokenType = tokenObj.Type()

	// The "error" built-in type, it will be used for matching error-handling
	// productions.
	errorObj := gotypes.Universe.Lookup("error")
	if errorObj == nil {
		panic("error is undefined")
	}
	c.ErrorType = errorObj.Type()

	// The programmer provided parser implementation.
	c.lookupParserType(scope)

	return !c.Errs.HasError()
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
func (c *context) lookupParserType(scope *gotypes.Scope) {
	loxObj := scope.Lookup(parserStateName)
	if loxObj == nil {
		// This type is generated so this should always succeed.
		panic(fmt.Errorf("could not find type %q", parserStateName))
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
			c.Errs.Errorf(
				c.Fset.Position(obj.Pos()),
				"parser %v cannot have type parameters",
				namedType.Obj().Name())
			return
		}

		// There can be only one.
		if parserObj != nil {
			c.Errs.Errorf(
				c.Fset.Position(obj.Pos()),
				"there can only be one parser implementation")
			c.Errs.Infof(
				c.Fset.Position(parserObj.Obj().Pos()),
				"here is the other one")
			return
		}

		parserObj = namedType
	}

	if parserObj == nil {
		c.Errs.GeneralErrorf(
			"parser implementation undefined")
		c.Errs.Infof(
			gotoken.Position{},
			"You must define a struct for the parser that embeds \"lox\".\n"+
				"Example:\n"+
				"type myParser struct {\n"+
				"  lox  // <== must embed this\n"+
				"}")
		return
	}

	c.ParserType = parserObj
}
