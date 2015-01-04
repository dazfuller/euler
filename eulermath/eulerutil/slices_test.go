package eulerutil

import (
	"sort"
	"testing"
)

func TestSliceToInteger(t *testing.T) {
	digits := sort.IntSlice{3, 1, 4, 1, 5, 9}
	expected := int64(314159)
	actual := ToInteger(digits)

	if expected != actual {
		t.Errorf("Expected value %d but got %d", expected, actual)
	}
}

func BenchmarkSliceToInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digits := sort.IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
		ToInteger(digits)
	}
}
