package primes

import (
	"testing"
)

func TestForNEqualToOne(t *testing.T) {
	primes := GetPrimeNumbersBelowN(1)
	if len(primes) != 0 {
		t.Error("Expected an empty slice")
	}
}

func TestForNEqualToTwo(t *testing.T) {
	primes := GetPrimeNumbersBelowN(3)
	if len(primes) != 1 {
		t.Error("Expected a single prime number")
	} else if primes[0] != 2 {
		t.Errorf("The prime number returned should be 2, got %d", primes[0])
	}
}

func TestForNEqualToTen(t *testing.T) {
	primes := GetPrimeNumbersBelowN(10)
	expected := []int64 {
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

func BenchmarkForNEqualTo1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimeNumbersBelowN(1000)
	}
}