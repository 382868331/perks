package histogram

import "testing"

func TestTaskPerks013EqualMeanAggregation(t *testing.T) {
	h := New(10)
	h.Insert(4)
	h.Insert(4)
	if got := len(h.Bins()); got != 1 {
		t.Fatalf("bins=%d", got)
	}
}

func TestTaskPerks013EqualMeanAggregationBoundary(t *testing.T) {
	h := New(10)
	h.Insert(-2)
	h.Insert(-2)
	h.Insert(-2)
	if h.Bins()[0].Count != 3 {
		t.Fatalf("count=%d", h.Bins()[0].Count)
	}
}
