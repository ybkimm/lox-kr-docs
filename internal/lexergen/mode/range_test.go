package mode

import "testing"

func TestRangeContains(t *testing.T) {
	test := func(name string, ab, ae, bb, be rune, expect bool) {
		t.Run(name, func(t *testing.T) {
			a := Range{B: ab, E: ae}
			b := Range{B: bb, E: be}
			result := a.Contains(b)
			if result != expect {
				t.Fatalf(
					"Ranges: %v %v expected: %v actual: %v",
					a, b, expect, result)
			}
		})
	}
	//    1     2     3     4     5     6    7    8
	// a ---    ---   --   ----  --    --  --      --
	// b  ---  ---   ----   --    --  --     --  --
	//    F      F     F    T     F    F     F    F
	test("1", 'a', 'd', 'b', 'e', false)
	test("2", 'b', 'e', 'a', 'f', false)
	test("3", 'c', 'd', 'a', 'f', false)
	test("4", 'a', 'f', 'c', 'd', true)
	test("5", 'a', 'b', 'b', 'c', false)
	test("6", 'b', 'c', 'a', 'b', false)
	test("7", 'a', 'b', 'c', 'd', false)
	test("8", 'c', 'd', 'a', 'b', false)
	test("9", 'c', 'd', 'c', 'd', true)
}

func TestRangeIntersects(t *testing.T) {
	test := func(name string, ab, ae, bb, be rune, expect bool) {
		t.Run(name, func(t *testing.T) {
			a := Range{B: ab, E: ae}
			b := Range{B: bb, E: be}
			result := a.Intersects(b)
			if result != expect {
				t.Fatalf(
					"Ranges: %v %v expected: %v actual: %v",
					a, b, expect, result)
			}
		})
	}

	//    1     2     3     4     5     6    7    8
	// a ---    ---   --   ----  --    --  --      --
	// b  ---  ---   ----   --    --  --     --  --
	//    T      T     T    T     T    T     F    F
	test("1", 'a', 'd', 'b', 'e', true)
	test("2", 'b', 'e', 'a', 'd', true)
	test("3", 'c', 'd', 'a', 'f', true)
	test("4", 'a', 'f', 'c', 'd', true)
	test("5", 'a', 'b', 'b', 'c', true)
	test("6", 'b', 'c', 'a', 'b', true)
	test("7", 'a', 'b', 'c', 'd', false)
	test("8", 'c', 'd', 'a', 'b', false)
}
