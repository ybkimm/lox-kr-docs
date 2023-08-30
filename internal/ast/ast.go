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

type TermType int

const (
	Simple     TermType = iota
	ZeroOrMore          // *
	OneOrMore           // +
	ZeroOrOne           // ?
	List                // @list
)

type Parser struct {
	baseAST
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
	Type  TermType
	Name  string
	Child *Term
	Sep   *Term
}

type CustomTokenDecl struct {
	baseAST
	parserDecl
	CustomTokens []*CustomToken
}

type CustomToken struct {
	baseAST
	Name string
}
