package codegen2

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type table[E int32 | uint32] struct {
	maxIndex int
	rowMap   map[string]int
	index    map[int]int
	arr      []E
}

func newTable[E int32 | uint32]() *table[E] {
	return &table[E]{
		maxIndex: -1,
		rowMap:   make(map[string]int),
		index:    make(map[int]int),
	}
}

func (r *table[E]) AddRow(index int, row []E) {
	if index <= r.maxIndex {
		panic("index must be monotonically increasing")
	}
	r.maxIndex = index
	key := r.rowKey(row)
	existing, ok := r.rowMap[key]
	if ok {
		r.index[index] = existing
		return
	}
	r.index[index] = len(r.arr)
	r.rowMap[key] = len(r.arr)
	r.arr = append(r.arr, E(len(row)))
	r.arr = append(r.arr, row...)
}

func (r *table[E]) Array() []E {
	arr := make([]E, 0, r.maxIndex+1+len(r.arr))
	for i := 0; i <= r.maxIndex; i++ {
		x, ok := r.index[i]
		if ok {
			x += r.maxIndex + 1
		} else {
			x = -1
		}
		arr = append(arr, E(x))
	}
	arr = append(arr, r.arr...)
	return arr
}

func (r *table[E]) String() string {
	var str strings.Builder
	WriteArray(&str, r.Array())
	return str.String()
}

func (r *table[E]) rowKey(xs []E) string {
	key := make([]byte, 0, binary.MaxVarintLen32*len(xs))
	for _, x := range xs {
		key = binary.AppendVarint(key, int64(x))
	}
	return string(key)
}

func WriteArray[T int32 | uint32](w *strings.Builder, xs []T) {
	for i, x := range xs {
		if i != 0 && i%14 == 0 {
			w.WriteByte('\n')
		}
		fmt.Fprintf(w, "%d, ", x)
	}
}
