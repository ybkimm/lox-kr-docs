package rang3

type rangeHeap []Range

func (h rangeHeap) Len() int { return len(h) }

func (h rangeHeap) Less(i, j int) bool {
	return Compare(h[i], h[j]) < 0
}

func (h rangeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *rangeHeap) Push(x any) {
	item := x.(Range)
	*h = append(*h, item)
}

func (h *rangeHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
