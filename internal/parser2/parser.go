package parser

import (
	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/token"
)

type Parser struct {
	loxParser
}

func (p *Parser) reduceSpec(s1 []ast.Section) *ast.Spec {
	return &ast.Spec{
		Sections: s1,
	}
}

func (p *Parser) reduceSection(s1 ast.Section) ast.Section {
	return s1
}

func (p *Parser) reduceParser(decls2 []ast.ParserDecl) ast.Section {
	return &ast.Parser{
		Decls: decls2,
	}
}

func (p *Parser) reducePdecl(r1 *ast.Rule) ast.ParserDecl {
	return r1
}

func (p *Parser) reducePrule(name1 token.Token, prods3 []*ast.Prod) *ast.Rule {
	return &ast.Rule{
		Name:  name1.Str,
		Prods: prods3,
	}
}

func (p *Parser) reducePprod(terms1 []*ast.Term, label2 *ast.Label) *ast.Prod {
	return &ast.Prod{
		Terms: terms1,
		Label: label2,
	}
}

func (p *Parser) reducePterm(name1 token.Token, q2 ast.Qualifier) *ast.Term {
	return &ast.Term{
		Name:      name1.Str,
		Qualifier: q2,
	}
}

func (p *Parser) reducePcard(card1 token.Token) ast.Qualifier {
	switch card1.Type {
	case token.ZERO_OR_MANY:
		return ast.ZeroOrMore
	case token.ONE_OR_MANY:
		return ast.OneOrMore
	case token.ZERO_OR_ONE:
		return ast.ZeroOrOne
	default:
		panic("unreachable")
	}
}

func (p *Parser) reduceLabel(l1 token.Token) *ast.Label {
	return &ast.Label{
		Label: l1.Str,
	}
}

func (p *Parser) reduceLexer(decls2 []ast.LexerDecl) *ast.Lexer {
	return &ast.Lexer{
		Decls: decls2,
	}
}

func (p *Parser) reduceLcustom(names2 []token.Token) ast.LexerDecl {
	d := &ast.CustomTokenDecl{
		CustomTokens: make([]*ast.CustomToken, len(names2)),
	}
	for i, name := range names2 {
		d.CustomTokens[i] = &ast.CustomToken{
			Name: name.Str,
		}
	}
	return d
}
