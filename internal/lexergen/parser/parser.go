package parser

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/errlogger"
)

type parser struct {
	lox
	file *gotoken.File
	errs *errlogger.ErrLogger
}

func Parse(file *gotoken.File, data []byte, errs *errlogger.ErrLogger) {
	lex := newLex(file, data, errs)
	parser := parser{
		file: file,
		errs: errs,
	}
	parser.parse(lex)
}

func (p *parser) on_spec(stmts []any) any {
	return nil
}

func (p *parser) on_statement(s any) any {
	return s
}

func (p *parser) on_rule(r any) any {
	return r
}

func (p *parser) on_mode(_ Token, id Token, _ Token, rules []any, _ Token) any {
	return nil
}

func (p *parser) on_token_rule(id Token, _ Token, expr any, action any, _ Token) any {
	return nil
}

func (p *parser) on_frag_rule(_ Token, expr any, action any, _ Token) any {
	return nil
}

func (p *parser) on_macro_rule(_ Token, id Token, _ Token, expr any, _ Token) any {
	return nil
}

func (p *parser) on_expr(factors []any) any {
	return nil
}

func (p *parser) on_factor(terms []any) any {
	return nil
}

func (p *parser) on_term_card(term any, card any) any {
	return nil
}

func (p *parser) on_card(c Token) any {
	return nil
}

func (p *parser) on_term__tok(tok Token) any {
	return nil
}

func (p *parser) on_term__char_class(charClass any) any {
	return nil
}

func (p *parser) on_term__expr(_ Token, expr any, _ Token) any {
	return nil
}

func (p *parser) on_char_class(neg Token, _ Token, items []any, _ Token) any {
	return nil
}

func (p *parser) on_char_class_item__range(a Token, _ Token, b Token) any {
	return nil
}

func (p *parser) on_char_class_item__char(c Token) any {
	return nil
}

func (p *parser) on_actions(_ Token, actions []any) any {
	return nil
}

func (p *parser) on_action(action any) any {
	return action
}

func (p *parser) on_action_skip(_ Token) any {
	return nil
}

func (p *parser) on_action_push_mode(_, _ Token, id Token, _ Token) any {
	return nil
}

func (p *parser) on_action_pop_mode(_ Token) any {
	return nil
}

func (p *parser) onError() {
	p.errs.Errorf(
		p.file.Position(p.errorToken().Pos),
		"unexpected %v", p.errorToken())
}
