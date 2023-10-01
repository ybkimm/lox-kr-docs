package parser

type TokenType int

const (
	EOF             TokenType = 0
	ERROR           TokenType = 1
	ID              TokenType = 2
	LITERAL         TokenType = 3
	CLASS_CHAR      TokenType = 4
	NUM             TokenType = 5
	SEMICOLON       TokenType = 6
	COMMA           TokenType = 7
	EQ              TokenType = 8
	OR              TokenType = 9
	ARROW           TokenType = 10
	OCURLY          TokenType = 11
	CCURLY          TokenType = 12
	CLASS_DASH      TokenType = 13
	TILDE           TokenType = 14
	OBRACKET        TokenType = 15
	CBRACKET        TokenType = 16
	OPAREN          TokenType = 17
	CPAREN          TokenType = 18
	ZERO_OR_ONE     TokenType = 19
	ZERO_OR_MORE    TokenType = 20
	ONE_OR_MORE     TokenType = 21
	ZERO_OR_MORE_NG TokenType = 22
	ONE_OR_MORE_NG  TokenType = 23
	PARSER          TokenType = 24
	LEXER           TokenType = 25
	SKIP            TokenType = 26
	MACRO           TokenType = 27
	FRAG            TokenType = 28
	MODE            TokenType = 29
	PUSH_MODE       TokenType = 30
	POP_MODE        TokenType = 31
	ERROR_KEYWORD   TokenType = 32
	LEFT            TokenType = 33
	LIST            TokenType = 34
	RIGHT           TokenType = 35
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "end-of-file"
	case ERROR:
		return "error"
	case ID:
		return "ID"
	case LITERAL:
		return "LITERAL"
	case CLASS_CHAR:
		return "CLASS_CHAR"
	case NUM:
		return "NUM"
	case SEMICOLON:
		return ";"
	case COMMA:
		return ","
	case EQ:
		return "="
	case OR:
		return "|"
	case ARROW:
		return "->"
	case OCURLY:
		return "{"
	case CCURLY:
		return "}"
	case CLASS_DASH:
		return "-"
	case TILDE:
		return "~"
	case OBRACKET:
		return "["
	case CBRACKET:
		return "]"
	case OPAREN:
		return "("
	case CPAREN:
		return ")"
	case ZERO_OR_ONE:
		return "?"
	case ZERO_OR_MORE:
		return "*"
	case ONE_OR_MORE:
		return "+"
	case ZERO_OR_MORE_NG:
		return "*?"
	case ONE_OR_MORE_NG:
		return "+?"
	case PARSER:
		return "@parser"
	case LEXER:
		return "@lexer"
	case SKIP:
		return "@skip"
	case MACRO:
		return "@macro"
	case FRAG:
		return "@frag"
	case MODE:
		return "@mode"
	case PUSH_MODE:
		return "@push_mode"
	case POP_MODE:
		return "@pop_mode"
	case ERROR_KEYWORD:
		return "@error"
	case LEFT:
		return "@left"
	case LIST:
		return "@list"
	case RIGHT:
		return "@right"
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
