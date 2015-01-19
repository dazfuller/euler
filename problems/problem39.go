/*
// If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
//
// {20,48,52}, {24,45,51}, {30,40,50}
//
// For which value of p ≤ 1000, is the number of solutions maximised?
*/

package problems

import (
	"fmt"
	"math"
	"time"
)

type Problem39 struct{}

func (p39 *Problem39) testForB(p, a int) (b int, valid bool) {
	fp := float64(p)
	fa := float64(a)

	fb := (fp * (fp - 2.0*fa)) / (2.0 * (fp - fa))
	b = int(fb)
	valid = math.Abs(fb-math.Floor(fb)) < 0.000001

	return
}

func (p39 *Problem39) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	result := []int{0, 0}

	for p := 2; p <= 1000; p += 2 {
		numberOfSolutions := 0
		for a := 2; a < p/3; a++ {
			if b, isValid := p39.testForB(p, a); isValid && a <= b {
				numberOfSolutions++
			}
		}
		if numberOfSolutions > result[1] {
			result[0] = p
			result[1] = numberOfSolutions
		}
	}

	t1 := time.Now()

	answer = fmt.Sprintf("%d with %d solutions", result[0], result[1])
	runTime = t1.Sub(t0)

	return
}
