package quantile

import "testing"

func TestTaskPerks007CountBufferedSamples(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(1)
	if got := s.Count(); got != 1 {
		t.Fatalf("count=%d", got)
	}
}
