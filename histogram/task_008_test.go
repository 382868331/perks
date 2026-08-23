package histogram

import "testing"

func TestTaskPerks008BinCountUpdate(t *testing.T) {
	b := &Bin{Count: 3, Sum: 6}
	b.Update(&Bin{Count: 2, Sum: 4})
	if b.Count != 5 {
		t.Fatalf("count=%d", b.Count)
	}
}
