package parser

import (
	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/token"
)

type Token = *token.Token

type Parser struct {
	loxParser
}

func (p *Parser) reduceSpec(s []ast.Section) *ast.Spec {
	return &ast.Spec{
		Sections: s,
	}
}

func (p *Parser) reduceSection(s ast.Section) ast.Section {
	return s
}

func (p *Parser) reduceParser(_ Token, decls []ast.ParserDecl) ast.Section {
	return &ast.Parser{
		Decls: decls,
	}
}

func (p *Parser) reducePdecl(r *ast.Rule) ast.ParserDecl {
	return r
}

func (p *Parser) reducePrule(name Token, _ Token, prods []*ast.Prod, _ Token) *ast.Rule {
	return &ast.Rule{
		Name:  name.Str,
		Prods: prods,
	}
}

func (p *Parser) reducePprod(terms []*ast.Term, label *ast.Label) *ast.Prod {
	return &ast.Prod{
		Terms: terms,
		Label: label,
	}
}

func (p *Parser) reducePterm(name Token, q ast.Qualifier) *ast.Term {
	return &ast.Term{
		Name:      name.Str,
		Qualifier: q,
	}
}

func (p *Parser) reducePcard(card Token) ast.Qualifier {
	switch card.Type {
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

func (p *Parser) reduceLabel(l Token) *ast.Label {
	return &ast.Label{
		Label: l.Str,
	}
}

func (p *Parser) reduceLexer(_ Token, decls []ast.LexerDecl) *ast.Lexer {
	return &ast.Lexer{
		Decls: decls,
	}
}

func (p *Parser) reduceLcustom(_ Token, names []Token, _ Token) ast.LexerDecl {
	d := &ast.CustomTokenDecl{
		CustomTokens: make([]*ast.CustomToken, len(names)),
	}
	for i, name := range names {
		d.CustomTokens[i] = &ast.CustomToken{
			Name: name.Str,
		}
	}
	return d
}
