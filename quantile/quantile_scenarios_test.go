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
