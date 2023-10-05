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
	m.list.prev = m.list
	m.list.next = m.list
}

func Len[K comparable, V any](m *Map[K, V]) int {
	if m.nodes == nil {
		return 0
	}
	return len(m.nodes)
}

func Has[K comparable, V any](m *Map[K, V], k K) bool {
	if m.nodes == nil {
		return false
	}
	n := m.nodes[k]
	return n != nil
}

func Put[K comparable, V any](m *Map[K, V], k K, v V) {
	initMap(m)
	n := m.nodes[k]
	if n == nil {
		n = &node[K, V]{key: k}
		insertNodeAfter(n, m.list.prev)
		m.nodes[k] = n
	}
	n.value = v
}

func Get[K comparable, V any](m *Map[K, V], k K) (V, bool) {
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

func Remove[K comparable, V any](m *Map[K, V], k K) {
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

func ForEach[K comparable, V any](m *Map[K, V], f func(key K, value V)) {
	if m.list == nil {
		return
	}
	for n := m.list.next; n != m.list; n = n.next {
		f(n.key, n.value)
	}
}

func Keys[K comparable, V any](m *Map[K, V]) []K {
	keys := make([]K, 0, len(m.nodes))
	ForEach(m, func(k K, v V) {
		keys = append(keys, k)
	})
	return keys
}
