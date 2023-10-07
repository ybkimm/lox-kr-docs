package ast

import "github.com/dcaiafa/lox/internal/lexergen/mode"

type Mode struct {
	baseStatement

	Name  string
	Rules []Statement

	Mode *mode.Mode
}

func (m *Mode) RunPass(ctx *Context, pass Pass) {
	if pass == CreateNames {
		if !ctx.RegisterName(m.Name, m) {
			return
		}
		m.Mode = ctx.CreateMode(m.Name)
	}

	ctx.CurrentLexerMode.Push(m.Mode)
	defer ctx.CurrentLexerMode.Pop()

	RunPass(ctx, m.Rules, pass)
}
