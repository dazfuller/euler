/*
// The first two consecutive numbers to have two distinct prime factors are:
//
// 14 = 2 × 7
// 15 = 3 × 5
//
// The first three consecutive numbers to have three distinct prime factors are:
//
// 644 = 2² × 7 × 23
// 645 = 3 × 5 × 43
// 646 = 2 × 17 × 19.
//
// Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/primes"
	"time"
)

type Problem47 struct{}

func (p *Problem47) Solve() (answer string, runTime time.Duration) {
	var start int64 = 2 * 3 * 5 * 7 // Start from the smallest possible value
	var first int64 = 0             // The first integer of the possible consecutive values
	k := 4                          // This determines the number of consecutive values
	ci := 0                         // A count of the consecutive integers

	t0 := time.Now()

	// Start the searcg from the smallest possible value and keep searching until the consecutive
	// count equals k
	i := start
	for ci != k {
		// Get the prime factors of the value, if the number of distinct prime factors is equal
		// to k then increment the count and record the current value if this is the first with
		// the correct count found. Otherwise reset the consecutive count
		pf := primes.PrimeFactors(i)
		if len(pf) == k {
			ci++
			if ci == 1 {
				first = i
			}
		} else {
			ci = 0
		}
		i++
	}

	t1 := time.Now()

	answer = fmt.Sprint(first)
	runTime = t1.Sub(t0)

	return
}
