package eulermath

import (
	"testing"
)

func TestForInvalidParameterValues(t *testing.T) {
	_, err := NumberOfCombinations(5, 6)
	expectedError := "r cannot be greater than n for values (n=5, r=6)"
	if err == nil {
		t.Error("Expected call to return an error but none was produced")
	} else if err.Error() != expectedError {
		t.Errorf("The returned error was not expected: %s", err)
	}
}

func TestFor10Combinations(t *testing.T) {
	result, err := NumberOfCombinations(5, 3)
	expected := uint64(10)

	if err != nil {
		t.Errorf("Got error: %s", err)
	} else if result != expected {
		t.Errorf("Expected 10 but got %d", result)
	}
}

func BenchmarkFor10Combinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfCombinations(5, 3)
	}
}
