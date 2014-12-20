/*
// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
//
// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

package problems

import (
	"math/big"
	"time"
)

type Problem48 struct{}

func (p *Problem48) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	sum := big.NewInt(1)
	mod := big.NewInt(10000000000)

	for i := int64(2); i < 1000; i++ {
		n := big.NewInt(i)
		sum.Add(sum, n.Exp(n, n, nil))
	}

	sum.Mod(sum, mod)

	t1 := time.Now()

	answer = sum.String()
	runTime = t1.Sub(t0)

	return
}
