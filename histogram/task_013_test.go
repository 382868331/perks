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
