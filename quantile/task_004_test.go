package quantile

import "testing"

func TestTaskPerks004InsertedWidth(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(3)
	if got := s.b[0].Width; got != 1 {
		t.Fatalf("width=%v", got)
	}
}

func TestTaskPerks004InsertedWidthBoundary(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(7)
	if got := s.Count(); got != 1 {
		t.Fatalf("count=%d", got)
	}
}
