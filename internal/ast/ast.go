package ast

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Section interface {
	isSection()
}

type section struct{}

func (s *section) isSection() {}

type Spec struct {
	Sections []Section
}

type Parser struct {
	section
	Decls []ParserDecl
}

type ParserDecl interface {
	isParserDecl()
}

type parserDecl struct{}

func (d *parserDecl) isParserDecl() {}

type Rule struct {
	parserDecl
	Name  string
	Prods []*Prod
}

type Prod struct {
	Terms []*Term
	Label *Label
}

type Term struct {
	Name      string
	Literal   string
	Qualifier Qualifier
}

type Label struct {
	Label string
}

type Lexer struct {
	section
	Decls []LexerDecl
}

type LexerDecl interface {
	isLexerDecl()
}
type lexerDecl struct{}

func (d *lexerDecl) isLexerDecl() {}

type CustomTokenDecl struct {
	lexerDecl
	CustomTokens []*CustomToken
}

type CustomToken struct {
	Name    string
	Literal string
}
