/*
// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n \
// exactly once. For example, 2143 is a 4-digit pandigital and is also prime.
//
// What is the largest n-digit pandigital prime that exists?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/eulerutil"
	"github.com/dazfuller/euler/eulermath/primes"
	"time"
)

type Problem41 struct{}

func (p *Problem41) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	d := []int{7, 6, 5, 4, 3, 2, 1}

	for eulerutil.PrevPermutation(d) {
		t := eulerutil.ToInteger(d)
		if primes.IsPrime(t) {
			answer = fmt.Sprint(t)
			break
		}
	}

	t1 := time.Now()

	runTime = t1.Sub(t0)

	return
}
