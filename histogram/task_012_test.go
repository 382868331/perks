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

func TestTaskPerks012HeapPopSequence(t *testing.T) {
	bs := Bins{}
	for _, mean := range []float64{8, 2, 5, 1, 3} {
		heap.Push(&bs, &Bin{Count: 1, Sum: mean})
	}
	for _, want := range []float64{1, 2, 3, 5, 8} {
		got := heap.Pop(&bs).(*Bin).Mean()
		if got != want {
			t.Fatalf("pop=%v want=%v", got, want)
		}
	}
	if len(bs) != 0 {
		t.Fatalf("remaining=%d", len(bs))
	}
}
