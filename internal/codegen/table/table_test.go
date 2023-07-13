package table

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRows(t *testing.T) {
	r := New()

	r.AddRow(0, []int32{1, 2, 3})
	r.AddRow(1, []int32{3, 4})
	r.AddRow(2, []int32{1, 2, 3})
	r.AddRow(4, []int32{3, 4})
	r.AddRow(5, []int32{1, 2})

	fmt.Println(r.String())

	expected := []int32{
		6, 10, 6, -1, 10, 13,
		3, 1, 2, 3,
		2, 3, 4,
		2, 1, 2,
	}

	if !reflect.DeepEqual(r.Array(), expected) {
		t.Log("Expected:\n", expected)
		t.Log("Actual:\n", r.Array())
		t.Fatalf("Result does not match expectations")
	}
}
