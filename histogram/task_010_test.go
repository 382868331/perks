package histogram

import "testing"

func TestTaskPerks010BinMeanDivisor(t *testing.T) {
	b := &Bin{Count: 1, Sum: 8}
	if got := b.Mean(); got != 8 {
		t.Fatalf("mean=%v", got)
	}
}

func TestTaskPerks010BinMeanDivisorBoundary(t *testing.T) {
	b := &Bin{Count: 4, Sum: 10}
	if got := b.Mean(); got != 2.5 {
		t.Fatalf("mean=%v", got)
	}
}
