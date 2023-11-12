package main

import (
	"bytes"
	gotoken "go/token"
	"io"
)

type StateMachine interface {
	PushRune(r rune) int
	Token() int
	Reset()
}

type Token struct {
	Type TokenType
	Str  []byte
	Pos  gotoken.Pos
}

type Lexer struct {
	sm          StateMachine
	onError     func(l *Lexer)
	file        *gotoken.File
	input       []byte
	inputReader *bytes.Reader
	char        rune
	start       int
	pos         gotoken.Pos
}

func NewLexer(
	sm StateMachine,
	onError func(l *Lexer),
	file *gotoken.File,
	input []byte,
) *Lexer {
	l := &Lexer{
		sm:          sm,
		onError:     onError,
		file:        file,
		input:       input,
		inputReader: bytes.NewReader(input),
	}
	l.consume()
	return l
}

func (l *Lexer) offset() int {
	offset, _ := l.inputReader.Seek(0, io.SeekCurrent)
	return int(offset) - 1
}

func (l *Lexer) consume() {
	if l.char == '\n' {
		// Line starts at the character after the \n.
		l.file.AddLine(l.offset() + 1)
	}
	r, _, err := l.inputReader.ReadRune()
	if err != nil {
		l.char = 0
		return
	}
	l.char = r
}

func (l *Lexer) Pos() gotoken.Pos {
	return l.pos
}

func (l *Lexer) Peek() rune {
	return l.char
}

func (l *Lexer) ReadToken() (Token, TokenType) {
	l.start = -1
	l.pos = 0

	for {
		if l.start == -1 {
			l.start = l.offset()
			l.pos = l.file.Pos(l.offset())
		}

		action := l.sm.PushRune(l.char)

		switch action {
		case 0: // consume
			l.consume()

		case 1: // accept
			end := l.offset()
			t := Token{
				Type: TokenType(l.sm.Token()),
				Str:  l.input[l.start:end],
				Pos:  l.pos,
			}
			l.start = -1
			return t, t.Type

		case 2: // discard
			l.start = -1

		case 3: // EOF
			t := Token{
				Type: EOF,
				Pos:  l.pos,
			}
			return t, t.Type

		default: // Error
			t := Token{
				Type: ERROR,
				Pos:  l.pos,
			}

			l.onError(l)

			for l.char != '\n' && l.char != 0 {
				l.consume()
			}
			l.sm.Reset()

			return t, t.Type
		}
	}
}
