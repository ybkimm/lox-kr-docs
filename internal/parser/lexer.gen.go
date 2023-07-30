package parser

import (
  _i0 "go/token"
)

type TokenType int

const (
	EOF TokenType = 0
	ID TokenType = 1
	LITERAL TokenType = 2
	NUM TokenType = 3
	ZERO_OR_MANY TokenType = 4
	ONE_OR_MANY TokenType = 5
	ZERO_OR_ONE TokenType = 6
	DEFINE TokenType = 7
	OR TokenType = 8
	SEMICOLON TokenType = 9
	OPAREN TokenType = 10
	CPAREN TokenType = 11
	PARSER TokenType = 12
	LEXER TokenType = 13
	TOKEN TokenType = 14
	LEFT TokenType = 15
	RIGHT TokenType = 16
)

func (t TokenType) String() string {
	switch t {
	case EOF: 
		return "EOF"
	case ID: 
		return "ID"
	case LITERAL: 
		return "LITERAL"
	case NUM: 
		return "NUM"
	case ZERO_OR_MANY: 
		return "ZERO_OR_MANY"
	case ONE_OR_MANY: 
		return "ONE_OR_MANY"
	case ZERO_OR_ONE: 
		return "ZERO_OR_ONE"
	case DEFINE: 
		return "DEFINE"
	case OR: 
		return "OR"
	case SEMICOLON: 
		return "SEMICOLON"
	case OPAREN: 
		return "OPAREN"
	case CPAREN: 
		return "CPAREN"
	case PARSER: 
		return "PARSER"
	case LEXER: 
		return "LEXER"
	case TOKEN: 
		return "TOKEN"
	case LEFT: 
		return "LEFT"
	case RIGHT: 
		return "RIGHT"
	default:
		return "???"
	}
}

type Token struct {
	Pos  _i0.Pos
	Type TokenType
	Str  string
}
