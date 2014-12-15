package primes

import (
	"math"
)

func GetPrimeNumbersBelowN(end int64) []int64 {
	if end < int64(2) {
		return make([]int64, 0)
	}

	ar := make([]bool, end)
	ar[0] = true
	ar[1] = true
	ar[2] = false

	limit := int64(math.Sqrt(float64(end))) + 1

	n := float64(end)
	approxNumberPrimes := int64((n / (math.Log(n))) * (1 + (1.2762 / math.Log(n))))

	primes := make([]int64, approxNumberPrimes)
	primes[0] = 2
	count := 1

	for i := int64(3); i < end; i += 2 {
		if ar[i] == false {
			primes[count] = int64(i)
			count++

			if i < limit {
				for j := i * i; j < end; j += i {
					ar[j] = true
				}
			}
		}
	}

	return primes[:count]
}
