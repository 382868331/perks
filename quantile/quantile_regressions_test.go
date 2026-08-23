package quantile

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"unicode/utf8"
)

var (
	_ = context.Background
	_ = errors.Is
	_ = reflect.DeepEqual
	_ = utf8.ValidString
)

func TestQuantileNormalizeBounds(t *testing.T) {
	if got := QuantileNormalizeBounds(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileNormalizeBoundsRegression(t *testing.T) {
	TestQuantileNormalizeBounds(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileNormalizeBounds(t)
}

func TestQuantileSaturatingAdd(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := QuantileSaturatingAdd(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileSaturatingAddRegression(t *testing.T) {
	TestQuantileSaturatingAdd(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileSaturatingAdd(t)
}

func TestQuantileSplitEscapedTokens(t *testing.T) {
	got := QuantileSplitEscapedTokens("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileSplitEscapedTokensRegression(t *testing.T) {
	TestQuantileSplitEscapedTokens(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileSplitEscapedTokens(t)
}

func TestQuantileStableUnique(t *testing.T) {
	got := QuantileStableUnique([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileStableUniqueRegression(t *testing.T) {
	TestQuantileStableUnique(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileStableUnique(t)
}

func TestQuantilePartitionValues(t *testing.T) {
	if got := QuantilePartitionValues([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestQuantilePartitionValuesRegression(t *testing.T) {
	TestQuantilePartitionValues(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantilePartitionValues(t)
}

func TestQuantileTruncateLabel(t *testing.T) {
	got := QuantileTruncateLabel("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileTruncateLabelRegression(t *testing.T) {
	TestQuantileTruncateLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileTruncateLabel(t)
}

func TestQuantileParseBooleanOption(t *testing.T) {
	got, err := QuantileParseBooleanOption(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}
