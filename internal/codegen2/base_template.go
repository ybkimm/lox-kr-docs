package codegen2

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

const baseTemplate = `
const (
{{- range i, t := terminals }}
	{{ t.Name }} int = {{ i }}
{{- end }}
)

func _TokenToString(t int) string {
	switch t {
{{- range i, t := terminals }}
	case {{ t.Name }}: 
		return "{{ t.Alias != "" ? t.Alias : t.Name }}"
{{- end }}
	default:
		return "???"
	}
}

type _Stack[T any] []T

func (s *_Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *_Stack[T]) Pop(n int) {
	*s = (*s)[:len(*s)-n]
}

func (s _Stack[T]) Peek(n int) T {
	return s[len(s)-n-1]
}
`

type baseTemplateInputs struct {
	Package     string
	PackagePath string
	Terminals   []*lr2.Terminal
}

func renderBaseTemplate(in *baseTemplateInputs) string {
	vars := make(jet.VarMap)
	vars.Set("terminals", in.Terminals)
	return renderTemplate(in.Package, in.PackagePath, baseTemplate, vars)
}
