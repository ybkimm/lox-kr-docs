package parser

import (
	gotoken "go/token"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
	"github.com/dcaiafa/lox/internal/base/baselexer"
)

type Token = baselexer.Token

type parser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
	unit *ast.Unit
}

func Parse(file *gotoken.File, data []byte, errs *errlogger.ErrLogger) *ast.Unit {
	onError := func(l *baselexer.Lexer) {
		errs.Errorf(file.Position(l.Pos()), "unexpected character: %c", l.Peek())
	}

	lex := baselexer.New(new(_LexerStateMachine), onError, file, data)

	p := &parser{
		file: file,
		errs: errs,
	}
	p.parse(lex)
	return p.unit
}

func (p *parser) on_spec(sections [][]ast.Statement) *ast.Unit {
	n := 0
	for _, section := range sections {
		n += len(section)
	}
	stmts := make([]ast.Statement, 0, n)
	for _, section := range sections {
		stmts = append(stmts, section...)
	}
	p.unit = &ast.Unit{
		Statements: stmts,
	}
	return p.unit
}

func (p *parser) on_section(sectionStmts []ast.Statement) []ast.Statement {
	return sectionStmts
}

// Parser
// ======

func (p *parser) on_parser_section(_ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_parser_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_parser_rule(start Token, name Token, _ Token, prods []*ast.ParserProd, _ Token) *ast.ParserRule {
	return &ast.ParserRule{
		IsStart: start.Type == START,
		Name:    string(name.Str),
		Prods:   prods,
	}
}

func (p *parser) on_parser_prod(terms []*ast.ParserTerm, qualif *ast.ProdQualifier) *ast.ParserProd {
	return &ast.ParserProd{
		Terms:     terms,
		Qualifier: qualif,
	}
}

func (p *parser) on_parser_term_card(term *ast.ParserTerm, typ ast.ParserTermType) *ast.ParserTerm {
	if typ == ast.ParserTermSimple || typ == ast.ParserTermError {
		return term
	}
	return &ast.ParserTerm{
		Type:  typ,
		Child: term,
	}
}

func (p *parser) on_parser_term__token(tok Token) *ast.ParserTerm {
	switch tok.Type {
	case ID:
		return &ast.ParserTerm{Name: string(tok.Str)}
	case LITERAL:
		return &ast.ParserTerm{Alias: fixLiteral(tok.Str)}
	case ERROR_KEYWORD:
		return &ast.ParserTerm{Type: ast.ParserTermError}
	default:
		panic("not-reached")
	}
}

func (p *parser) on_parser_term__list(listTerm *ast.ParserTerm) *ast.ParserTerm {
	return listTerm
}

func (p *parser) on_parser_list(_, _ Token, elem *ast.ParserTerm, _ Token, sep *ast.ParserTerm, _ Token) *ast.ParserTerm {
	return &ast.ParserTerm{
		Type:  ast.ParserTermList,
		Child: elem,
		Sep:   sep,
	}
}

func (p *parser) on_parser_card(card Token) ast.ParserTermType {
	switch card.Type {
	case ZERO_OR_MORE:
		return ast.ParserTermZeroOrMore
	case ONE_OR_MORE:
		return ast.ParserTermOneOrMore
	case ZERO_OR_ONE:
		return ast.ParserTermZeroOrOne
	default:
		panic("unreachable")
	}
}

func (p *parser) on_parser_qualif(assoc Token, _ Token, prec Token, _ Token) *ast.ProdQualifier {
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
	q.Precedence, err = strconv.Atoi(string(prec.Str))
	if err != nil {
		panic(err)
	}
	if q.Precedence <= 0 {
		panic("not-reached")
	}

	return q
}

// Lexer
// =====

func (p *parser) on_lexer_section(_ Token, stmts []ast.Statement) []ast.Statement {
	return stmts
}

func (p *parser) on_lexer_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_lexer_rule(r ast.Statement) ast.Statement {
	return r
}

func (p *parser) on_mode(_ Token, name Token, _ Token, rules []ast.Statement, _ Token) *ast.Mode {
	return &ast.Mode{
		Name:  string(name.Str),
		Rules: rules,
	}
}

func (p *parser) on_token_rule(id Token, _ Token, expr *ast.LexerExpr, actions []ast.Action, _ Token) *ast.TokenRule {
	return &ast.TokenRule{
		Name:    string(id.Str),
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_frag_rule(_ Token, expr *ast.LexerExpr, actions []ast.Action, _ Token) *ast.FragRule {
	return &ast.FragRule{
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_macro_rule(_ Token, name Token, _ Token, expr *ast.LexerExpr, _ Token) *ast.MacroRule {
	return &ast.MacroRule{
		Name: string(name.Str),
		Expr: expr,
	}
}

func (p *parser) on_lexer_expr(factors []*ast.LexerFactor) *ast.LexerExpr {
	return &ast.LexerExpr{
		Factors: factors,
	}
}

func (p *parser) on_lexer_factor(terms []*ast.LexerTermCard) *ast.LexerFactor {
	return &ast.LexerFactor{
		Terms: terms,
	}
}

func (p *parser) on_lexer_term_card(term ast.LexerTerm, card ast.Card) *ast.LexerTermCard {
	return &ast.LexerTermCard{
		Term: term,
		Card: card,
	}
}

func (p *parser) on_lexer_card(c Token) ast.Card {
	switch c.Type {
	case ZERO_OR_ONE:
		return ast.ZeroOrOne
	case ZERO_OR_MORE:
		return ast.ZeroOrMore
	case ONE_OR_MORE:
		return ast.OneOrMore
	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__tok(tok Token) ast.LexerTerm {
	switch tok.Type {
	case LITERAL:
		return &ast.LexerTermLiteral{
			Literal: fixLiteral(tok.Str),
		}
	case ID:
		return &ast.LexerTermRef{
			Ref: string(tok.Str),
		}
	default:
		panic("unreachable")
	}
}

func (p *parser) on_lexer_term__char_class(charClass *ast.LexerTermCharClass) ast.LexerTerm {
	return charClass
}

func (p *parser) on_lexer_term__expr(_ Token, expr *ast.LexerExpr, _ Token) ast.LexerTerm {
	return expr
}

func (p *parser) on_char_class(neg Token, _ Token, chars []Token, _ Token) *ast.LexerTermCharClass {
	items := make([]*ast.CharClassItem, 0, len(chars))

	toRune := func(t Token) rune {
		r, _ := utf8.DecodeRuneInString(unescape(t.Str))
		return r
	}
	addItem := func(b, e Token) {
		items = append(items, &ast.CharClassItem{
			From: toRune(b),
			To:   toRune(e),
		})
	}

	for i := 0; i < len(chars); i++ {
		if i+2 > len(chars)-1 || chars[i+1].Type != CLASS_DASH {
			addItem(chars[i], chars[i])
		} else {
			addItem(chars[i], chars[i+2])
			i += 2
		}
	}

	return &ast.LexerTermCharClass{
		Neg:            neg.Type != EOF,
		CharClassItems: items,
	}
}

func (p *parser) on_char_class_item(c Token) Token {
	return c
}

func (p *parser) on_action(action ast.Action) ast.Action {
	return action
}

func (p *parser) on_action_discard(_ Token) *ast.ActionDiscard {
	return &ast.ActionDiscard{}
}

func (p *parser) on_action_push_mode(_, _ Token, mode Token, _ Token) *ast.ActionPushMode {
	return &ast.ActionPushMode{
		Mode: string(mode.Str),
	}
}

func (p *parser) on_action_pop_mode(_ Token) *ast.ActionPopMode {
	return &ast.ActionPopMode{}
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

func (p *parser) onError() {
	tok := p.errorToken()
	p.errs.Errorf(
		p.file.Position(tok.Pos),
		"unexpected %v %q", _TokenToString(tok.Type), string(tok.Str))
}

func fixLiteral(lit []byte) string {
	return unescape(lit[1 : len(lit)-1])
}

func unescape(lit []byte) string {
	var str strings.Builder

	for i := 0; i < len(lit); i++ {
		if lit[i] == '\\' {
			switch lit[i+1] {
			case 'n':
				str.WriteRune('\n')
				i++
			case 'r':
				str.WriteRune('\r')
				i++
			case 't':
				str.WriteRune('\t')
				i++
			case '\'':
				str.WriteRune('\'')
				i++
			case '\\':
				str.WriteRune('\\')
				i++
			case '-':
				str.WriteRune('-')
				i++
			case 'x':
				str.WriteByte(byte(hexToRune(string(lit[i+2 : i+4]))))
				i += 3
			case 'u':
				str.WriteRune(hexToRune(string(lit[i+2 : i+6])))
				i += 5
			case 'U':
				str.WriteRune(hexToRune(string(lit[i+2 : i+10])))
				i += 9
			default:
				panic("unreachable")
			}
		} else {
			str.WriteByte(lit[i])
		}
	}

	return str.String()
}

func hexToRune(str string) rune {
	v, err := strconv.ParseUint(string(str), 16, 32)
	if err != nil {
		panic(err)
	}
	return rune(v)
}
