package histogram

import "testing"

func TestTaskPerks014ReservoirCapacity(t *testing.T) {
	h := New(2)
	h.Insert(1)
	h.Insert(10)
	if got := len(h.Bins()); got != 2 {
		t.Fatalf("bins=%d", got)
	}
}
