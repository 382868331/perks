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

func TestQuantileRankIntervalClamp(t *testing.T) {
	if got := QuantileRankIntervalClamp(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileRankIntervalClampRegression(t *testing.T) {
	TestQuantileRankIntervalClamp(t)
	TestQuantileRankIntervalClamp(t)
}

func TestQuantileSampleCountSaturation(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := QuantileSampleCountSaturation(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileSampleCountSaturationRegression(t *testing.T) {
	TestQuantileSampleCountSaturation(t)
	TestQuantileSampleCountSaturation(t)
}

func TestQuantileEscapedObjectiveList(t *testing.T) {
	got := QuantileEscapedObjectiveList("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileEscapedObjectiveListRegression(t *testing.T) {
	TestQuantileEscapedObjectiveList(t)
	TestQuantileEscapedObjectiveList(t)
}

func TestQuantileDuplicateSampleStability(t *testing.T) {
	got := QuantileDuplicateSampleStability([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileDuplicateSampleStabilityRegression(t *testing.T) {
	TestQuantileDuplicateSampleStability(t)
	TestQuantileDuplicateSampleStability(t)
}

func TestQuantileZeroCompressionBatch(t *testing.T) {
	if got := QuantileZeroCompressionBatch([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileZeroCompressionBatchRegression(t *testing.T) {
	TestQuantileZeroCompressionBatch(t)
	TestQuantileZeroCompressionBatch(t)
}

func TestQuantileUnicodeMetricCutoff(t *testing.T) {
	got := QuantileUnicodeMetricCutoff("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileUnicodeMetricCutoffRegression(t *testing.T) {
	TestQuantileUnicodeMetricCutoff(t)
	TestQuantileUnicodeMetricCutoff(t)
}

func TestQuantileObjectiveBooleanOption(t *testing.T) {
	got, err := QuantileObjectiveBooleanOption(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestQuantileObjectiveBooleanOptionRegression(t *testing.T) {
	TestQuantileObjectiveBooleanOption(t)
	TestQuantileObjectiveBooleanOption(t)
}
