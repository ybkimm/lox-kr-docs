package stablemap

import (
	"testing"

	"github.com/dcaiafa/lox/internal/testutil"
)

func TestMap(t *testing.T) {
	var m Map[int, string]

	check := func(ns ...any) {
		t.Helper()
		expectedKeys := make([]int, 0, len(ns)/2)
		for i := 0; i < len(ns); i += 2 {
			expectedKeys = append(expectedKeys, ns[i].(int))
		}
		testutil.RequireEqual(t, Keys(&m), expectedKeys)
		testutil.RequireEqual(t, Len(&m), len(expectedKeys))

		i := 0
		ForEach(&m, func(k int, v string) {
			testutil.RequireEqual(t, k, ns[i].(int))
			testutil.RequireEqual(t, v, ns[i+1].(string))
			v2, ok := Get(&m, k)
			testutil.RequireTrue(t, ok)
			testutil.RequireEqual(t, v2, v)
			ok = Has(&m, k)
			testutil.RequireTrue(t, ok)
			i += 2
		})
	}

	Remove(&m, 0)
	_, ok := Get(&m, 1)
	testutil.RequireFalse(t, ok)
	ok = Has(&m, 1)
	testutil.RequireFalse(t, ok)
	check()

	Put(&m, 1, "1")
	Put(&m, 2, "2")
	Put(&m, 3, "3")
	check(1, "1", 2, "2", 3, "3")

	Put(&m, 2, "2'")
	check(1, "1", 2, "2'", 3, "3")

	Remove(&m, 1)
	Remove(&m, 1)
	check(2, "2'", 3, "3")
	_, ok = Get(&m, 1)
	testutil.RequireFalse(t, ok)

	Put(&m, 1, "1'")
	Put(&m, 1, "1'")
	check(2, "2'", 3, "3", 1, "1'")

	Remove(&m, 1)
	check(2, "2'", 3, "3")

	Put(&m, 4, "4")
	check(2, "2'", 3, "3", 4, "4")

	Remove(&m, 3)
	check(2, "2'", 4, "4")

	Remove(&m, 2)
	Remove(&m, 4)
	check()

	Put(&m, 1, "1")
	Put(&m, 2, "2")
	Put(&m, 3, "3")
	check(1, "1", 2, "2", 3, "3")
}
