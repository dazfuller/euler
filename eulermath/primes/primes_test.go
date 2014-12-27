package primes

import (
	"reflect"
	"testing"
)

func TestGetPrimesForNEqualToOne(t *testing.T) {
	primes := GetPrimeNumbersBelowN(1)
	if len(primes) != 0 {
		t.Error("Expected an empty slice")
	}
}

func TestGetPrimesForNEqualToTwo(t *testing.T) {
	primes := GetPrimeNumbersBelowN(3)
	if len(primes) != 1 {
		t.Error("Expected a single prime number")
	} else if primes[0] != 2 {
		t.Errorf("The prime number returned should be 2, got %d", primes[0])
	}
}

func TestGetPrimesForNEqualToTen(t *testing.T) {
	primes := GetPrimeNumbersBelowN(10)
	expected := []int64{
		2,
		3,
		5,
		7,
	}
	if len(primes) != len(expected) {
		t.Errorf("Expected %d primes but got %d", len(expected), len(primes))
		return
	}

	for i, v := range expected {
		if v != primes[i] {
			t.Errorf("Expected value %d at position %d but found %d", v, i, primes[i])
		}
	}
}

func TestGetPrimesForNEqualTo1Million(t *testing.T) {
	primes := GetPrimeNumbersBelowN(1e6)
	largest := primes[len(primes)-1]
	if largest != 999983 {
		t.Errorf("Expected 999983 but got %d", largest)
	}
}

func BenchmarkGetPrimesForNEqualTo1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimeNumbersBelowN(1000)
	}
}

func TestGetPrimeFactorsFor32(t *testing.T) {
	factors := PrimeFactors(32)
	expected := make(map[int64]int64)
	expected[2] = 5

	if reflect.DeepEqual(expected, factors) == false {
		t.Errorf("Expected %v but got %v", expected, factors)
	}
}

func TestGetPrimeFactorsFor60(t *testing.T) {
	factors := PrimeFactors(60)
	expected := make(map[int64]int64)
	expected[2] = 2
	expected[3] = 1
	expected[5] = 1

	if reflect.DeepEqual(expected, factors) == false {
		t.Errorf("Expected %v but got %v", expected, factors)
	}
}

func TestGetPrimeFactorsFor132(t *testing.T) {
	factors := PrimeFactors(132)
	expected := make(map[int64]int64)
	expected[2] = 2
	expected[3] = 1
	expected[11] = 1

	if reflect.DeepEqual(expected, factors) == false {
		t.Errorf("Expected %v but got %v", expected, factors)
	}
}

func TestGetPrimeFactorsFor9367(t *testing.T) {
	factors := PrimeFactors(9367)
	expected := make(map[int64]int64)
	expected[17] = 1
	expected[19] = 1
	expected[29] = 1

	if reflect.DeepEqual(expected, factors) == false {
		t.Errorf("Expected %v but got %v", expected, factors)
	}
}

func BenchmarkGetPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(9367)
	}
}
