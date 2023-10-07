package codegen2

import (
	gotoken "go/token"
	"os"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/lexergen/parser"
)

const LoxFileExtension = ".lox"

// ParseLox parses and checks all .lox files in a directory.
func ParseLox(
	fset *gotoken.FileSet,
	dir string,
	errs *errlogger.ErrLogger,
) *ast.Context {
	loxFiles, err := filepath.Glob(filepath.Join(dir, "*"+LoxFileExtension))
	if err != nil {
		errs.GeneralError(err)
		return nil
	}

	if len(loxFiles) == 0 {
		errs.GeneralErrorf("%v contains no %v files", dir, LoxFileExtension)
		return nil
	}

	spec := new(ast.Spec)
	for _, loxFileName := range loxFiles {
		data, err := os.ReadFile(loxFileName)
		if err != nil {
			errs.GeneralError(err)
			return nil
		}
		file := fset.AddFile(loxFileName, -1, len(data))
		unit := parser.Parse(file, data, errs)
		if errs.HasError() {
			return nil
		}
		spec.Units = append(spec.Units, unit)
	}

	ctx := ast.NewContext(fset, errs)
	ctx.Analyze(spec, ast.AllPasses)
	if errs.HasError() {
		return nil
	}

	return ctx
}
