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
