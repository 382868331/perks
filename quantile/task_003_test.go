package quantile

import "testing"

func TestTaskPerks003TargetSliceCapacity(t *testing.T) {
	if got := len(targetMapToSlice(map[float64]float64{.5: .1})); got != 1 {
		t.Fatalf("len=%d", got)
	}
}

func TestTaskPerks003TargetSliceCapacityBoundary(t *testing.T) {
	if got := len(targetMapToSlice(map[float64]float64{.1: .01, .9: .01})); got != 2 {
		t.Fatalf("len=%d", got)
	}
}
