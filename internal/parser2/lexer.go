package parser

import (
	"bytes"
	gotoken "go/token"
	"io"
	"strconv"

	"github.com/dcaiafa/lox/internal/errlogger"
)

type Token struct {
	Type TokenType
	Str  string
	Pos  gotoken.Pos
}

func (t Token) String() string {
	switch t.Type {
	case LITERAL, CLASS_CHAR:
		return t.Type.String() + ` '` + t.Str + `'`
	case ID:
		return t.Type.String() + ` "` + t.Str + `"`
	default:
		return t.Type.String()
	}
}

var keywords = map[string]TokenType{
	"@frag":      FRAG,
	"@lexer":     LEXER,
	"@macro":     MACRO,
	"@mode":      MODE,
	"@parser":    PARSER,
	"@pop_mode":  POP_MODE,
	"@push_mode": PUSH_MODE,
	"@skip":      SKIP,
	"@list":      LIST,
	"@left":      LEFT,
	"@right":     RIGHT,
	"@start":     START,
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isHex(r rune) bool {
	return isNumber(r) || (r >= 'A' && r <= 'F') || (r >= 'a' && r <= 'f')
}

func isLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

type modeFunc func()

type lex struct {
	file      *gotoken.File
	input     *bytes.Reader
	errLogger *errlogger.ErrLogger
	buf       bytes.Buffer
	char      rune
	mode      modeFunc
	tok       Token
}

func newLex(file *gotoken.File, input []byte, errLogger *errlogger.ErrLogger) *lex {
	l := &lex{
		file:      file,
		input:     bytes.NewReader(input),
		errLogger: errLogger,
	}
	l.mode = l.modeDefault
	l.advance()
	return l
}

func (l *lex) ReadToken() (Token, TokenType) {
	l.readToken()
	return l.tok, l.tok.Type
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

func (l *lex) recover() {
	for {
		l.advance()
		ch := l.peek()
		if ch == ' ' || ch == '\n' || ch == 0 {
			break
		}
	}
	l.mode = l.modeDefault
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

func (l *lex) readToken() {
	l.tok.Type = -1
	for l.tok.Type == -1 {
		r := l.peek()
		l.tok.Pos = l.pos()
		if r == 0 {
			l.tok.Type = EOF
			return
		}
		l.mode()
	}
}

func (l *lex) modeDefault() {
	r := l.peek()
	switch r {
	case '/':
		l.scanComment()
	case '=':
		l.advance()
		l.tok.Type = EQ
	case ',':
		l.advance()
		l.tok.Type = COMMA
	case ';':
		l.advance()
		l.tok.Type = SEMICOLON
	case '|':
		l.advance()
		l.tok.Type = OR
	case '*':
		l.advance()
		l.tok.Type = ZERO_OR_MORE
	case '+':
		l.advance()
		l.tok.Type = ONE_OR_MORE
	case '?':
		l.advance()
		l.tok.Type = ZERO_OR_ONE
	case '(':
		l.advance()
		l.tok.Type = OPAREN
	case ')':
		l.advance()
		l.tok.Type = CPAREN
	case '{':
		l.advance()
		l.tok.Type = OCURLY
	case '}':
		l.advance()
		l.tok.Type = CCURLY
	case '~':
		l.advance()
		l.tok.Type = TILDE
	case '-':
		l.advance()
		if l.peek() == '>' {
			l.advance()
			l.tok.Type = ARROW
		} else {
			l.unexpectedChar()
		}
	case '[':
		l.advance()
		l.tok.Type = OBRACKET
		l.mode = l.modeCharClass
	case '\'':
		l.scanLiteral()
	case '@':
		l.scanKeyword()
	default:
		if isSpace(r) || r == '\n' {
			l.advance()
		} else if isLetter(r) || r == '_' {
			l.scanIdentifier()
		} else if isNumber(r) {
			l.scanNumber()
		} else {
			l.unexpectedChar()
		}
	}
}

func (l *lex) modeCharClass() {
	r := l.peek()
	switch r {
	case ']':
		l.advance()
		l.tok.Type = CBRACKET
		l.mode = l.modeDefault
	case '\\':
		l.buf.Reset()
		l.consumeEscapedChar()
		l.tok.Type = CLASS_CHAR
		l.tok.Str = l.buf.String()
	case '-':
		// Parser expects DASH to have a Str value.
		l.buf.Reset()
		l.buf.WriteRune('-')
		l.advance()
		l.tok.Type = CLASS_DASH
		l.tok.Str = l.buf.String()
	case '\r', '\n', '\t':
		l.unexpectedChar()
	default:
		l.buf.Reset()
		l.buf.WriteRune(r)
		l.advance()
		l.tok.Type = CLASS_CHAR
		l.tok.Str = l.buf.String()
	}
}

func (l *lex) consumeEscapedChar() {
	r := l.peek()
	if r != '\\' {
		panic("not-reached")
	}

	l.advance()
	switch l.peek() {
	case 'n':
		l.buf.WriteRune('\n')
		l.advance()
	case 'r':
		l.buf.WriteRune('\r')
		l.advance()
	case 't':
		l.buf.WriteRune('\t')
		l.advance()
	case '-':
		l.buf.WriteRune('-')
		l.advance()
	case '\\':
		l.buf.WriteRune('\\')
		l.advance()
	case 'u':
		l.advance()
		var buf [4]byte
		for i := 0; i < 4; i++ {
			r := l.peek()
			if !isHex(r) {
				l.unexpectedChar()
				return
			}
			buf[i] = byte(r)
			l.advance()
		}
		ru, _ := strconv.ParseUint(string(buf[:]), 16, 16)
		l.buf.WriteRune(rune(ru))
	default:
		l.unexpectedChar()
	}
}

func (l *lex) scanComment() {
	l.advance()
	if l.peek() != '/' {
		l.errLogger.Errorf(
			l.file.Position(l.pos()), "unexpected character: %v", l.peek())
		l.recover()
		l.tok.Type = ERROR
		return
	}
	for l.peek() != '\n' {
		l.advance()
	}
}

func (l *lex) scanLiteral() {
	l.buf.Reset()

	// Skip the first '
	l.advance()

	for {
		r := l.peek()
		switch r {
		case '\'':
			l.advance()
			l.tok.Str = l.buf.String()
			l.tok.Type = LITERAL
			return

		case '\\':
			l.advance()
			r = l.peek()
			switch r {
			case '\'', '\\':
				l.buf.WriteRune(r)
				l.advance()
			case 'n':
				l.buf.WriteRune('\n')
				l.advance()
			default:
				l.errLogger.Errorf(
					l.file.Position(l.pos()), "unexpected character %v in string literal", r)
				l.recover()
				l.tok.Type = ERROR
				return
			}

		case '\n':
			l.errLogger.Errorf(
				l.file.Position(l.pos()), "newline in string literal")
			// Don't recover. We will start parsing from the '\n'.
			l.tok.Type = ERROR
			return

		default:
			l.buf.WriteRune(r)
			l.advance()
		}
	}
}

func (l *lex) scanIdentifier() {
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
	l.tok.Type = ID
	l.tok.Str = l.buf.String()
}

func (l *lex) scanKeyword() {
	l.buf.Reset()

	r := l.peek()
	l.advance()
	l.buf.WriteRune(r)

	for {
		r := l.peek()
		if !isLetter(r) && r != '_' {
			break
		}
		l.advance()
		l.buf.WriteRune(r)
	}

	tokStr := l.buf.String()
	keyword, ok := keywords[tokStr]
	if !ok {
		l.errLogger.Errorf(l.file.Position(l.tok.Pos), "invalid keyword %v", tokStr)
		l.tok.Type = ERROR
		l.recover()
		return
	}
	l.tok.Type = keyword
	l.tok.Str = l.buf.String()
}

func (l *lex) scanNumber() {
	l.buf.Reset()
	for isNumber(l.peek()) {
		l.buf.WriteRune(l.peek())
		l.advance()
	}
	l.tok.Type = NUM
	l.tok.Str = l.buf.String()
}

func (l *lex) unexpectedChar() {
	l.errLogger.Errorf(
		l.file.Position(l.pos()), "unexpected character: %c", l.peek())
	l.recover()
	l.tok.Type = ERROR
}
