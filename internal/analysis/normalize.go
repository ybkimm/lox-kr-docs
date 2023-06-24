package analysis

import (
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/grammar"
)

type analyzer struct {
	errs   *errlogger.ErrLogger
	syntax *grammar.Syntax
	prods  map[string]*grammar.Production
}

func Analyze(s *grammar.Syntax, errs *errlogger.ErrLogger) {
	a := &analyzer{
		errs:   errs,
		syntax: s,
	}
	a.prepare()
	if a.errs.HasErrors() {
		return
	}
	a.normalize()
	if a.errs.HasErrors() {
		return
	}
}

func (a *analyzer) prepare() {
	a.prods = make(map[string]*grammar.Production, len(a.syntax.Productions))
	for _, prod := range a.syntax.Productions {
		if a.prods[prod.Name] != nil {
			a.errs.Errorf("%q redeclared", prod.Name)
		}
		a.prods[prod.Name] = prod
	}
	if a.errs.HasErrors() {
		return
	}
}

func (a *analyzer) normalize() bool {
	for _, prod := range a.prods {
		for _, term := range prod.Terms {
			for _, factor := range term.Factors {
			}
		}
	}
}
