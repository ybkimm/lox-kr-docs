package parser

import (
  _i0 "go/token"
)

type TokenType int

const (
	EOF TokenType = 0
	ID TokenType = 1
	LITERAL TokenType = 2
	LABEL TokenType = 3
	ZERO_OR_MANY TokenType = 4
	ONE_OR_MANY TokenType = 5
	ZERO_OR_ONE TokenType = 6
	DEFINE TokenType = 7
	OR TokenType = 8
	SEMICOLON TokenType = 9
	PARSER TokenType = 10
	LEXER TokenType = 11
	TOKEN TokenType = 12
)

func (t TokenType) String() string {
	switch t {
	case EOF: 
		return "EOF"
	case ID: 
		return "ID"
	case LITERAL: 
		return "LITERAL"
	case LABEL: 
		return "LABEL"
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
	case PARSER: 
		return "PARSER"
	case LEXER: 
		return "LEXER"
	case TOKEN: 
		return "TOKEN"
	default:
		return "???"
	}
}

type Token struct {
	Pos  _i0.Pos
	Type TokenType
	Str  string
}
