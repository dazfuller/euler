package primes

import (
	"math"
)

type boolSlice []bool

type predicate func(x bool) bool

func (ar boolSlice) count(fn predicate) (c int64) {
	c = int64(0)
	for _, v := range ar {
		if fn(v) {
			c++
		}
	}
	return
}

func GetPrimeNumbersBelowN(end int64) []int64 {
	if end < int64(2) {
		return make([]int64, 0)
	}

	ar := make(boolSlice, end)
	ar[0] = true
	ar[1] = true

	limit := int64(math.Sqrt(float64(end))) + 1

	for i := int64(2); i < limit; i++ {
		if ar[i] == false {
			for j := i * 2; j < end; j += i {
				ar[j] = true
			}
		}
	}

	count := ar.count(func (x bool) bool {
		if x == false {
			return true
		}
		return false
	})

	primes := make([]int64, count)
	index := int64(0)
	for i, v := range ar {
		if v == false {
			primes[index] = int64(i)
			index++
		}
	}

	return primes
}