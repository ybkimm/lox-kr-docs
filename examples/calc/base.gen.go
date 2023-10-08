package main


type TokenType int

const (
	EOF TokenType = 0
	ERROR TokenType = 1
	NUM TokenType = 2
	ADD TokenType = 3
	SUB TokenType = 4
	MUL TokenType = 5
	DIV TokenType = 6
	REM TokenType = 7
	POW TokenType = 8
	O_PAREN TokenType = 9
	C_PAREN TokenType = 10
)

func (t TokenType) String() string {
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
