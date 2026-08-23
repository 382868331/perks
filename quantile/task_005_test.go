package quantile

import "testing"

func TestTaskPerks005BufferFlushFence(t *testing.T) {
	s := NewLowBiased(.01)
	for i := 0; i < 500; i++ {
		s.Insert(float64(i))
	}
	if !s.flushed() {
		t.Fatal("buffer did not flush")
	}
}
