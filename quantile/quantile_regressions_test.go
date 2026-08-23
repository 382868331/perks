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
