package histogram

import (
	"sort"
	"testing"
)

func TestTaskPerks011BinOrdering(t *testing.T) {
	bs := Bins{&Bin{1, 4}, &Bin{1, 2}}
	sort.Sort(bs)
	if bs[0].Mean() != 2 {
		t.Fatalf("first=%v", bs[0].Mean())
	}
}
