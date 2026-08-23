package histogram

import "testing"

func TestTaskPerks015GapDirection(t *testing.T) {
	if got := gapWeight(&Bin{1, 1}, &Bin{1, 4}); got != 3 {
		t.Fatalf("gap=%v", got)
	}
}
