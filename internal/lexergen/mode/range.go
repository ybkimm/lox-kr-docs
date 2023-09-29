package mode

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/dcaiafa/lox/internal/util/stack"
)

func NormalizeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		switch {
		case ranges[i].B < ranges[j].B:
			return true
		case ranges[i].B > ranges[j].B:
			return false
		case ranges[i].E < ranges[j].B:
			return true
		default:
			return false
		}
	})
	min := func(a, b rune) rune {
		if a < b {
			return a
		} else {
			return b
		}
	}
	max := func(a, b rune) rune {
		if a > b {
			return a
		} else {
			return b
		}
	}

	ranges2 := stack.Stack[Range]{}
	ranges2.SetCapacity(len(ranges))
	for _, r := range ranges {
		if !ranges2.Empty() {
			tip := ranges2.Peek()
			if tip.Intersects(r) {
				ranges2.Pop()
				ranges2.Push(Range{B: min(tip.B, r.B), E: max(tip.E, r.E)})
				continue
			}
		}
		ranges2.Push(r)
	}
	return ranges2.Elements()
}

func NegateRanges(ranges []Range) []Range {
	v := stack.Stack[Range]{}
	v.SetCapacity(len(ranges))
	var b rune = 0
	for _, r := range ranges {
		if r.B > 0 {
			v.Push(Range{B: b, E: r.B - 1})
		}
		b = r.E + 1
	}
	if b < 0xFFFF {
		v.Push(Range{B: b, E: 0xFFFF})
	}
	return v.Elements()
}

type Range struct {
	B rune
	E rune
}

func (r Range) Contains(o Range) bool {
	return r.B <= o.B && r.E >= o.E
}

func (r Range) Intersects(o Range) bool {
	a, b := r, o
	if a.B > b.B {
		a, b = b, a
	}
	return b.B <= a.E
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
