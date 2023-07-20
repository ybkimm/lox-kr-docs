package parser2

import (
	"bytes"
	"fmt"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/loc"
	"github.com/dcaiafa/lox/internal/token"
)

var keywords = map[string]token.Type{
	"@lexer":  token.LEXER,
	"@parser": token.PARSER,
	"@custom": token.CUSTOM,
}

type lex struct {
	Spec *ast.Spec

	char    rune
	input   *bytes.Reader
	errs    *errs.Errs
	pos     loc.Loc
	lastPos loc.Loc
	buf     bytes.Buffer
}

func newLex(filename string, input []byte, errs *errs.Errs) *lex {
	l := &lex{
		input: bytes.NewReader(input),
		errs:  errs,
		pos: loc.Loc{
			Filename: filename,
			Line:     1,
			Column:   1,
		},
	}
	l.advance()
	return l
}

func (l *lex) NextToken() (token.Token, error) {
	var tok token.Token
	err := l.nextToken(&tok)
	if tok.Pos.Line != 0 {
		l.lastPos = tok.Pos
	}
	return tok, err
}

func (l *lex) nextToken(tok *token.Token) error {
	for {
		r := l.peek()
		if r == 0 {
			tok.Type = token.EOF
			return nil
		}
		if isSpace(r) {
			l.advance()
			continue
		}
		if r == '\n' {
			l.advance()
			continue
		}

		tok.Pos = l.pos

		switch r {
		case '=':
			l.advance()
			tok.Type = token.DEFINE
		case ';':
			l.advance()
			tok.Type = token.SEMICOLON
		case '|':
			l.advance()
			tok.Type = token.OR
		case '*':
			l.advance()
			tok.Type = token.ZERO_OR_MANY
		case '+':
			l.advance()
			tok.Type = token.ONE_OR_MANY
		case '?':
			l.advance()
			tok.Type = token.ZERO_OR_ONE
		case '@':
			err := l.scanKeyword(tok)
			if err != nil {
				return err
			}
		default:
			if isLetter(r) || r == '_' {
				l.scanIdentifier(tok)
			} else {
				return fmt.Errorf("unexpected character: %v", r)
			}
		}
		return nil
	}
}

func (l *lex) scanIdentifier(tok *token.Token) {
	l.buf.Reset()

	r := l.peek()
	l.advance()
	l.buf.WriteRune(r)

	for {
		r := l.peek()
		if !isLetter(r) && !isNumber(r) && r != '_' {
			break
		}
		l.advance()
		l.buf.WriteRune(r)
	}
	tok.Type = token.ID
	tok.Str = l.buf.String()
}

func (l *lex) scanKeyword(tok *token.Token) error {
	l.buf.Reset()

	r := l.peek()
	l.advance()
	l.buf.WriteRune(r)

	for {
		r := l.peek()
		if !isLetter(r) {
			break
		}
		l.advance()
		l.buf.WriteRune(r)
	}

	tokStr := l.buf.String()
	keyword, ok := keywords[tokStr]
	if !ok {
		return fmt.Errorf("invalid keyword %v", tokStr)
	}
	tok.Type = keyword
	tok.Str = l.buf.String()
	return nil
}

func (l *lex) peek() rune {
	return l.char
}

func (l *lex) advance() {
	if l.char == '\n' {
		l.pos.Column = 1
		l.pos.Line++
	} else if l.char != 0 {
		l.pos.Column++
	}
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
}

func (l *lex) Error(s string) {
	l.errs.Errorf(l.lastPos, "%v", s)
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
