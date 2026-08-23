package topk

import "testing"

func TestTaskPerks017MonitorCapacity(t *testing.T) {
	s := New(2)
	s.Insert("a")
	s.Insert("b")
	if len(s.mon) != 2 {
		t.Fatalf("mon=%d", len(s.mon))
	}
}
