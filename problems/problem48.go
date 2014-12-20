/*
// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
//
// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package problems

import (
	"fmt"
	"time"
)

type Problem48 struct{}

func (p *Problem48) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	sum := uint64(1)

	// calculate each term but only keep the last 10 digits, this allows us to use
	// integers without overflowing
	for i := uint64(2); i < 1000; i++ {
		t := uint64(1)
		for j := uint64(1); j <= i; j++ {
			t *= i
			t %= 10000000000
		}
		sum += t
		sum %= 10000000000
	}

	t1 := time.Now()

	answer = fmt.Sprint(sum)
	runTime = t1.Sub(t0)

	return
}
