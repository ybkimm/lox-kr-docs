package ast

import (
	gotoken "go/token"

	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

type Spec struct {
	baseAST

	Units []*Unit

	DefaultLexerMode *mode.ModeBuilder
}

func (s *Spec) RunPass(ctx *Context, pass Pass) {
	switch pass {
	case CreateNames:
		s.DefaultLexerMode = ctx.CreateMode(DefaultModeName)
	}

	ctx.CurrentLexerMode.Push(s.DefaultLexerMode)
	defer ctx.CurrentLexerMode.Pop()

	RunPass(ctx, s.Units, pass)

	if ctx.Errs.HasError() {
		return
	}

	switch pass {
	case GenerateGrammar:
		if ctx.HasParserRules {
			if ctx.StartParserRule == nil {
				ctx.Errs.Errorf(gotoken.Position{}, "@start rule undefined")
				return
			}
			ctx.Grammar.SetStart(ctx.StartParserRule.Rule)
		}
	}
}
