package parsergen

func makeTerm(sym Symbol, q ...Qualifier) *Term {
	t := &Term{
		Name: sym.SymName(),
		sym:  sym,
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
