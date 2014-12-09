/* 
// Combinatoric selections
// 
// There are exactly ten ways of selecting three from five, 12345:
// 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
// 
// In combinatorics, we use the notation, 5C3 = 10.  In general,
// 
// nCr = n! / r!(n−r)!
// where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.
// It is not until n = 23, that a value exceeds one-million: 23C10 = 1144066.
// 
// How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath"
	"log"
	"time"
)

type Problem53 struct {}

func (p *Problem53) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()
	count := 0

	for n := int64(23); n <= 100; n++ {
		for r := int64(1); r <= n; r++ {
			combinations, err := eulermath.NumberOfCombinations(n, r)
			if err != nil {
				log.Fatal(err)
			}
			if combinations > 1e6 {
				count++
			}
		}
	}

	t1 := time.Now()

	answer = fmt.Sprintf("The answer to problem 53 is: %d", count)
	runTime = t1.Sub(t0)

	return
}