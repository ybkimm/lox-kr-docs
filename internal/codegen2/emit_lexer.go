package codegen2

import (
	"cmp"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/lexergen/dfa"
	"github.com/dcaiafa/lox/internal/lexergen/mode"
	"github.com/dcaiafa/lox/internal/lexergen/rang3"
)

const lexerTemplatePlaceholder = `
type _LexerStateMachine struct {
	token int
	state int
	mode  []uint32
}
func (l *_LexerStateMachine) PushRune(r rune) int {
	panic("not implemented")
}
func (l *_LexerStateMachine) Reset() {
	panic("not implemented")
}
func (l *_LexerStateMachine) Token() int {
	panic("not implemented")
}
`

func renderLexerTemplatePlaceholder(pkgName string) string {
	return renderTemplate(
		pkgName, "", lexerTemplatePlaceholder, make(jet.VarMap))
}

const lexerTemplate = `
{{ range _, mode := modes() }}

var _lexerMode{{mode.Index}} = []uint32 {
	{{ mode_table(mode) }}
}

{{ end }}

const (
	_lexerConsume  = 0
	_lexerAccept   = 1
	_lexerDiscard  = 2
	_lexerEOF      = 3
	_lexerError    = -1
)

type _LexerStateMachine struct {
	token int
	state int
	mode  []uint32
}

func (l *_LexerStateMachine) PushRune(r rune) int {
	if l.mode == nil {
		l.mode = _lexerMode0
	}

	i := int(l.mode[int(l.state)])
	count := int(l.mode[i])
	i++
	end := i + count

	gotoEnd := i + int(l.mode[i])* 3
	i++
	for ; i < gotoEnd; i += 3 {
		if r >= rune(l.mode[i]) &&
			r <= rune(l.mode[i+1]) {
			l.state = int(l.mode[i+2])
			return _lexerConsume
		}
	}

	for ; i < end; i += 2 {
		switch l.mode[i] {
		case 3: // Accept
			l.token = int(l.mode[i+1])
			l.state = 0
			return _lexerAccept
		case 4: // Discard
			l.state = 0
			return _lexerDiscard
		case 5: // Accum
			l.state = 0
			return _lexerConsume
		}
	}

	if l.state == 0 && r == 0 {
		return _lexerEOF
	}

	return _lexerError
}

func (l *_LexerStateMachine) Reset() {
	l.mode = nil
	l.state = 0
}

func (l *_LexerStateMachine) Token() int {
	return l.token
}
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

	vars.Set("mode_table", func(m *mode.Mode) string {
		table := newTable[uint32]()
		for _, state := range m.DFA.States {
			var row []uint32
			actions := state.Data.(*mode.Actions)

			row = append(row, uint32(state.Transitions.Len()))
			state.Transitions.ForEach(func(eventRaw any, toState *dfa.State) {
				event := eventRaw.(rang3.Range)
				row = append(row, uint32(event.B), uint32(event.E), toState.ID)
			})

			if actions != nil {
				for _, action := range actions.Actions {
					switch action.Type {
					case mode.ActionAccept:
						row = append(row, 3, uint32(action.Terminal))
					case mode.ActionDiscard:
						row = append(row, 4, 0)
					case mode.ActionAccum:
						row = append(row, 5, 0)
					default:
						panic("unreachable")
					}
				}
			}

			assert.True(len(row) > 0)
			table.AddRow(int(state.ID), row)
		}
		return table.String()
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
