package analysis

import (
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/grammar"
	"github.com/dcaiafa/lox/internal/loc"
)

type analyzer struct {
	syntax *grammar.Spec
	errs   *errs.Errs
	decls  map[string]grammar.Decl
}

func Analyze(s *grammar.Spec, errs *errs.Errs) {
	a := &analyzer{
		syntax: s,
		errs:   errs,
		decls:  make(map[string]grammar.Decl),
	}
	a.prepare()
	if a.errs.HasErrors() {
		return
	}
	/*
		a.checkReferences()
		if a.errs.HasErrors() {
			return
		}
		for a.normalize() {
			// Run until completely normalized.
		}
	*/
}

func (a *analyzer) prepare() {
	for _, rule := range a.syntax.Parser.Rules {
		if _, ok := a.decls[rule.DeclName()]; ok {
			a.errs.Errorf(loc.Loc{}, "%q redeclared", rule.DeclName())
		}
		a.decls[rule.DeclName()] = rule
	}
	if a.errs.HasErrors() {
		return
	}
}

/*
func (a *analyzer) checkReferences() {
	for _, decl := range a.syntax.Decls {
		rule, ok := decl.(*grammar.Rule)
		if !ok {
			continue
		}
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				decl := a.decls[term.Name]
				if decl == nil {
					a.errs.Errorf(loc.Loc{}, "undefined: %v", term.Name)
				}
			}
		}
	}
}

func (a *analyzer) normalize() bool {
	changed := false
	for _, decl := range a.syntax.Decls {
		rule, ok := decl.(*grammar.Rule)
		if !ok {
			continue
		}

		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				switch term.Qualifier {
				case grammar.NoQualifier:
				case grammar.ZeroOrMore:
					// a = b c*
					//  =>
					// a = b a_0
					// a_0 = c+ | e
					srule := a.synthesizeRule(rule.Name)
					srule.Prods = []*grammar.Prod{
						{Terms: []*grammar.Term{{Name: term.Name, Qualifier: grammar.OneOrMore}}},
						{Terms: []*grammar.Term{}},
					}
					changed = true
				case grammar.OneOrMore:
					// a = b c+
					//  =>
					// a = b a_0
					// a_0 = a_0 c
					//     | c
					srule := a.synthesizeRule(rule.Name)
					srule.Prods = []*grammar.Prod{
						{Terms: []*grammar.Term{{Name: srule.Name}, {Name: term.Name}}},
						{Terms: []*grammar.Term{{Name: term.Name}}},
					}
					changed = true
				case grammar.ZeroOrOne:
					// a = b c?
					//  =>
					// a = b a_0
					// a_0 = c | e
					srule := a.synthesizeRule(rule.Name)
					srule.Prods = []*grammar.Prod{
						{Terms: []*grammar.Term{{Name: term.Name}}},
						{Terms: []*grammar.Term{}},
					}
					changed = true
				default:
					panic("not reached")
				}
			}
		}
	}
	return changed
}

func (a *analyzer) synthesizeRule(namePrefix string) *grammar.Rule {
	r := &grammar.Rule{
		Name: fmt.Sprintf("%s__%d", namePrefix, len(a.decls)),
	}
	a.decls[r.Name] = r
	a.syntax.Decls = append(a.syntax.Decls, r)
	return r
}
*/
