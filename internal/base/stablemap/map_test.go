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
		testutil.RequireEqual(t, m.Keys(), expectedKeys)
		testutil.RequireEqual(t, m.Len(), len(expectedKeys))

		i := 0
		m.ForEach(func(k int, v string) {
			testutil.RequireEqual(t, k, ns[i].(int))
			testutil.RequireEqual(t, v, ns[i+1].(string))
			v2, ok := m.Get(k)
			testutil.RequireTrue(t, ok)
			testutil.RequireEqual(t, v2, v)
			ok = m.Has(k)
			testutil.RequireTrue(t, ok)
			i += 2
		})
	}

	m.Remove(0)
	_, ok := m.Get(1)
	testutil.RequireFalse(t, ok)
	ok = m.Has(1)
	testutil.RequireFalse(t, ok)
	check()

	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")
	check(1, "1", 2, "2", 3, "3")

	m.Put(2, "2'")
	check(1, "1", 2, "2'", 3, "3")

	m.Remove(1)
	m.Remove(1)
	check(2, "2'", 3, "3")
	_, ok = m.Get(1)
	testutil.RequireFalse(t, ok)

	m.Put(1, "1'")
	m.Put(1, "1'")
	check(2, "2'", 3, "3", 1, "1'")

	m.Remove(1)
	check(2, "2'", 3, "3")

	m.Put(4, "4")
	check(2, "2'", 3, "3", 4, "4")

	m.Remove(3)
	check(2, "2'", 4, "4")

	m.Remove(2)
	m.Remove(4)
	check()

	m.Put(1, "1")
	m.Put(2, "2")
	m.Put(3, "3")
	check(1, "1", 2, "2", 3, "3")
}
