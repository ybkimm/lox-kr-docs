package codegen

import (
	"bytes"
	"fmt"
	goformat "go/format"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

const lexerGenGo = "base.gen.go"

const lexerTemplate = `
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

type _Bounds struct {
	Begin Token
	End   Token
}

type lexer interface {
	ReadToken() (Token, TokenType)
}
`

type LexerGenState struct {
	ImplDir string
	Grammar *grammar.AugmentedGrammar
	imports *importBuilder
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
	s.imports = newImportBuilder()

	packageName, err := computePackageName(s.ImplDir)
	if err != nil {
		return err
	}

	vars := jet.VarMap{}
	vars.Set("p", prefix)
	vars.Set("i", s.i)
	vars.Set("terminals", s.Grammar.Terminals)
	body := renderTemplate(lexerTemplate, vars)
	out := bytes.NewBuffer(make([]byte, 0, len(body)+2048))
	fmt.Fprintf(out, "package %v\n\n", packageName)
	s.imports.WriteTo(out)
	out.WriteString(body)

	outFinal, err := goformat.Source(out.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format %v: %w", lexerGenGo, err)
	}

	err = os.WriteFile(filepath.Join(s.ImplDir, lexerGenGo), outFinal, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (s *LexerGenState) i(path string) string {
	return s.imports.Import(path)
}
