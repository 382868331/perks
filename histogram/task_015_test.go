package histogram

import "testing"

func TestTaskPerks015GapDirection(t *testing.T) {
	if got := gapWeight(&Bin{1, 1}, &Bin{1, 4}); got != 3 {
		t.Fatalf("gap=%v", got)
	}
}

func TestTaskPerks015GapDirectionBoundary(t *testing.T) {
	if got := gapWeight(&Bin{2, -4}, &Bin{2, 6}); got != 5 {
		t.Fatalf("gap=%v", got)
	}
}
