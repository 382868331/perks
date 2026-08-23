package topk

import "testing"

func TestTaskPerks018EvictionCleanup(t *testing.T) {
	s := New(1)
	s.Insert("a")
	s.Insert("b")
	s.Insert("c")
	if len(s.mon) > 2 {
		t.Fatalf("mon=%d", len(s.mon))
	}
}

func TestTaskPerks018EvictionCleanupBoundary(t *testing.T) {
	s := New(2)
	for _, v := range []string{"a", "b", "c", "d"} {
		s.Insert(v)
	}
	if len(s.mon) > 3 {
		t.Fatalf("mon=%d", len(s.mon))
	}
}
