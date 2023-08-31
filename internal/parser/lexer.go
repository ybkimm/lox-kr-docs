package parser

import (
	"bytes"
	gotoken "go/token"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
)

type Token struct {
	Type TokenType
	Str  string
	Pos  gotoken.Pos
}

var keywords = map[string]TokenType{
	"@left":  LEFT,
	"@list":  LIST,
	"@right": RIGHT,
	"@token": TOKEN,
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

type lex struct {
	file      *gotoken.File
	input     *bytes.Reader
	errLogger *errlogger.ErrLogger
	buf       bytes.Buffer
	char      rune
}

func newLex(file *gotoken.File, input []byte, errLogger *errlogger.ErrLogger) *lex {
	l := &lex{
		file:      file,
		input:     bytes.NewReader(input),
		errLogger: errLogger,
	}
	l.advance()
	return l
}

func (l *lex) NextToken() (Token, TokenType) {
	var tok Token
	l.nextToken(&tok)
	return tok, tok.Type
}

func (l *lex) advance() {
	if l.char == '\n' {
		// Line starts at the character after the \n.
		l.file.AddLine(l.offset() + 1)
	}
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
}

func (l *lex) offset() int {
	offset, _ := l.input.Seek(0, io.SeekCurrent)
	return int(offset) - 1
}

func (l *lex) pos() gotoken.Pos {
	return l.file.Pos(l.offset())
}

func (l *lex) peek() rune {
	return l.char
}

func (l *lex) nextToken(tok *Token) {
	tok.Type = -1
	for tok.Type == -1 {
		r := l.peek()
		if r == 0 {
			tok.Type = EOF
			return
		}
		if isSpace(r) {
			l.advance()
			continue
		}
		if r == '\n' {
			l.advance()
			continue
		}

		tok.Pos = l.pos()

		switch r {
		case '/':
			l.scanComment(tok)
		case '=':
			l.advance()
			tok.Type = DEFINE
		case ';':
			l.advance()
			tok.Type = SEMICOLON
		case '|':
			l.advance()
			tok.Type = OR
		case '*':
			l.advance()
			tok.Type = ZERO_OR_MANY
		case '+':
			l.advance()
			tok.Type = ONE_OR_MANY
		case '?':
			l.advance()
			tok.Type = ZERO_OR_ONE
		case '(':
			l.advance()
			tok.Type = OPAREN
		case ')':
			l.advance()
			tok.Type = CPAREN
		case ',':
			l.advance()
			tok.Type = COMMA
		case '\'':
			l.scanLiteral(tok)
		case '@':
			l.scanKeyword(tok)
		default:
			if isLetter(r) || r == '_' {
				l.scanIdentifier(tok)
			} else if isNumber(r) {
				l.scanNum(tok)
			} else {
				l.errLogger.Errorf(
					l.file.Position(l.pos()), "unexpected character: %v", r)
				tok.Type = ERROR
			}
		}
	}
}

func (l *lex) scanComment(tok *Token) {
	l.advance()
	if l.peek() != '/' {
		l.errLogger.Errorf(
			l.file.Position(l.pos()), "unexpected character: %v", l.peek())
		tok.Type = ERROR
		return
	}
	for l.peek() != '\n' {
		l.advance()
	}
}

func (l *lex) scanLiteral(tok *Token) {
	l.buf.Reset()

	// Skip the first '.
	l.advance()

	for {
		r := l.peek()
		switch r {
		case '\'':
			l.advance()
			tok.Str = l.buf.String()
			tok.Type = LITERAL
			return

		case '\\':
			l.advance()
			r = l.peek()
			switch r {
			case '\'', '\\':
				l.buf.WriteRune(r)
				l.advance()
			default:
				l.errLogger.Errorf(
					l.file.Position(l.pos()), "unexpected character %v in string literal", r)
				tok.Type = ERROR
				return
			}

		case '\n':
			l.errLogger.Errorf(
				l.file.Position(l.pos()), "newline in string literal")
			tok.Type = ERROR
			return

		default:
			l.buf.WriteRune(r)
			l.advance()
		}
	}
}

func (l *lex) scanIdentifier(tok *Token) {
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
	tok.Type = ID
	tok.Str = l.buf.String()
}

func (l *lex) scanNum(tok *Token) {
	l.buf.Reset()

	for isNumber(l.peek()) {
		l.buf.WriteRune(l.peek())
		l.advance()
	}

	tok.Type = NUM
	tok.Str = l.buf.String()
}

func (l *lex) scanKeyword(tok *Token) {
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
		l.errLogger.Errorf(l.file.Position(tok.Pos), "invalid keyword %v", tokStr)
		tok.Type = ERROR
		return
	}
	tok.Type = keyword
	tok.Str = l.buf.String()
}
