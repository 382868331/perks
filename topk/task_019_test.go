package topk

import "testing"

func TestTaskPerks019ReplacementCount(t *testing.T) {
	s := New(1)
	s.Insert("a")
	s.Insert("b")
	s.Insert("c")
	q := s.Query()
	if q[0].Count < 2 {
		t.Fatalf("count=%d", q[0].Count)
	}
}
