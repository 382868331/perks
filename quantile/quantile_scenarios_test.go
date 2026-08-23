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

func TestQuantileInvariantGrowthOverflow(t *testing.T) {
	if got := QuantileInvariantGrowthOverflow(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileInvariantGrowthOverflowRegression(t *testing.T) {
	TestQuantileInvariantGrowthOverflow(t)
	TestQuantileInvariantGrowthOverflow(t)
}

func TestQuantileQuantileOneIndex(t *testing.T) {
	if got := QuantileQuantileOneIndex([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileQuantileOneIndexRegression(t *testing.T) {
	TestQuantileQuantileOneIndex(t)
	TestQuantileQuantileOneIndex(t)
}

func TestQuantileSummaryCloneIsolation(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := QuantileSummaryCloneIsolation(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestQuantileSummaryCloneIsolationRegression(t *testing.T) {
	TestQuantileSummaryCloneIsolation(t)
	TestQuantileSummaryCloneIsolation(t)
}

func TestQuantileReverseMetricRunes(t *testing.T) {
	if got := QuantileReverseMetricRunes("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileReverseMetricRunesRegression(t *testing.T) {
	TestQuantileReverseMetricRunes(t)
	TestQuantileReverseMetricRunes(t)
}

func TestQuantileWindowInclusiveTail(t *testing.T) {
	got := QuantileWindowInclusiveTail([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileWindowInclusiveTailRegression(t *testing.T) {
	TestQuantileWindowInclusiveTail(t)
	TestQuantileWindowInclusiveTail(t)
}

func TestQuantileEmptySummaryRendering(t *testing.T) {
	if got := QuantileEmptySummaryRendering(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestQuantileEmptySummaryRenderingRegression(t *testing.T) {
	TestQuantileEmptySummaryRendering(t)
	TestQuantileEmptySummaryRendering(t)
}

func TestQuantileParallelInsertCount(t *testing.T) {
	if got := QuantileParallelInsertCount(64); got != 64 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileParallelInsertCountRegression(t *testing.T) {
	TestQuantileParallelInsertCount(t)
	TestQuantileParallelInsertCount(t)
}

func TestQuantileCanceledStreamDrain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := QuantileCanceledStreamDrain(ctx, 20); got != 0 {
		t.Fatalf("got %d", got)
	}
}

func TestQuantileCanceledStreamDrainRegression(t *testing.T) {
	TestQuantileCanceledStreamDrain(t)
	TestQuantileCanceledStreamDrain(t)
}

func TestQuantileObjectiveErrorChain(t *testing.T) {
	base := errors.New("root")
	if got := QuantileObjectiveErrorChain(base); !errors.Is(got, base) {
		t.Fatalf("chain lost: %v", got)
	}
}

func TestQuantileObjectiveErrorChainRegression(t *testing.T) {
	TestQuantileObjectiveErrorChain(t)
	TestQuantileObjectiveErrorChain(t)
}

func TestQuantileBufferReleaseOnQuery(t *testing.T) {
	active = 0
	QuantileBufferReleaseOnQuery(true)
	if active != 0 {
		t.Fatalf("active=%d", active)
	}
}

func TestQuantileBufferReleaseOnQueryRegression(t *testing.T) {
	TestQuantileBufferReleaseOnQuery(t)
	TestQuantileBufferReleaseOnQuery(t)
}

func TestQuantileTupleDeleteIterator(t *testing.T) {
	got := QuantileTupleDeleteIterator([]int{2, 4, 5, 6})
	if !reflect.DeepEqual(got, []int{5}) {
		t.Fatalf("got %v", got)
	}
}

func TestQuantileTupleDeleteIteratorRegression(t *testing.T) {
	TestQuantileTupleDeleteIterator(t)
	TestQuantileTupleDeleteIterator(t)
}

func TestQuantileEstimateBoolContract(t *testing.T) {
	v, ok := QuantileEstimateBoolContract(nil)
	if ok || v != 0 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestQuantileEstimateBoolContractRegression(t *testing.T) {
	TestQuantileEstimateBoolContract(t)
	TestQuantileEstimateBoolContract(t)
}
