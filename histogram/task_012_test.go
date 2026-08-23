package histogram

import (
	"container/heap"
	"testing"
)

func TestTaskPerks012HeapPopPosition(t *testing.T) {
	for _, means := range [][]float64{{1, 4}, {1, 3, 7}, {1, 2, 5, 9}} {
		bs := Bins{}
		for _, mean := range means {
			heap.Push(&bs, &Bin{Count: 1, Sum: mean})
		}
		got := heap.Pop(&bs)
		if got == nil || got.(*Bin).Mean() != 1 {
			t.Fatalf("means=%v pop=%v", means, got)
		}
		if len(bs) != len(means)-1 {
			t.Fatalf("means=%v remaining=%d", means, len(bs))
		}
	}
}
