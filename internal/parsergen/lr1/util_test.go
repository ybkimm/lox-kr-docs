package lr1

import "github.com/dcaiafa/lox/internal/parsergen/grammar"

func prod(terms ...*grammar.Term) *grammar.Prod {
	return &grammar.Prod{
		Terms: terms,
	}
}

func term(symName string) *grammar.Term {
	t := &grammar.Term{
		Name: symName,
	}
	return t
}
