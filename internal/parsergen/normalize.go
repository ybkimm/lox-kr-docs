package parsergen

import (
	"fmt"
)

func (g *Grammar) Analyze() error {
	ctx := newContext()

	g.prepare(ctx)
	if err := ctx.Err(); err != nil {
		return err
	}

	g.resolveRefs(ctx)
	if err := ctx.Err(); err != nil {
		return err
	}

	g.normalize(ctx)
	if err := ctx.Err(); err != nil {
		return err
	}

	return nil
}

func (g *Grammar) prepare(ctx *context) {
	g.defs = make(map[string]Def)

	for _, terminal := range g.Terminals {
		if other := g.defs[terminal.DefName()]; other != nil {
			ctx.Fail(&RedeclaredError{Def: terminal, Other: other})
			continue
		}
		g.defs[terminal.DefName()] = terminal
	}

	for _, rule := range g.Rules {
		if other := g.defs[rule.DefName()]; other != nil {
			ctx.Fail(&RedeclaredError{Def: rule, Other: other})
			continue
		}
		g.defs[rule.DefName()] = rule
	}
}

func (g *Grammar) resolveRefs(ctx *context) {
	for _, rule := range g.Rules {
		for _, prod := range rule.Prods {
			for _, term := range prod.Terms {
				def := g.defs[term.Name]
				if def == nil {
					ctx.Fail(&UndefinedError{Term: term, Prod: prod, Rule: rule})
					continue
				}
				term.def = def
			}
		}
	}
}

func (g *Grammar) normalize(ctx *context) {
	changed := true
	for changed {
		changed = false
		for _, rule := range g.Rules {
			for _, prod := range rule.Prods {
				for i, term := range prod.Terms {
					switch term.Qualifier {
					case NoQualifier:
					case ZeroOrMore:
						// a = b c*
						//  =>
						// a = b a_0
						// a_0 = c+ | e
						srule := g.synthesizeRule(ctx, rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(term.def, OneOrMore)),
							makeProd(),
						}
						prod.Terms[i] = makeTerm(srule)
						changed = true
					case OneOrMore:
						// a = b c+
						//  =>
						// a = b a_0
						// a_0 = a_0 c
						//     | c
						srule := g.synthesizeRule(ctx, rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(srule), makeTerm(term.def)),
							makeProd(makeTerm(term.def)),
						}
						prod.Terms[i] = makeTerm(srule)
						changed = true
					case ZeroOrOne:
						// a = b c?
						//  =>
						// a = b a_0
						// a_0 = c | e
						srule := g.synthesizeRule(ctx, rule.Name)
						srule.Prods = []*Prod{
							makeProd(makeTerm(term.def)),
							makeProd(),
						}
						prod.Terms[i] = makeTerm(srule)
						changed = true
					default:
						panic("not reached")
					}
				}
			}
		}
	}
}

func (g *Grammar) synthesizeRule(ctx *context, namePrefix string) *Rule {
	r := &Rule{
		Name: fmt.Sprintf("%s__%d", namePrefix, len(g.defs)),
	}
	g.defs[r.Name] = r
	g.Rules = append(g.Rules, r)
	return r
}

func makeTerm(def Def, q ...Qualifier) *Term {
	t := &Term{
		Name: def.DefName(),
		def:  def,
	}
	if len(q) != 0 {
		t.Qualifier = q[0]
	}
	return t
}

func makeProd(terms ...*Term) *Prod {
	return &Prod{
		Terms: terms,
	}
}
