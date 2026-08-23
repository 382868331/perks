package quantile

import (
	"sort"
	"testing"
)

func TestTaskPerks001SampleOrdering(t *testing.T) {
	v := Samples{{Value: 2}, {Value: 1}}
	sort.Sort(v)
	if v[0].Value != 1 {
		t.Fatalf("first=%v", v[0].Value)
	}
}
