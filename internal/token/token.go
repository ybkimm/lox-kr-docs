package token

import "github.com/dcaiafa/lox/internal/loc"

type Type int

const (
	Undefined Type = iota
	ID
	Literal
	Keyword
)

type Token struct {
	Pos  loc.Loc
	Type Type
	Str  string
}
