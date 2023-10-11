package parser


type TokenType int

const (
	EOF TokenType = 0
	ERROR TokenType = 1
	ID TokenType = 2
	LITERAL TokenType = 3
	CLASS_CHAR TokenType = 4
	NUM TokenType = 5
	SEMICOLON TokenType = 6
	COMMA TokenType = 7
	EQ TokenType = 8
	OR TokenType = 9
	ARROW TokenType = 10
	OCURLY TokenType = 11
	CCURLY TokenType = 12
	CLASS_DASH TokenType = 13
	TILDE TokenType = 14
	OBRACKET TokenType = 15
	CBRACKET TokenType = 16
	OPAREN TokenType = 17
	CPAREN TokenType = 18
	ZERO_OR_ONE TokenType = 19
	ZERO_OR_MORE TokenType = 20
	ONE_OR_MORE TokenType = 21
	PARSER TokenType = 22
	LEXER TokenType = 23
	START TokenType = 24
	SKIP TokenType = 25
	MACRO TokenType = 26
	FRAG TokenType = 27
	MODE TokenType = 28
	PUSH_MODE TokenType = 29
	POP_MODE TokenType = 30
	ERROR_KEYWORD TokenType = 31
	LEFT TokenType = 32
	LIST TokenType = 33
	RIGHT TokenType = 34
)

func (t TokenType) String() string {
	switch t {
	case EOF: 
		return "EOF"
	case ERROR: 
		return "ERROR"
	case ID: 
		return "ID"
	case LITERAL: 
		return "LITERAL"
	case CLASS_CHAR: 
		return "CLASS_CHAR"
	case NUM: 
		return "NUM"
	case SEMICOLON: 
		return "SEMICOLON"
	case COMMA: 
		return "COMMA"
	case EQ: 
		return "EQ"
	case OR: 
		return "OR"
	case ARROW: 
		return "ARROW"
	case OCURLY: 
		return "OCURLY"
	case CCURLY: 
		return "CCURLY"
	case CLASS_DASH: 
		return "CLASS_DASH"
	case TILDE: 
		return "TILDE"
	case OBRACKET: 
		return "OBRACKET"
	case CBRACKET: 
		return "CBRACKET"
	case OPAREN: 
		return "OPAREN"
	case CPAREN: 
		return "CPAREN"
	case ZERO_OR_ONE: 
		return "ZERO_OR_ONE"
	case ZERO_OR_MORE: 
		return "ZERO_OR_MORE"
	case ONE_OR_MORE: 
		return "ONE_OR_MORE"
	case PARSER: 
		return "PARSER"
	case LEXER: 
		return "LEXER"
	case START: 
		return "START"
	case SKIP: 
		return "SKIP"
	case MACRO: 
		return "MACRO"
	case FRAG: 
		return "FRAG"
	case MODE: 
		return "MODE"
	case PUSH_MODE: 
		return "PUSH_MODE"
	case POP_MODE: 
		return "POP_MODE"
	case ERROR_KEYWORD: 
		return "ERROR_KEYWORD"
	case LEFT: 
		return "LEFT"
	case LIST: 
		return "LIST"
	case RIGHT: 
		return "RIGHT"
	default:
		return "???"
	}
}
