package codegen2

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

const baseTemplate = `
type TokenType int

const (
{{- range i, t := terminals }}
	{{ t.Name }} TokenType = {{ i }}
{{- end }}
)

func (t TokenType) String() string {
	switch t {
{{- range i, t := terminals }}
	case {{ t.Name }}: 
		return "{{ t.Alias != "" ? t.Alias : t.Name }}"
{{- end }}
	default:
		return "???"
	}
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
