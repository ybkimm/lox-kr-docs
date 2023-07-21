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
		return "{{ t.Name }}"
{{- end }}
	default:
		return "???"
	}
}

type Token struct {
	Pos  {{i("go/token")}}.Pos
	Type TokenType
	Str  string
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
	err = os.WriteFile(filepath.Join(s.ImplDir, lexerGenGo), out.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (s *LexerGenState) i(path string) string {
	return s.imports.Import(path)
}
