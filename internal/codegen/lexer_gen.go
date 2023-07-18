package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

const lexerGenGo = "lexer.gen.go"

const lexerTemplate = `
const (
{{- range i, t :=  terminals }}
	{{ t.Name }} = {{ i }}
{{- end }}
)
`

type LexerGenState struct {
	ImplDir string
	Grammar *grammar.AugmentedGrammar
}

func NewLexerGenState(
	implDir string,
	g *grammar.AugmentedGrammar,
) *LexerGenState {
	return &LexerGenState{
		ImplDir: implDir,
		Grammar: g,
	}
}

func (s *LexerGenState) Generate() error {
	packageName, err := computePackageName(s.ImplDir)
	if err != nil {
		return err
	}

	vars := jet.VarMap{}
	vars.Set("terminals", s.Grammar.Terminals)
	body := renderTemplate(lexerTemplate, vars)
	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", packageName)
	out.WriteString(body)
	err = os.WriteFile(filepath.Join(s.ImplDir, lexerGenGo), out.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}
