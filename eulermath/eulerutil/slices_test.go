package eulerutil

import (
	"reflect"
	"testing"
)

func TestSliceToInteger(t *testing.T) {
	digits := []int{3, 1, 4, 1, 5, 9}
	expected := int64(314159)
	actual := ToInteger(digits)

	if expected != actual {
		t.Errorf("Expected value %d but got %d", expected, actual)
	}
}

func BenchmarkSliceToInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		ToInteger(digits)
	}
}

func TestSlicePermutationsPartial(t *testing.T) {
	d := []int{1, 2, 3, 4}
	for i := 0; i < 7; i++ {
		NextPermutation(d)
	}

	expected := []int{2, 1, 4, 3}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func TestSlicePermutationsFull(t *testing.T) {
	d := []int{1, 2, 3, 4}
	for NextPermutation(d) {
	}

	expected := []int{4, 3, 2, 1}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := []int{1, 2, 3, 4}
		for NextPermutation(d) {
		}
	}
}
