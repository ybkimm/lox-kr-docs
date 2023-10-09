package codegen2

import (
	"os"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/lexergen/parser"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

const LoxFileExtension = ".lox"

// parseLox parses and checks all .lox files in the project directory.
func (c *context) ParseLox() bool {
	loxFiles, err := filepath.Glob(filepath.Join(c.Dir, "*"+LoxFileExtension))
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	if len(loxFiles) == 0 {
		c.Errs.GeneralErrorf("%v contains no %v files", c.Dir, LoxFileExtension)
		return false
	}

	spec := new(ast.Spec)
	for _, loxFileName := range loxFiles {
		data, err := os.ReadFile(loxFileName)
		if err != nil {
			c.Errs.GeneralError(err)
			return false
		}
		file := c.Fset.AddFile(loxFileName, -1, len(data))
		unit := parser.Parse(file, data, c.Errs)
		if c.Errs.HasError() {
			return false
		}
		spec.Units = append(spec.Units, unit)
	}

	astctx := ast.NewContext(c.Fset, c.Errs)
	astctx.Analyze(spec, ast.AllPasses)
	if c.Errs.HasError() {
		return false
	}

	c.ParserGrammar = astctx.Grammar
	c.ParserTable = lr2.ConstructLALR(c.ParserGrammar)

	return !c.Errs.HasError()
}
