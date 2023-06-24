package token

type Type int

const (
	Undefined Type = iota
	ID
	Literal
	Eq
	Period
	VBar
	Star
	Plus
	QMark
	Hash
)

type Token struct {
	Type Type
	Str  string
}
