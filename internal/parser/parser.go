package parser

import (
	gotoken "go/token"
	"strconv"

	"github.com/dcaiafa/lox/internal/ast"
)

type parser struct {
	loxParser
	spec *ast.Spec
}

func Parse(file *gotoken.File, data []byte, errLogger _lxErrorLogger) (*ast.Spec, bool) {
	var parser parser
	lex := newLex(file, data, errLogger)
	ok := parser.parse(lex, errLogger)
	return parser.spec, ok
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

func (p *parser) reducePprod(terms []*ast.Term, qualif *ast.ProdQualifier) *ast.Prod {
	return &ast.Prod{
		Terms:     terms,
		Qualifier: qualif,
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

func (p *parser) reduceLexer(_ Token, decls []ast.LexerDecl) *ast.Lexer {
	return &ast.Lexer{
		Decls: decls,
	}
}

func (p *parser) reduceLdecl(d ast.LexerDecl) ast.LexerDecl {
	return d
}

func (p *parser) reducePqualif(assoc Token, _ Token, prec Token, _ Token) *ast.ProdQualifier {
	q := &ast.ProdQualifier{}

	switch assoc.Type {
	case LEFT:
		q.Associativity = ast.Left
	case RIGHT:
		q.Associativity = ast.Right
	default:
		panic("not-reached")
	}

	var err error
	q.Precedence, err = strconv.Atoi(prec.Str)
	if err != nil {
		panic(err)
	}
	if q.Precedence <= 0 {
		panic("not-reached")
	}

	return q
}

func (p *parser) reduceLtoken(_ Token, names []Token, _ Token) ast.LexerDecl {
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
