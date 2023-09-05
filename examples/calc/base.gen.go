package main

type TokenType int

const (
	EOF     TokenType = 0
	ERROR   TokenType = 1
	NUM     TokenType = 2
	PLUS    TokenType = 3
	MINUS   TokenType = 4
	MUL     TokenType = 5
	DIV     TokenType = 6
	REM     TokenType = 7
	POW     TokenType = 8
	O_PAREN TokenType = 9
	C_PAREN TokenType = 10
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "end-of-file"
	case ERROR:
		return "error"
	case NUM:
		return "number"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case REM:
		return "%"
	case POW:
		return "^"
	case O_PAREN:
		return "("
	case C_PAREN:
		return ")"
	default:
		return "???"
	}
}

type _Bounds struct {
	Begin Token
	End   Token
}

type lexer interface {
	ReadToken() (Token, TokenType)
}
