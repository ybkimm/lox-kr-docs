package parser

import (
	"bytes"
	"fmt"
	gotoken "go/token"
	"io"
)

var keywords = map[string]TokenType{
	"@lexer":  LEXER,
	"@parser": PARSER,
	"@token":  TOKEN,
	"@left":   LEFT,
	"@right":  RIGHT,
}

type lex struct {
	file  *gotoken.File
	input *bytes.Reader
	buf   bytes.Buffer
	char  rune
}

func newLex(file *gotoken.File, input []byte) *lex {
	l := &lex{
		file:  file,
		input: bytes.NewReader(input),
	}
	l.advance()
	return l
}

func (l *lex) NextToken() (Token, error) {
	var tok Token
	err := l.nextToken(&tok)
	return tok, err
}

func (l *lex) offset() int {
	offset, _ := l.input.Seek(0, io.SeekCurrent)
	return int(offset) - 1
}

func (l *lex) nextToken(tok *Token) error {
	for {
		r := l.peek()
		if r == 0 {
			tok.Type = EOF
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

		tok.Pos = l.file.Pos(l.offset())

		switch r {
		case '/':
			err := l.scanComment()
			if err != nil {
				return err
			}
			continue
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
		case '@':
			err := l.scanKeyword(tok)
			if err != nil {
				return err
			}
		default:
			if isLetter(r) || r == '_' {
				l.scanIdentifier(tok)
			} else if isNumber(r) {
				l.scanNum(tok)
			} else {
				return fmt.Errorf("unexpected character: %v", r)
			}
		}
		return nil
	}
}

func (l *lex) scanComment() error {
	l.advance()
	if l.peek() != '/' {
		return fmt.Errorf("unexpected character: %v", l.peek())
	}
	for l.peek() != '\n' {
		l.advance()
	}
	return nil
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

func (l *lex) scanKeyword(tok *Token) error {
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
		// The line starts at the character after the \n.
		l.file.AddLine(l.offset() + 1)
	}
	r, _, err := l.input.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
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
