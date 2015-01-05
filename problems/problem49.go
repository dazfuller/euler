/*
// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual
// in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations
// of one another.
//
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but
// there is one other 4-digit increasing sequence.
//
// What 12-digit number do you form by concatenating the three terms in this sequence?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/eulerutil"
	"github.com/dazfuller/euler/eulermath/primes"
	"math"
	"sort"
	"time"
)

var problem49Result *candidate

type Problem49 struct{}

func (problem *Problem49) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	var p []int64 = primes.GetPrimeNumbersBelowN(1e5)

	// Find the index of the first prime over 1000
	for i, v := range p {
		if v > 1000 {
			p = p[i:]
			break
		}
	}

	for _, v := range p {
		checkPrimeValue(v, p)
	}

	t1 := time.Now()

	answer = problem49Result.String()
	runTime = t1.Sub(t0)

	return
}

func int64ToArray(n int64) []int {
	a := make([]int, 4) // default to 4 as we're only dealing with values between 1000 and 9999

	for i := 3; i >= 0; i-- {
		a[i] = int(n % 10)
		n /= 10
	}

	sort.Ints(a) // Return the digits sorted

	return a
}

func intArrayToInt64(a []int) int64 {
	value := float64(0)
	n := len(a)

	for i, v := range a {
		exp := float64(n - i - 1)
		value += math.Pow(float64(10), exp) * float64(v)
	}

	return int64(value)
}

type candidate struct {
	First  int64
	Second int64
	Third  int64
}

func (c *candidate) String() string {
	return fmt.Sprintf("%d%d%d", c.First, c.Second, c.Third)
}

func hasCandidates(p []int64) (found bool, c *candidate) {
	found = false

	for x := 0; x < len(p)-2; x++ {
		for y := x + 1; y < len(p)-1; y++ {
			for z := y + 1; z < len(p); z++ {
				if p[y]-p[x] == p[z]-p[y] {
					c = new(candidate)
					c.First = p[x]
					c.Second = p[y]
					c.Third = p[z]

					found = true

					return
				}
			}
		}
	}

	c = nil
	return
}

func checkPrimeValue(prime int64, primeValues []int64) {
	validPrimes := make([]int64, 0)

	// Turn the prime value into an array of its digits
	a := int64ToArray(prime)

	for eulerutil.NextPermutation(a) {
		t := intArrayToInt64(a) // Convert the permutation back into an integer

		// Check to see if the permutation is a valid integer
		primeIndex := sort.Search(len(primeValues), func(i int) bool { return primeValues[i] >= t })
		if primeIndex < len(primeValues) && primeValues[primeIndex] == t {
			validPrimes = append(validPrimes, t)
		}
	}

	if len(validPrimes) >= 3 {
		if found, candidate := hasCandidates(validPrimes); found {
			if problem49Result == nil || problem49Result.First < candidate.First {
				problem49Result = candidate
			}
		}
	}
}
