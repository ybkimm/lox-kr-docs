package mode

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dcaiafa/lox/internal/lexergen/nfa"
)

type Range struct {
	B rune
	E rune
}

func (r Range) String() string {
	p := func(c rune) string {
		var buf strings.Builder
		switch c {
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '-':
			buf.WriteString(`\-`)
		default:
			if unicode.IsGraphic(c) {
				buf.WriteRune(c)
			} else {
				fmt.Fprintf(&buf, "\\u%04x", c)
			}
		}
		return buf.String()
	}
	if r.B == r.E {
		return p(r.B)
	} else {
		return p(r.B) + "-" + p(r.E)
	}
}

type NFAComposite struct {
	B *nfa.State
	E *nfa.State
}

type Mode struct {
	Name string
	NFA  *nfa.StateFactory
}

func New(name string) *Mode {
	return &Mode{
		Name: name,
		NFA:  nfa.NewStateFactory(),
	}
}
