package quantile

import "testing"

func TestTaskPerks002LowBiasInvariant(t *testing.T) {
	s := NewLowBiased(.1)
	s.n = 10
	if got := s.ƒ(s.stream, 5); got != 1 {
		t.Fatalf("bound=%v", got)
	}
}

func TestTaskPerks002LowBiasInvariantBoundary(t *testing.T) {
	s := NewLowBiased(.2)
	s.n = 20
	if got := s.ƒ(s.stream, 10); got != 4 {
		t.Fatalf("bound=%v", got)
	}
}
