package token

import "github.com/dcaiafa/lox/internal/loc"

type Type int

const (
	EOF Type = iota
	ID
	LITERAL
	LABEL
	ZERO_OR_MANY
	ONE_OR_MANY
	ZERO_OR_ONE
	DEFINE
	OR
	SEMICOLON
	PARSER
	LEXER
	CUSTOM
)

var names = []string{
	"EOF",
	"ID",
	"LITERAL",
	"LABEL",
	"ZERO_OR_MANY",
	"ONE_OR_MANY",
	"ZERO_OR_ONE",
	"DEFINE",
	"OR",
	"SEMICOLON",
	"PARSER",
	"LEXER",
	"CUSTOM",
}

type Token struct {
	Pos  loc.Loc
	Type Type
	Str  string
}
