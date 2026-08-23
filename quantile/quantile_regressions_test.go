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

func TestQuantileParseBooleanOptionRegression(t *testing.T) {
	TestQuantileParseBooleanOption(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileParseBooleanOption(t)
}

func TestQuantileBoundedBackoff(t *testing.T) {
	if got := QuantileBoundedBackoff(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileBoundedBackoffRegression(t *testing.T) {
	TestQuantileBoundedBackoff(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileBoundedBackoff(t)
}

func TestQuantileSelectUpperQuantile(t *testing.T) {
	if got := QuantileSelectUpperQuantile([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileSelectUpperQuantileRegression(t *testing.T) {
	TestQuantileSelectUpperQuantile(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileSelectUpperQuantile(t)
}

func TestQuantileCloneNestedState(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := QuantileCloneNestedState(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestQuantileCloneNestedStateRegression(t *testing.T) {
	TestQuantileCloneNestedState(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileCloneNestedState(t)
}

func TestQuantileReverseUnicodeLabel(t *testing.T) {
	if got := QuantileReverseUnicodeLabel("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileReverseUnicodeLabelRegression(t *testing.T) {
	TestQuantileReverseUnicodeLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileReverseUnicodeLabel(t)
}

func TestQuantileSlidingWindows(t *testing.T) {
	got := QuantileSlidingWindows([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileSlidingWindowsRegression(t *testing.T) {
	TestQuantileSlidingWindows(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileSlidingWindows(t)
}

func TestQuantileJoinOptionalParts(t *testing.T) {
	if got := QuantileJoinOptionalParts(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileJoinOptionalPartsRegression(t *testing.T) {
	TestQuantileJoinOptionalParts(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileJoinOptionalParts(t)
}

func TestQuantileCountConcurrentUpdates(t *testing.T) {
	if got := QuantileCountConcurrentUpdates(64); got != 64 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileCountConcurrentUpdatesRegression(t *testing.T) {
	TestQuantileCountConcurrentUpdates(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileCountConcurrentUpdates(t)
}

func TestQuantileProcessUntilCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := QuantileProcessUntilCanceled(ctx, 20); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileProcessUntilCanceledRegression(t *testing.T) {
	TestQuantileProcessUntilCanceled(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileProcessUntilCanceled(t)
}

func TestQuantileWrapCause(t *testing.T) {
	base := errors.New("root")
	if got := QuantileWrapCause(base); !errors.Is(got, base) {
		t.Fatalf("chain lost: %v", got)
	}
}

func TestQuantileWrapCauseRegression(t *testing.T) {
	TestQuantileWrapCause(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileWrapCause(t)
}

func TestQuantileResetResourceState(t *testing.T) {
	active = 0
	QuantileResetResourceState(true)
	if active != 0 {
		t.Fatalf("active=%d", active)
	}
}

func TestQuantileResetResourceStateRegression(t *testing.T) {
	TestQuantileResetResourceState(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestQuantileResetResourceState(t)
}
