package main


const (
	EOF int = 0
	ERROR int = 1
	NUM int = 2
	ADD int = 3
	SUB int = 4
	MUL int = 5
	DIV int = 6
	REM int = 7
	POW int = 8
	O_PAREN int = 9
	C_PAREN int = 10
)

func _TokenToString(t int) string {
	switch t {
	case EOF: 
		return "EOF"
	case ERROR: 
		return "ERROR"
	case NUM: 
		return "NUM"
	case ADD: 
		return "ADD"
	case SUB: 
		return "SUB"
	case MUL: 
		return "MUL"
	case DIV: 
		return "DIV"
	case REM: 
		return "REM"
	case POW: 
		return "POW"
	case O_PAREN: 
		return "O_PAREN"
	case C_PAREN: 
		return "C_PAREN"
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
