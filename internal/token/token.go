package token

import "github.com/dcaiafa/lox/internal/fileloc"

type Type int

const (
	Undefined Type = iota
	ID
	Literal
	Keyword
)

type Token struct {
	Pos  fileloc.FileLoc
	Type Type
	Str  string
}
