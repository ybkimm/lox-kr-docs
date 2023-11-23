package stablemap

type Map[K comparable, V any] struct {
	nodes map[K]*node[K, V]
	list  *node[K, V]
}

func initMap[K comparable, V any](m *Map[K, V]) {
	if m.nodes != nil {
		return
	}
	m.nodes = make(map[K]*node[K, V])
	m.list = &node[K, V]{}
	initList(m.list)
}

func (m *Map[K, V]) Len() int {
	if m.nodes == nil {
		return 0
	}
	return len(m.nodes)
}

func (m *Map[K, V]) Has(k K) bool {
	if m.nodes == nil {
		return false
	}
	n := m.nodes[k]
	return n != nil
}

func (m *Map[K, V]) Clear() {
	if m.nodes == nil {
		return
	}
	clear(m.nodes)
	initList(m.list)
}

func (m *Map[K, V]) Put(k K, v V) {
	initMap(m)
	n := m.nodes[k]
	if n == nil {
		n = &node[K, V]{key: k}
		insertNodeAfter(n, m.list.prev)
		m.nodes[k] = n
	}
	n.value = v
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	var v V
	if m.nodes == nil {
		return v, false
	}
	n := m.nodes[k]
	if n == nil {
		return v, false
	}
	return n.value, true
}

func (m *Map[K, V]) GetOrZero(k K) V {
	v, _ := m.Get(k)
	return v
}

func (m *Map[K, V]) Remove(k K) {
	if m.nodes == nil {
		return
	}
	n := m.nodes[k]
	if n == nil {
		return
	}
	removeNode(n)
	delete(m.nodes, k)
}

func (m *Map[K, V]) ForEach(f func(key K, value V)) {
	if m.list == nil {
		return
	}
	for n := m.list.next; n != m.list; n = n.next {
		f(n.key, n.value)
	}
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.nodes))
	m.ForEach(func(k K, _ V) {
		keys = append(keys, k)
	})
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, len(m.nodes))
	m.ForEach(func(_ K, v V) {
		values = append(values, v)
	})
	return values
}
