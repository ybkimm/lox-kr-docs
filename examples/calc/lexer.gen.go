package main

import (
  _i0 "fmt"
)

type TokenType int

const (
	EOF TokenType = 0
	ERROR TokenType = 1
	NUM TokenType = 2
	PLUS TokenType = 3
	MINUS TokenType = 4
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
	case PLUS: 
		return "PLUS"
	case MINUS: 
		return "MINUS"
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
