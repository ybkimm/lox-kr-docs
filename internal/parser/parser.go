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

func (p *parser) on_prod(terms []*ast.Term, qualif *ast.ProdQualifier) *ast.Prod {
	return &ast.Prod{
		Terms:     terms,
		Qualifier: qualif,
	}
}

func (p *parser) on_term_card(term *ast.Term, typ ast.TermType) *ast.Term {
	if typ == ast.Simple {
		return term
	}
	return &ast.Term{
		Type:  typ,
		Child: term,
	}
}

func (p *parser) on_term__token(tok Token) *ast.Term {
	switch tok.Type {
	case ID:
		return &ast.Term{Name: tok.Str}
	case LITERAL:
		return &ast.Term{Alias: tok.Str}
	default:
		panic("not-reached")
	}
}

func (p *parser) on_term__list(listTerm *ast.Term) *ast.Term {
	return listTerm
}

func (p *parser) on_list(_, _ Token, elem *ast.Term, _ Token, sep *ast.Term, _ Token) *ast.Term {
	return &ast.Term{
		Type:  ast.List,
		Child: elem,
		Sep:   sep,
	}
}

func (p *parser) on_card(card Token) ast.TermType {
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

func (p *parser) on_token_decl(_ Token, tokens []*ast.CustomToken, _ Token) *ast.CustomTokenDecl {
	return &ast.CustomTokenDecl{
		CustomTokens: tokens,
	}
}

func (p *parser) on_token(name Token, alias Token) *ast.CustomToken {
	t := &ast.CustomToken{
		Name: name.Str,
	}
	if alias.Type == LITERAL {
		t.Alias = alias.Str
	}
	return t
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
