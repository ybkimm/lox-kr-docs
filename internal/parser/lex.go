package parser

import (
	"bytes"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/fileloc"
	"github.com/dcaiafa/lox/internal/grammar"
	"github.com/dcaiafa/lox/internal/token"
)

var keywords = map[string]int{}

type lex struct {
	Syntax *grammar.Syntax

	input   *bytes.Reader
	errs    *errlogger.ErrLogger
	pos     fileloc.FileLoc
	lastPos fileloc.FileLoc
	buf     bytes.Buffer
}

func newLex(filename string, input []byte, errs *errlogger.ErrLogger) *lex {
	return &lex{
		input: bytes.NewReader(input),
		errs:  errs,
		pos: fileloc.FileLoc{
			Filename: filename,
			Line:     1,
			Column:   1,
		},
	}
}

func (l *lex) Lex(lval *yySymType) int {
	return l.scan(lval)
}

func (l *lex) scan(lval *yySymType) int {
	defer func() {
		if lval.tok.Pos.Line != 0 {
			l.lastPos = lval.tok.Pos
		}
	}()

	lval.tok = token.Token{}

	for {
		r := l.read()
		if r == 0 {
			return EOF
		}
		if isSpace(r) {
			continue
		}
		if r == '\n' {
			continue
		}

		lval.tok.Pos = l.pos

		switch r {
		case '\'':
			l.unread()
			return l.scanSingleQuotedString(lval)
		case '=', '.', '|', '*', '+', '?', '#':
			return int(r)
		default:
			if isLetter(r) || r == '_' {
				l.unread()
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

	r := l.read()
	if !isLetter(r) && r != '_' {
		return LEXERR
	}
	l.buf.WriteRune(r)

	for {
		r := l.read()
		if !isLetter(r) && !isNumber(r) && r != '_' {
			l.unread()
			break
		}
		l.buf.WriteRune(r)
	}

	lval.tok.Str = l.buf.String()

	keyword, ok := keywords[lval.tok.Str]
	if ok {
		lval.tok.Type = token.Keyword
		return keyword
	}

	lval.tok.Type = token.ID
	return ID
}

func (l *lex) scanSingleQuotedString(lval *yySymType) int {
	l.buf.Reset()

	if l.read() != '\'' {
		return LEXERR
	}

	for {
		r := l.read()
		if r == '\'' {
			break
		}
		l.buf.WriteRune(r)
	}

	lval.tok.Type = token.Literal
	lval.tok.Str = l.buf.String()
	return LITERAL
}

func (l *lex) read() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		return 0
	}
	if r == '\n' {
		l.pos.Column = 0
		l.pos.Line++
	} else {
		l.pos.Column++
	}
	return r
}

func (l *lex) unread() {
	l.input.UnreadRune()
	l.pos.Column--
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
