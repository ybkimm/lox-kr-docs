package ast

import (
	gotoken "go/token"
)

type Bounds struct {
	Begin gotoken.Pos
	End   gotoken.Pos
}

type AST interface {
	SetBounds(b Bounds)
	Bounds() Bounds
}

type baseAST struct {
	bounds Bounds
}

func (a *baseAST) SetBounds(b Bounds) {
	a.bounds = b
}

func (a *baseAST) Bounds() Bounds {
	return a.bounds
}

type Qualifier int

const (
	NoQualifier Qualifier = iota
	ZeroOrMore            // *
	OneOrMore             // +
	ZeroOrOne             // ?
)

type Section interface {
	AST
	isSection()
}

type section struct{}

func (s *section) isSection() {}

type Spec struct {
	baseAST
	Sections []Section
}

type Parser struct {
	baseAST
	section
	Decls []ParserDecl
}

type ParserDecl interface {
	AST
	isParserDecl()
}

type parserDecl struct{}

func (d *parserDecl) isParserDecl() {}

type Rule struct {
	baseAST
	parserDecl
	Name  string
	Prods []*Prod
}

type Associativity int

const (
	Left  Associativity = 0
	Right Associativity = 1
)

type ProdQualifier struct {
	baseAST
	Precedence    int
	Associativity Associativity
}

type Prod struct {
	baseAST
	Terms     []*Term
	Qualifier *ProdQualifier
}

type Term struct {
	baseAST
	Name      string
	Qualifier Qualifier
}

type Lexer struct {
	baseAST
	section
	Decls []LexerDecl
}

type LexerDecl interface {
	AST
	isLexerDecl()
}

type lexerDecl struct{}

func (d *lexerDecl) isLexerDecl() {}

type CustomTokenDecl struct {
	baseAST
	lexerDecl
	CustomTokens []*CustomToken
}

type CustomToken struct {
	baseAST
	Name string
}
