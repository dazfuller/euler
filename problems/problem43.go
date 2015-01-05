/*
// The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.
//
// Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:
//
// d2 d3 d4=406 is divisible by 2
// d3 d4 d5=063 is divisible by 3
// d4 d5 d6=635 is divisible by 5
// d5 d6 d7=357 is divisible by 7
// d6 d7 d8=572 is divisible by 11
// d7 d8 d9=728 is divisible by 13
// d8 d9 d10=289 is divisible by 17
//
// Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath/eulerutil"
	"time"
)

type Problem43 struct{}

func (p *Problem43) toIntegerForCheck(d []int) int {
	return (d[0] * 100) + (d[1] * 10) + d[2]
}

func (p *Problem43) checkDivisibility(d []int) bool {
	if d[3]%2 != 0 {
		return false
	}

	if p.toIntegerForCheck(d[2:5])%3 != 0 {
		return false
	}

	if d[5]%5 != 0 {
		return false
	}

	if p.toIntegerForCheck(d[4:7])%7 != 0 {
		return false
	}

	if p.toIntegerForCheck(d[5:8])%11 != 0 {
		return false
	}

	if p.toIntegerForCheck(d[6:9])%13 != 0 {
		return false
	}

	if p.toIntegerForCheck(d[7:10])%17 != 0 {
		return false
	}

	return true
}

func (p *Problem43) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	digits := []int{1, 4, 0, 6, 3, 5, 7, 2, 8, 9}
	sum := int64(0)

	if p.checkDivisibility(digits) {
		sum += eulerutil.ToInteger(digits)
	}

	for eulerutil.NextPermutation(digits) {
		if p.checkDivisibility(digits) {
			sum += eulerutil.ToInteger(digits)
		}
	}

	t1 := time.Now()

	answer = fmt.Sprint(sum)
	runTime = t1.Sub(t0)

	return
}
