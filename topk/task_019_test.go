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

func TestTaskPerks019ReplacementCountBoundary(t *testing.T) {
	s := New(1)
	s.Merge(Samples{&Element{"a", 2}, &Element{"b", 3}, &Element{"c", 4}})
	if s.Query()[0].Count <= 4 {
		t.Fatalf("count=%d", s.Query()[0].Count)
	}
}
