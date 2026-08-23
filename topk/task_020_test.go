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
