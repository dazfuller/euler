package eulermath

import (
	"math/big"
	"testing"
)

func TestGetFactorialFor0(t *testing.T) {
	result := Factorial(big.NewInt(0))
	expected := big.NewInt(1)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected 1 but got %d", result)
	}
}

func TestGetFactorialFor1(t *testing.T) {
	result := Factorial(big.NewInt(1))
	expected := big.NewInt(1)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected 1 but got %d", result)
	}
}

func TestGetFactorialFor6(t *testing.T) {
	result := Factorial(big.NewInt(6))
	expected := big.NewInt(720)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected 720 but got %d", result)
	}
}

func TestGetFactorialFor20(t *testing.T) {
	result := Factorial(big.NewInt(20))
	expected := big.NewInt(2432902008176640000)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestGetFactorialFor100(t *testing.T) {
	result := Factorial(big.NewInt(100))
	expected := new(big.Int)
	expected.SetString("93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000", 10)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func BenchmarkFactorialFor6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(big.NewInt(6))
	}
}

func BenchmarkFactorialFor20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(big.NewInt(20))
	}
}

func BenchmarkFactorialFor100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(big.NewInt(100))
	}
}
