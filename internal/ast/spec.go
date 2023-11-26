package ast

import (
	"fmt"
	"slices"

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
				ctx.Errs.GeneralError(fmt.Errorf("@start rule undefined"))
				return
			}
			ctx.Grammar.SetStart(ctx.StartParserRule.Rule)
		}

		modeNames := make([]string, 0, len(ctx.LexerModes))
		for name := range ctx.LexerModes {
			modeNames = append(modeNames, name)
		}
		slices.Sort(modeNames)

		ctx.LexerDFAs = make(map[string]*mode.Mode, len(ctx.LexerModes))
		for i, name := range modeNames {
			mode := ctx.LexerModes[name].Build(ctx.Errs, ctx.FSet)
			mode.Index = i
			ctx.LexerDFAs[name] = mode
		}
	}
}
