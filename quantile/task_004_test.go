package quantile

import "testing"

func TestTaskPerks004InsertedWidth(t *testing.T) {
	s := NewLowBiased(.01)
	s.Insert(3)
	if got := s.b[0].Width; got != 1 {
		t.Fatalf("width=%v", got)
	}
}
