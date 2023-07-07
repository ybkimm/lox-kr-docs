package lr1

import "github.com/dcaiafa/lox/internal/parsergen/grammar"

func prod(terms ...*grammar.Term) *grammar.Prod {
	return &grammar.Prod{
		Terms: terms,
	}
}

func term(symName string, q ...grammar.Cardinality) *grammar.Term {
	t := &grammar.Term{
		Name: symName,
	}
	if len(q) != 0 {
		t.Cardinality = q[0]
	}
	return t
}
