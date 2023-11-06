package ast

import (
	gotoken "go/token"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
	"github.com/dcaiafa/lox/internal/util/stack"
)

const DefaultModeName = "$default"

type Context struct {
	FSet *gotoken.FileSet
	Errs *errlogger.ErrLogger

	StartParserRule   *ParserRule
	HasParserRules    bool
	CurrentUnit       stack.Stack[*Unit]
	CurrentParserRule stack.Stack[*ParserRule]
	CurrentParserProd stack.Stack[*ParserProd]
	CurrentPrinter    stack.Stack[*Printer]
	CurrentLexerMode  stack.Stack[*mode.ModeBuilder]
	Grammar           *lr2.Grammar
	LexerModes        map[string]*mode.ModeBuilder

	names   map[string]AST
	aliases map[string]*TokenRule
}

func NewContext(fset *gotoken.FileSet, errs *errlogger.ErrLogger) *Context {
	c := &Context{
		FSet:       fset,
		Errs:       errs,
		Grammar:    lr2.NewGrammar(),
		LexerModes: make(map[string]*mode.ModeBuilder),
		names:      make(map[string]AST),
		aliases:    make(map[string]*TokenRule),
	}
	defaultMode := mode.New("")
	c.CurrentLexerMode.Push(defaultMode)
	return c
}

func (c *Context) RegisterName(name string, ast AST) bool {
	otherAST, alreadyExists := c.names[name]
	if alreadyExists {
		c.Errs.Errorf(c.Position(ast), "%v redefined", name)
		c.Errs.Infof(c.Position(otherAST), "other %v defined here", name)
		return false
	}
	c.names[name] = ast
	return true
}

func (c *Context) CreateMode(name string) *mode.ModeBuilder {
	_, ok := c.LexerModes[name]
	if ok {
		// This should never happen because the name conflict should have been
		// caught already when calling CreateName.
		panic("mode redefined")
	}
	m := mode.New(name)
	c.LexerModes[name] = m
	return m
}

func (c *Context) Lookup(name string) AST {
	return c.names[name]
}

var AmbiguousAlias = &TokenRule{}

func (c *Context) CreateAlias(name string, t *TokenRule) {
	existing := c.aliases[name]
	if existing != nil {
		// Can't use aliases if there is more than one token with the same literal.
		c.aliases[name] = AmbiguousAlias
		return
	}
	c.aliases[name] = t
}

func (c *Context) LookupAlias(name string) *TokenRule {
	return c.aliases[name]
}

func (c *Context) Position(ast AST) gotoken.Position {
	return c.FSet.Position(ast.Bounds().Begin)
}

func (c *Context) Mode() *mode.ModeBuilder {
	return c.CurrentLexerMode.Peek()
}

func (c *Context) Print(ast AST, out io.Writer) {
	c.CurrentPrinter.Push(NewPrinter(out))
	ast.RunPass(c, Print)
	c.CurrentPrinter.Pop()
}

func RunPass[T AST](ctx *Context, asts []T, pass Pass) {
	for _, ast := range asts {
		ast.RunPass(ctx, pass)
	}
}

func (c *Context) Analyze(ast AST, untilPass Pass) bool {
	for _, pass := range passes {
		ast.RunPass(c, pass)
		if c.Errs.HasError() {
			return false
		}
		if pass == untilPass {
			break
		}
	}
	return true
}
