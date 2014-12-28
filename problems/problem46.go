/*
// It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime
// and twice a square.
//
// 9 = 7 + 2×12
// 15 = 7 + 2×22
// 21 = 3 + 2×32
// 25 = 7 + 2×32
// 27 = 19 + 2×22
// 33 = 31 + 2×12
//
// It turns out that the conjecture was false.
//
// What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/primes"
	"math"
	"time"
)

type Problem46 struct {
	primes []int64
}

func (p *Problem46) isConjectureValid(n int64) bool {
	// Check each prime value which is less than the value being evaluated
	for _, p := range p.primes {
		if p >= n {
			break
		}

		// Test the conjecture that the value can be written as the sum of a prime and twice a square
		// p + (2*(x*x))
		d := float64(n-p) / 2.0
		s := math.Sqrt(d)

		// If s is a whole number then the conjecture is valid for the current value
		if s == math.Floor(s) {
			return true
		}
	}

	// If this point is reached then the conjecture has failed (which is the scenario we are seeking)
	return false
}

func (p *Problem46) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	// Start from 9 as this is the smallest, odd composite number
	i := int64(9)

	// Set the first few prime numbers
	p.primes = []int64{2, 3, 5, 7}

	// Keep on checking values
	for true {
		// If the value is prime then add it to the collection to be used for testing the
		// conjecture
		if primes.IsPrime(i) {
			p.primes = append(p.primes, i)
		} else if p.isConjectureValid(i) == false {
			break
		}
		i += 2
	}

	t1 := time.Now()

	answer = fmt.Sprint(i)
	runTime = t1.Sub(t0)

	return
}
