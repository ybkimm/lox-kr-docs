package token

import (
	gotoken "go/token"
)

type Token struct {
	Pos  gotoken.Pos
	Type int
	Str  string
}

func (t Token) ID() int {
	return int(t.Type)
}
