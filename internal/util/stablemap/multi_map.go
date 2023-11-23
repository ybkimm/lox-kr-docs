package stablemap

import "github.com/dcaiafa/lox/internal/util/array"

type MultiMap[K comparable, V any] struct {
	Map[K, *array.Array[V]]
}

func (m *MultiMap[K, V]) Add(k K, v V) {
	arr, ok := m.Get(k)
	if !ok {
		arr = new(array.Array[V])
		m.Put(k, arr)
	}
	arr.Add(v)
}
