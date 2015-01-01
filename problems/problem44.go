/*
// Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten pentagonal numbers are:
//
// 1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...
//
// It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.
//
// Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and
// D = |Pk − Pj| is minimised; what is the value of D?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath"
	"time"
)

func findSolution() (int, int) {
	pentagonals := []int{}
	i := 1
	k := 0
	for true {
		k = eulermath.GetPentagonal(i)
		for _, j := range pentagonals {
			if eulermath.IsPentagonal(k-j) && eulermath.IsPentagonal(k+j) {
				return k, j
			}
		}
		pentagonals = append(pentagonals, k)
		i++
	}
	return 0, 0
}

type Problem44 struct{}

func (p *Problem44) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	k, j := findSolution()

	t1 := time.Now()

	answer = fmt.Sprint(k - j)
	runTime = t1.Sub(t0)

	return
}
