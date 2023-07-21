package parser

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/token"
)

//go:generate go run ../../cmd/lox .

type Token = token.Token

type parser struct {
	loxParser
	spec *ast.Spec
}

func Parse(file *gotoken.File, data []byte) (*ast.Spec, error) {
	var parser parser
	lex := newLex(file, data)
	err := parser.parse(lex)
	if err != nil {
		return nil, err
	}
	return parser.spec, nil
}

func (p *parser) reduceSpec(s []ast.Section) any {
	p.spec = &ast.Spec{
		Sections: s,
	}
	return nil
}

func (p *parser) reduceSection(s ast.Section) ast.Section {
	return s
}

func (p *parser) reduceParser(_ Token, decls []ast.ParserDecl) ast.Section {
	return &ast.Parser{
		Decls: decls,
	}
}

func (p *parser) reducePdecl(r *ast.Rule) ast.ParserDecl {
	return r
}

func (p *parser) reducePrule(name Token, _ Token, prods []*ast.Prod, _ Token) *ast.Rule {
	return &ast.Rule{
		Name:  name.Str,
		Prods: prods,
	}
}

func (p *parser) reducePprods(prods []*ast.Prod, _ Token, prod *ast.Prod) []*ast.Prod {
	return append(prods, prod)
}

func (p *parser) reducePprods_1(prod *ast.Prod) []*ast.Prod {
	return []*ast.Prod{prod}
}

func (p *parser) reducePprod(terms []*ast.Term, label *ast.Label) *ast.Prod {
	return &ast.Prod{
		Terms: terms,
		Label: label,
	}
}

func (p *parser) reducePterm(name Token, q ast.Qualifier) *ast.Term {
	return &ast.Term{
		Name:      name.Str,
		Qualifier: q,
	}
}

func (p *parser) reducePcard(card Token) ast.Qualifier {
	switch card.Type {
	case ZERO_OR_MANY:
		return ast.ZeroOrMore
	case ONE_OR_MANY:
		return ast.OneOrMore
	case ZERO_OR_ONE:
		return ast.ZeroOrOne
	default:
		panic("unreachable")
	}
}

func (p *parser) reduceLabel(l Token) *ast.Label {
	return &ast.Label{
		Label: l.Str,
	}
}

func (p *parser) reduceLexer(_ Token, decls []ast.LexerDecl) *ast.Lexer {
	return &ast.Lexer{
		Decls: decls,
	}
}

func (p *parser) reduceLdecl(d ast.LexerDecl) ast.LexerDecl {
	return d
}

func (p *parser) reduceLcustom(_ Token, names []Token, _ Token) ast.LexerDecl {
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
