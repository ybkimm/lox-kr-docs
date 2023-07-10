package codegen

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type row struct {
	Cols  []int32
	Index int
}

func (r *row) Add(v int32) {
	r.Cols = append(r.Cols, v)
}

func (r *row) Key() string {
	key := make([]byte, 0, binary.MaxVarintLen32*len(r.Cols))
	for n := range r.Cols {
		key = binary.AppendVarint(key, int64(n))
	}
	return string(key)
}

type rows struct {
	rowMap map[string]*row
	rows   []*row
	size   int
}

func newRows() *rows {
	return &rows{
		rowMap: make(map[string]*row),
	}
}

func (r *rows) Add(row *row) *row {
	existing, ok := r.rowMap[row.Key()]
	if ok {
		return existing
	}
	r.rows = append(r.rows, row)
	r.rowMap[row.Key()] = row
	r.size += len(row.Cols)
	return row
}

func (r *rows) ToArray() []int32 {
	arr := make([]int32, 0, r.size+len(r.rows))
	for _, row := range r.rows {
		row.Index = len(arr)
		arr = append(arr, int32(len(row.Cols)))
		arr = append(arr, row.Cols...)
	}
	return arr
}

func writeArray(w *strings.Builder, xs []int32) {
	for i, x := range xs {
		if i != 0 && i%20 == 0 {
			w.WriteByte('\n')
		}
		fmt.Fprintf(w, "%d, ", x)
	}
	w.WriteByte('\n')
}
