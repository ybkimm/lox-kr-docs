package codegen2

import (
	"cmp"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
)

const lexerTemplate = `
{{ range _, mode := modes() }}

var _lexerMode{{mode.Index}} = []uint32 {
	{{ mode_table(mode) | array }}
}

{{ end }}

/*
type _LexerStateMachine struct {
	OnToken   func(t TokenType)
	OnDiscard func()

	state int
}

func (l *_LexerStateMachine) PushRune(r rune) {
}
*/

`

func (c *context) EmitLexer() bool {
	vars := make(jet.VarMap)

	vars.Set("array", func(arr []uint32) string {
		var str strings.Builder
		WriteArray(&str, arr)
		return str.String()
	})

	vars.Set("modes", func() []*mode.Mode {
		modes := make([]*mode.Mode, 0, len(c.LexerModes))
		for _, mode := range c.LexerModes {
			modes = append(modes, mode)
		}
		slices.SortFunc(modes, func(a, b *mode.Mode) int {
			return cmp.Compare(a.Index, b.Index)
		})
		return modes
	})

	vars.Set("mode_table", func(mode *mode.Mode) []uint32 {
		return []uint32{1, 2, 3}
	})

	lexerGen := renderTemplate(
		c.GoPackageName, c.GoPackagePath, lexerTemplate, vars)

	err := os.WriteFile(
		filepath.Join(c.Dir, lexerGenGo), []byte(lexerGen), 0666)
	if err != nil {
		c.Errs.GeneralError(err)
		return false
	}

	return true
}
