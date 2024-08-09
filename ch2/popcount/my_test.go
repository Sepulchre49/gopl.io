package popcount_test

import (
	"fmt"
	"testing"

	"gopl.io/ch2/popcount"
)

func popcount_test(t *testing.T) {
    result := popcount.PopCount(10)
    fmt.Sprintf("%d", result)
}
