package parser

const (
	EOF             int = 0
	ERROR           int = 1
	SEMICOLON       int = 2
	COMMA           int = 3
	EQ              int = 4
	OR              int = 5
	OCURLY          int = 6
	CCURLY          int = 7
	TILDE           int = 8
	OPAREN          int = 9
	CPAREN          int = 10
	SUB             int = 11
	ZERO_OR_ONE     int = 12
	ZERO_OR_MORE    int = 13
	ONE_OR_MORE     int = 14
	ZERO_OR_MORE_NG int = 15
	ONE_OR_MORE_NG  int = 16
	PARSER          int = 17
	LEXER           int = 18
	START           int = 19
	DISCARD         int = 20
	MACRO           int = 21
	FRAG            int = 22
	MODE            int = 23
	PUSH_MODE       int = 24
	POP_MODE        int = 25
	ERROR_KEYWORD   int = 26
	LEFT            int = 27
	LIST            int = 28
	RIGHT           int = 29
	EMIT            int = 30
	EMPTY           int = 31
	KEYWORD         int = 32
	ID              int = 33
	NUM             int = 34
	LITERAL         int = 35
	OBRACKET        int = 36
	CBRACKET        int = 37
	CLASS_DASH      int = 38
	CLASS_CHAR      int = 39
)

func _TokenToString(t int) string {
	switch t {
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	case SEMICOLON:
		return "SEMICOLON"
	case COMMA:
		return "COMMA"
	case EQ:
		return "EQ"
	case OR:
		return "OR"
	case OCURLY:
		return "OCURLY"
	case CCURLY:
		return "CCURLY"
	case TILDE:
		return "TILDE"
	case OPAREN:
		return "OPAREN"
	case CPAREN:
		return "CPAREN"
	case SUB:
		return "SUB"
	case ZERO_OR_ONE:
		return "ZERO_OR_ONE"
	case ZERO_OR_MORE:
		return "ZERO_OR_MORE"
	case ONE_OR_MORE:
		return "ONE_OR_MORE"
	case ZERO_OR_MORE_NG:
		return "ZERO_OR_MORE_NG"
	case ONE_OR_MORE_NG:
		return "ONE_OR_MORE_NG"
	case PARSER:
		return "PARSER"
	case LEXER:
		return "LEXER"
	case START:
		return "START"
	case DISCARD:
		return "DISCARD"
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
	case EMIT:
		return "EMIT"
	case EMPTY:
		return "EMPTY"
	case KEYWORD:
		return "KEYWORD"
	case ID:
		return "ID"
	case NUM:
		return "NUM"
	case LITERAL:
		return "LITERAL"
	case OBRACKET:
		return "OBRACKET"
	case CBRACKET:
		return "CBRACKET"
	case CLASS_DASH:
		return "CLASS_DASH"
	case CLASS_CHAR:
		return "CLASS_CHAR"
	default:
		return "???"
	}
}

type _Stack[T any] []T

func (s *_Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}

func (s _Stack[T]) PeekSlice(n int) []T {
	return s[len(s)-n:]
}
