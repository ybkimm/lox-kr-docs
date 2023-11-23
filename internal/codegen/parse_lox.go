package codegen

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/parser"
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

	c.LexerModes = astctx.LexerDFAs
	c.ParserGrammar = astctx.Grammar
	c.ParserTable = lr2.ConstructLALR(c.ParserGrammar)

	if c.Report != nil {
		fmt.Fprintln(c.Report, "Parser Table")
		fmt.Fprintln(c.Report, "============")
		c.ParserTable.Print(c.Report)

		var modes []*mode.Mode
		for _, mode := range c.LexerModes {
			modes = append(modes, mode)
		}
		slices.SortFunc(modes, func(a, b *mode.Mode) int {
			return cmp.Compare(a.Index, b.Index)
		})
		fmt.Fprintln(c.Report, "")

		fmt.Fprintln(c.Report, "Lexer Modes")
		fmt.Fprintln(c.Report, "============")
		for _, mode := range modes {
			fmt.Fprintln(c.Report, mode.Name+":")
			mode.DFA.Print(c.Report)
			fmt.Fprintln(c.Report, "")
		}
	}

	return !c.Errs.HasError()
}
