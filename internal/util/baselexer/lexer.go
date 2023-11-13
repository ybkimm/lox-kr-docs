package baselexer

import (
	"bytes"
	gotoken "go/token"
)

const EOF = 0
const ERROR = 1

type StateMachine interface {
	PushRune(r rune) int
	Token() int
	Reset()
}

type Token struct {
	Type int
	Str  []byte
	Pos  gotoken.Pos
}

type Lexer struct {
	sm          StateMachine
	onError     func(l *Lexer)
	file        *gotoken.File
	input       []byte
	inputReader *bytes.Reader
	offset      int
	charLen     int
	char        rune
	start       int
	pos         gotoken.Pos
}

func New(
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

func (l *Lexer) consume() {
	if l.char == '\n' {
		l.file.AddLine(l.offset + 1)
	}
	l.offset += l.charLen
	r, charLen, err := l.inputReader.ReadRune()
	if err != nil {
		l.char = 0
		l.charLen = 0
		return
	}
	l.char = r
	l.charLen = charLen
}

func (l *Lexer) Pos() gotoken.Pos {
	return l.pos
}

func (l *Lexer) Peek() rune {
	return l.char
}

func (l *Lexer) ReadToken() (Token, int) {
	l.start = -1
	l.pos = 0

	for {
		if l.start == -1 {
			l.start = l.offset
			l.pos = l.file.Pos(l.offset)
		}

		action := l.sm.PushRune(l.char)

		switch action {
		case 0: // consume
			l.consume()

		case 1: // accept
			end := l.offset
			t := Token{
				Type: l.sm.Token(),
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

			// Read until the beginning of the next line.
			for l.char != '\n' && l.char != 0 {
				l.consume()
			}
			l.consume()

			l.sm.Reset()

			return t, t.Type
		}
	}
}
