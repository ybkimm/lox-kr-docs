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
	CLASS_DASH      TokenType = 11
	TILDE           TokenType = 12
	OBRACKET        TokenType = 13
	CBRACKET        TokenType = 14
	OPAREN          TokenType = 15
	CPAREN          TokenType = 16
	ZERO_OR_ONE     TokenType = 17
	ZERO_OR_MORE    TokenType = 18
	ONE_OR_MORE     TokenType = 19
	ZERO_OR_MORE_NG TokenType = 20
	ONE_OR_MORE_NG  TokenType = 21
	SKIP            TokenType = 22
	MACRO           TokenType = 23
	FRAG            TokenType = 24
	MODE            TokenType = 25
	PUSH_MODE       TokenType = 26
	POP_MODE        TokenType = 27
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
