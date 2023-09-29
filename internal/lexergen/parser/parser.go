package parser

import (
	gotoken "go/token"
	"unicode/utf8"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/ast"
)

type parser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
	unit *ast.Unit
}

func Parse(file *gotoken.File, data []byte, errs *errlogger.ErrLogger) *ast.Unit {
	lex := newLex(file, data, errs)
	p := &parser{
		file: file,
		errs: errs,
	}
	p.parse(lex)
	return p.unit
}

func (p *parser) on_spec(stmts []ast.Statement) *ast.Unit {
	p.unit = &ast.Unit{
		Statements: stmts,
	}
	return p.unit
}

func (p *parser) on_statement(s ast.Statement) ast.Statement {
	return s
}

func (p *parser) on_rule(r ast.Statement) ast.Statement {
	return r
}

func (p *parser) on_mode(_ Token, name Token, _ Token, rules []ast.Statement, _ Token) *ast.Mode {
	return &ast.Mode{
		Name:  name.Str,
		Rules: rules,
	}
}

func (p *parser) on_token_rule(id Token, _ Token, expr *ast.Expr, actions []ast.Action, _ Token) *ast.TokenRule {
	return &ast.TokenRule{
		Name:    id.Str,
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_frag_rule(_ Token, expr *ast.Expr, actions []ast.Action, _ Token) *ast.FragRule {
	return &ast.FragRule{
		Expr:    expr,
		Actions: actions,
	}
}

func (p *parser) on_macro_rule(_ Token, name Token, _ Token, expr *ast.Expr, _ Token) *ast.MacroRule {
	return &ast.MacroRule{
		Name: name.Str,
		Expr: expr,
	}
}

func (p *parser) on_expr(factors []*ast.Factor) *ast.Expr {
	return &ast.Expr{
		Factors: factors,
	}
}

func (p *parser) on_factor(terms []*ast.TermCard) *ast.Factor {
	return &ast.Factor{
		Terms: terms,
	}
}

func (p *parser) on_term_card(term ast.Term, card ast.Card) *ast.TermCard {
	return &ast.TermCard{
		Term: term,
		Card: card,
	}
}

func (p *parser) on_card(c Token) ast.Card {
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

func (p *parser) on_term__tok(tok Token) ast.Term {
	switch tok.Type {
	case LITERAL:
		return &ast.TermLiteral{
			Literal: tok.Str,
		}
	case ID:
		return &ast.TermRef{
			Ref: tok.Str,
		}
	default:
		panic("unreachable")
	}
}

func (p *parser) on_term__char_class(charClass *ast.TermCharClass) ast.Term {
	return charClass
}

func (p *parser) on_term__expr(_ Token, expr *ast.Expr, _ Token) ast.Term {
	return expr
}

func (p *parser) on_char_class(neg Token, _ Token, chars []Token, _ Token) *ast.TermCharClass {
	items := make([]*ast.CharClassItem, 0, len(chars))

	toRune := func(t Token) rune {
		r, _ := utf8.DecodeRuneInString(t.Str)
		return r
	}
	addItem := func(b, e Token) {
		items = append(items, &ast.CharClassItem{
			From: toRune(b),
			To:   toRune(e),
		})
	}

	for i := 0; i < len(chars); i++ {
		if i+2 > len(chars)-1 || chars[i+1].Type != DASH {
			addItem(chars[i], chars[i])
		} else {
			addItem(chars[i], chars[i+2])
			i += 2
		}
	}

	return &ast.TermCharClass{
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

func (p *parser) on_action_skip(_ Token) *ast.ActionSkip {
	return &ast.ActionSkip{}
}

func (p *parser) on_action_push_mode(_, _ Token, mode Token, _ Token) *ast.ActionPushMode {
	return &ast.ActionPushMode{
		Mode: mode.Str,
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
	p.errs.Errorf(
		p.file.Position(p.errorToken().Pos),
		"unexpected %v", p.errorToken())
}
