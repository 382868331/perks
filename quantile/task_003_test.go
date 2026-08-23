package quantile

import "testing"

func TestTaskPerks003TargetSliceCapacity(t *testing.T) {
	if got := len(targetMapToSlice(map[float64]float64{.5: .1})); got != 1 {
		t.Fatalf("len=%d", got)
	}
}
