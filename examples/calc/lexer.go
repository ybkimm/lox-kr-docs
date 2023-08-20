package main

import (
	"bytes"
	gotoken "go/token"
	"io"
)

type Token struct {
	Type TokenType
	Str  string
	Pos  gotoken.Pos
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
	errLogger *ErrLogger
	buf       bytes.Buffer
	char      rune
}

func newLex(file *gotoken.File, input []byte, errLogger *ErrLogger) *lex {
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

func (l *lex) offset() int {
	offset, _ := l.input.Seek(0, io.SeekCurrent)
	return int(offset) - 1
}

func (l *lex) advance() {
	if l.char == '\n' {
		// Register the new line so that go/token.Pos can be converted to
		// line/col. The line starts at the character after the \n.
		l.file.AddLine(l.offset() + 1)
	}
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
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
		case '+':
			tok.Type = PLUS
			l.advance()
		case '-':
			tok.Type = MINUS
			l.advance()
		case '*':
			tok.Type = MUL
			l.advance()
		case '/':
			tok.Type = DIV
			l.advance()
		case '%':
			tok.Type = REM
			l.advance()
		case '^':
			tok.Type = POW
			l.advance()
		case '(':
			tok.Type = O_PAREN
			l.advance()
		case ')':
			tok.Type = C_PAREN
			l.advance()
		default:
			if isNumber(r) {
				l.scanNum(tok)
			} else {
				l.errLogger.Errorf(
					l.pos(),
					"unexpected character: %c", r)
				tok.Type = ERROR
			}
		}
	}
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
