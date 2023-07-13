package parser2

import (
	"bytes"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/loc"
	"github.com/dcaiafa/lox/internal/token"
)

var keywords = map[string]int{
	"@lexer":  kLEXER,
	"@parser": kPARSER,
	"@custom": kCUSTOM,
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

func (l *lex) Lex(lval *yySymType) int {
	return l.scan(lval)
}

func (l *lex) scan(lval *yySymType) (tok int) {
	defer func() {
		if lval.tok.Pos.Line != 0 {
			l.lastPos = lval.tok.Pos
		}
	}()

	lval.tok = token.Token{}

	for {
		r := l.peek()
		if r == 0 {
			return 0
		}
		if isSpace(r) {
			l.advance()
			continue
		}
		if r == '\n' {
			l.advance()
			continue
		}

		lval.tok.Pos = l.pos

		switch r {
		case '\'':
			return l.scanSingleQuotedString(lval)
		case '@':
			return l.scanKeyword(lval)
		case '#':
			return l.scanLabel(lval)
		case '=', ';', '|', '*', '+', '?':
			l.advance()
			return int(r)
		default:
			if isLetter(r) || r == '_' {
				tok := l.scanIdentifier(lval)
				return tok
			} else {
				return LEXERR
			}
		}
	}
}

func (l *lex) scanIdentifier(lval *yySymType) int {
	l.buf.Reset()

	r := l.peek()
	if !isLetter(r) && r != '_' {
		return LEXERR
	}
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

	lval.tok.Str = l.buf.String()
	lval.tok.Type = token.ID
	return ID
}

func (l *lex) scanKeyword(lval *yySymType) int {
	l.buf.Reset()

	r := l.peek()
	if l.peek() != '@' {
		return LEXERR
	}
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

	lval.tok.Type = token.ID
	lval.tok.Str = l.buf.String()

	keyword, ok := keywords[lval.tok.Str]
	if !ok {
		return LEXERR
	}

	return keyword
}

func (l *lex) scanLabel(lval *yySymType) int {
	l.buf.Reset()

	r := l.peek()
	if l.peek() != '#' {
		return LEXERR
	}
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

	lval.tok.Type = token.LABEL
	lval.tok.Str = l.buf.String()

	return LABEL
}

func (l *lex) scanSingleQuotedString(lval *yySymType) int {
	l.buf.Reset()

	if l.peek() != '\'' {
		return LEXERR
	}
	l.advance()

	for {
		r := l.peek()
		if r == 0 {
			return LEXERR
		}
		l.advance()
		if r == '\'' {
			break
		}
		l.buf.WriteRune(r)
	}

	lval.tok.Type = token.LITERAL
	lval.tok.Str = l.buf.String()
	return LITERAL
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
