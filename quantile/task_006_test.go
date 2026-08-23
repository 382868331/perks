package quantile

import "testing"

func TestTaskPerks006QueryIndexAdjustment(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	if got := s.Query(.5); got != 2 {
		t.Fatalf("median=%v", got)
	}
}

func TestTaskPerks006QueryIndexAdjustmentBoundary(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(1)
	s.Insert(2)
	if got := s.Query(.5); got != 1 {
		t.Fatalf("median=%v", got)
	}
}
