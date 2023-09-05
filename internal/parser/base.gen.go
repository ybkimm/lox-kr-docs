package parser

type TokenType int

const (
	EOF           TokenType = 0
	ERROR         TokenType = 1
	ID            TokenType = 2
	LITERAL       TokenType = 3
	NUM           TokenType = 4
	ZERO_OR_MANY  TokenType = 5
	ONE_OR_MANY   TokenType = 6
	ZERO_OR_ONE   TokenType = 7
	COMMA         TokenType = 8
	CPAREN        TokenType = 9
	DEFINE        TokenType = 10
	OPAREN        TokenType = 11
	OR            TokenType = 12
	SEMICOLON     TokenType = 13
	ERROR_KEYWORD TokenType = 14
	LEFT          TokenType = 15
	LIST          TokenType = 16
	RIGHT         TokenType = 17
	TOKEN         TokenType = 18
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "end-of-file"
	case ERROR:
		return "error"
	case ID:
		return "identifier"
	case LITERAL:
		return "string"
	case NUM:
		return "number"
	case ZERO_OR_MANY:
		return "*"
	case ONE_OR_MANY:
		return "+"
	case ZERO_OR_ONE:
		return "?"
	case COMMA:
		return ","
	case CPAREN:
		return ")"
	case DEFINE:
		return "="
	case OPAREN:
		return "("
	case OR:
		return "|"
	case SEMICOLON:
		return ";"
	case ERROR_KEYWORD:
		return "@error"
	case LEFT:
		return "@left"
	case LIST:
		return "@list"
	case RIGHT:
		return "@right"
	case TOKEN:
		return "@token"
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
