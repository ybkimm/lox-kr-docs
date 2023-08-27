package parser

import (
	gotoken "go/token"
	"strconv"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errlogger"
)

type errLogger struct {
	*errlogger.ErrLogger
	File *gotoken.File
}

func (e *errLogger) ParserError(err error) {
	var pos gotoken.Position
	if err, ok := err.(interface{ Pos() Token }); ok {
		pos = e.File.Position(err.Pos().Pos)
	}
	e.Errorf(pos, "%v", err)
}

type parser struct {
	loxParser
	parserAST *ast.Parser
}

func Parse(file *gotoken.File, data []byte, errs *errlogger.ErrLogger) (*ast.Parser, bool) {
	errLogger := &errLogger{
		ErrLogger: errs,
		File:      file,
	}

	var parser parser
	lex := newLex(file, data, errs)
	ok := parser.parse(lex, errLogger)
	return parser.parserAST, ok
}

func (p *parser) on_parser(decls []ast.ParserDecl) *ast.Parser {
	p.parserAST = &ast.Parser{
		Decls: decls,
	}
	return p.parserAST
}

func (p *parser) on_decl(r ast.ParserDecl) ast.ParserDecl {
	return r
}

func (p *parser) on_rule(name Token, _ Token, prods []*ast.Prod, _ Token) *ast.Rule {
	return &ast.Rule{
		Name:  name.Str,
		Prods: prods,
	}
}

func (p *parser) on_prods(prods []*ast.Prod, _ Token, prod *ast.Prod) []*ast.Prod {
	return append(prods, prod)
}

func (p *parser) on_prods__1(prod *ast.Prod) []*ast.Prod {
	return []*ast.Prod{prod}
}

func (p *parser) on_prod(terms []*ast.Term, qualif *ast.ProdQualifier) *ast.Prod {
	return &ast.Prod{
		Terms:     terms,
		Qualifier: qualif,
	}
}

func (p *parser) on_term(name Token, q ast.Qualifier) *ast.Term {
	return &ast.Term{
		Name:      name.Str,
		Qualifier: q,
	}
}

func (p *parser) on_card(card Token) ast.Qualifier {
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

func (p *parser) on_qualif(assoc Token, _ Token, prec Token, _ Token) *ast.ProdQualifier {
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

func (p *parser) on_token(_ Token, names []Token, _ Token) *ast.CustomTokenDecl {
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

func (p *parser) onReduce(r any, begin, end Token) {
	rAST, ok := r.(ast.AST)
	if !ok {
		return
	}
	rAST.SetBounds(ast.Bounds{
		Begin: begin.Pos,
		End:   end.Pos,
	})
}
