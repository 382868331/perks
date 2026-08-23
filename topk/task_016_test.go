package topk

import "testing"

func TestTaskPerks016ExistingCountDirection(t *testing.T) {
	s := New(2)
	s.Insert("a")
	s.Insert("a")
	if got := s.Query()[0].Count; got != 2 {
		t.Fatalf("count=%d", got)
	}
}

func TestTaskPerks016ExistingCountDirectionBoundary(t *testing.T) {
	s := New(2)
	s.Insert("x")
	s.Insert("x")
	s.Insert("x")
	if got := s.Query()[0].Count; got != 3 {
		t.Fatalf("count=%d", got)
	}
}
