/*
// Take the number 192 and multiply it by each of 1, 2, and 3:
//
// 192 × 1 = 192
// 192 × 2 = 384
// 192 × 3 = 576
// By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated
// product of 192 and (1,2,3)
//
// The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital,
// 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
//
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an
// integer with (1,2, ... , n) where n > 1?
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/eulerutil"
	"sort"
	"time"
)

type Problem38 struct{}

func (p *Problem38) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	for i := int64(9876); i >= int64(9182); i-- {
		i2 := i * 2
		d := append(eulerutil.ToSlice(i), eulerutil.ToSlice(i2)...)
		if p.isPandigital(d) {
			answer = fmt.Sprintf("%d%d", i, i2)
			break
		}
	}

	t1 := time.Now()

	runTime = t1.Sub(t0)

	return
}

func (p *Problem38) isPandigital(digits []int) bool {
	panDigits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(digits)

	if len(digits) != len(panDigits) {
		return false
	}

	for i, v := range digits {
		if v != panDigits[i] {
			return false
		}
	}

	return true
}
