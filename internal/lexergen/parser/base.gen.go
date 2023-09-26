package parser

type TokenType int

const (
	EOF             TokenType = 0
	ERROR           TokenType = 1
	ID              TokenType = 2
	LITERAL         TokenType = 3
	CLASS_CHAR      TokenType = 4
	SEMICOLON       TokenType = 5
	EQ              TokenType = 6
	OR              TokenType = 7
	ARROW           TokenType = 8
	OCURLY          TokenType = 9
	CCURLY          TokenType = 10
	DASH            TokenType = 11
	HAT             TokenType = 12
	TILDE           TokenType = 13
	OBRACKET        TokenType = 14
	CBRACKET        TokenType = 15
	OPAREN          TokenType = 16
	CPAREN          TokenType = 17
	ZERO_OR_ONE     TokenType = 18
	ZERO_OR_MORE    TokenType = 19
	ONE_OR_MORE     TokenType = 20
	ZERO_OR_MORE_NG TokenType = 21
	ONE_OR_MORE_NG  TokenType = 22
	SKIP            TokenType = 23
	MACRO           TokenType = 24
	FRAG            TokenType = 25
	MODE            TokenType = 26
	PUSH_MODE       TokenType = 27
	POP_MODE        TokenType = 28
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
	case SEMICOLON:
		return ";"
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
	case DASH:
		return "-"
	case HAT:
		return "^"
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
