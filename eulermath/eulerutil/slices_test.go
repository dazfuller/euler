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

func TestSliceNextPermutationPartial(t *testing.T) {
	d := []int{1, 2, 3, 4}
	for i := 0; i < 7; i++ {
		NextPermutation(d)
	}

	expected := []int{2, 1, 4, 3}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func TestSliceNextPermutationFull(t *testing.T) {
	d := []int{1, 2, 3, 4}
	for NextPermutation(d) {
	}

	expected := []int{4, 3, 2, 1}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func BenchmarkNextPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := []int{1, 2, 3, 4}
		for NextPermutation(d) {
		}
	}
}

func TestSlicePrevPermutationPartial(t *testing.T) {
	d := []int{4, 3, 2, 1}
	for i := 0; i < 7; i++ {
		PrevPermutation(d)
	}

	expected := []int{3, 4, 1, 2}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func TestSlicePrevPermutationFull(t *testing.T) {
	d := []int{4, 3, 2, 1}
	for PrevPermutation(d) {
	}

	expected := []int{1, 2, 3, 4}
	if reflect.DeepEqual(expected, d) == false {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func BenchmarkPrevPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := []int{4, 3, 2, 1}
		for PrevPermutation(d) {
		}
	}
}
