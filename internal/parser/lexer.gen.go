package parser

import (
  _i0 "fmt"
)

type TokenType int

const (
	EOF TokenType = 0
	ERROR TokenType = 1
	ID TokenType = 2
	LITERAL TokenType = 3
	NUM TokenType = 4
	ZERO_OR_MANY TokenType = 5
	ONE_OR_MANY TokenType = 6
	ZERO_OR_ONE TokenType = 7
	DEFINE TokenType = 8
	OR TokenType = 9
	SEMICOLON TokenType = 10
	OPAREN TokenType = 11
	CPAREN TokenType = 12
	PARSER TokenType = 13
	LEXER TokenType = 14
	TOKEN TokenType = 15
	LEFT TokenType = 16
	RIGHT TokenType = 17
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

type _lxErrorLogger interface {
	ParserError(err error)
}

type _lxUnexpectedTokenError struct {
	Token Token
}

func (e _lxUnexpectedTokenError) Error() string {
	return _i0.Sprintf("unexpected token: %v", e.Token)
}

func (e _lxUnexpectedTokenError) Pos() Token {
	return e.Token
}

type _lxLexer interface {
	NextToken() (Token, TokenType)
}
