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

func PrimeFactors(n int64) map[int64]int64 {
	result := make(map[int64]int64)
	if n == 1 {
		return result
	}
	for n%2 == 0 {
		result[2]++
		n /= 2
	}

	limit := int64(math.Sqrt(float64(n))) + 1
	var i int64 = 3
	for i <= limit {
		if n%i == 0 {
			result[i]++
			n /= i
			limit = int64(math.Sqrt(float64(n))) + 1
		} else {
			i += 2
		}
	}

	if n != 1 {
		result[n]++
	}

	return result
}
