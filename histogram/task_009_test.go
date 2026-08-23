package histogram

import "testing"

func TestTaskPerks009BinSumUpdate(t *testing.T) {
	b := &Bin{Count: 2, Sum: 5}
	b.Update(&Bin{Count: 1, Sum: 3})
	if b.Sum != 8 {
		t.Fatalf("sum=%v", b.Sum)
	}
}
