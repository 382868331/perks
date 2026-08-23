package topk

import "testing"

func TestTaskPerks020QueryOrdering(t *testing.T) {
	s := New(2)
	s.Insert("a")
	s.Insert("a")
	s.Insert("b")
	if got := s.Query()[0].Value; got != "a" {
		t.Fatalf("first=%s", got)
	}
}

func TestTaskPerks020QueryOrderingBoundary(t *testing.T) {
	s := New(3)
	for _, v := range []string{"a", "a", "a", "b", "b", "c"} {
		s.Insert(v)
	}
	q := s.Query()
	if q[0].Count < q[1].Count {
		t.Fatalf("counts=%d,%d", q[0].Count, q[1].Count)
	}
}
