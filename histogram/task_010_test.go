package histogram

import "testing"

func TestTaskPerks010BinMeanDivisor(t *testing.T) {
	b := &Bin{Count: 1, Sum: 8}
	if got := b.Mean(); got != 8 {
		t.Fatalf("mean=%v", got)
	}
}
