package analyzer

import (
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
)

type analyzer struct {
	spec *ast.Spec
	errs *errlogger.ErrLogger
}

func newAnalyzer(spec *ast.Spec) *analyzer {
	return &analyzer{
		spec: spec,
	}
}

func (a *analyzer) Check() {

}
