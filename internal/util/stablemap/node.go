package stablemap

type node[K comparable, V any] struct {
	next, prev *node[K, V]
	key        K
	value      V
}

func initList[K comparable, V any](l *node[K, V]) {
	l.prev = l
	l.next = l
}

// insertNodeAfter inserts list node n after node o.
func insertNodeAfter[K comparable, V any](n, o *node[K, V]) {
	n.prev = o
	n.next = o.next
	o.next.prev = n
	o.next = n
}

// removeNode removes a node from its list.
func removeNode[K comparable, V any](n *node[K, V]) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}
