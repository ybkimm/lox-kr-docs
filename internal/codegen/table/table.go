package table

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type Table struct {
	maxIndex int
	rowMap   map[string]int
	index    map[int]int
	arr      []int32
}

func New() *Table {
	return &Table{
		maxIndex: -1,
		rowMap:   make(map[string]int),
		index:    make(map[int]int),
	}
}

func (r *Table) AddRow(index int, row []int32) {
	if index <= r.maxIndex {
		panic("index must be monotonically increasing")
	}
	r.maxIndex = index
	key := rowKey(row)
	existing, ok := r.rowMap[key]
	if ok {
		r.index[index] = existing
		return
	}
	r.index[index] = len(r.arr)
	r.rowMap[key] = len(r.arr)
	r.arr = append(r.arr, int32(len(row)))
	r.arr = append(r.arr, row...)
}

func (r *Table) Array() []int32 {
	arr := make([]int32, 0, r.maxIndex+1+len(r.arr))
	for i := 0; i <= r.maxIndex; i++ {
		x, ok := r.index[i]
		if ok {
			x += r.maxIndex + 1
		} else {
			x = -1
		}
		arr = append(arr, int32(x))
	}
	arr = append(arr, r.arr...)
	return arr
}

func (r *Table) String() string {
	var str strings.Builder
	WriteArray(&str, r.Array())
	return str.String()
}

func rowKey(xs []int32) string {
	key := make([]byte, 0, binary.MaxVarintLen32*len(xs))
	for _, x := range xs {
		key = binary.AppendVarint(key, int64(x))
	}
	return string(key)
}

func WriteArray(w *strings.Builder, xs []int32) {
	for i, x := range xs {
		if i != 0 && i%14 == 0 {
			w.WriteByte('\n')
		}
		fmt.Fprintf(w, "%d, ", x)
	}
	w.WriteByte('\n')
}
