package rang3

import (
	"container/heap"
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"

	"github.com/dcaiafa/lox/internal/util/stack"
)

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

func (r Range) Touches(o Range) bool {
	a, b := r, o
	if a.B > b.B {
		a, b = b, a
	}
	return b.B <= a.E || b.B-1 == a.E
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

func Compare(a, b Range) int {
	switch {
	case a.B < b.B:
		return -1
	case a.B > b.B:
		return 1
	case a.E < b.E:
		return -1
	case a.E > b.E:
		return 1
	default:
		return 0
	}
}

func Normalize(ranges []Range, onChange func(o, a, b, c Range)) {
	rh := rangeHeap(ranges)
	heap.Init(&rh)

	for len(rh) > 1 {
		x := heap.Pop(&rh).(Range)
		y := rh[0]

		if x == y {
			continue
		}

		if !x.Intersects(y) {
			continue
		}

		switch {
		case x.B == y.B && x.E < y.E:
			// x ---
			// y ------
			// ===================
			// x ---
			// a    ---
			a := Range{x.E + 1, y.E}
			onChange(y, x, a, a)
			heap.Pop(&rh) // pop y
			heap.Push(&rh, x)
			heap.Push(&rh, a)

		case x.B < y.B && x.E == y.E:
			// x ------
			// y    ---
			// ===================
			// a ---
			// y    ---
			a := Range{x.B, y.B - 1}
			onChange(x, a, y, y)
			heap.Push(&rh, a)

		case x.B < y.B && x.E < y.E:
			// x ------
			// y    ------
			// ==================
			// a ---
			// b    ---
			// c       ---
			a := Range{x.B, y.B - 1}
			b := Range{y.B, x.E}
			c := Range{x.E + 1, y.E}
			onChange(x, a, b, b)
			onChange(y, b, c, c)
			heap.Pop(&rh) // pop y
			heap.Push(&rh, a)
			heap.Push(&rh, b)
			heap.Push(&rh, c)

		case x.B < y.B && x.E > y.E:
			// x ---------
			// y    ---
			// ==================
			// a ---
			// y    ---
			// b       ---
			a := Range{x.B, y.B - 1}
			b := Range{y.E + 1, x.E}
			onChange(x, a, y, b)
			heap.Push(&rh, a)
			heap.Push(&rh, b)

		default:
			panic("not reached")
		}
	}
}

func Flatten(ranges []Range) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		return Compare(a, b)
	})

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
			if tip.Touches(r) {
				ranges2.Pop()
				ranges2.Push(Range{B: min(tip.B, r.B), E: max(tip.E, r.E)})
				continue
			}
		}
		ranges2.Push(r)
	}
	return ranges2.Elements()
}

func Negate(ranges []Range) []Range {
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
