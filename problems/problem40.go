/*
// An irrational decimal fraction is created by concatenating the positive integers:
//
// 0.123456789101112131415161718192021...
//
// It can be seen that the 12th digit of the fractional part is 1.
//
// If dn represents the nth digit of the fractional part, find the value of the following expression.
//
// d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package problems

import (
	"fmt"
	"math"
	"time"
)

type Problem40 struct{}

func (p *Problem40) intLength(x int) int {
	return int(math.Floor(math.Log10(float64(x)) + 1.0))
}

func (p *Problem40) getNthDigit(x, n int) int {
	l := p.intLength(x)
	a := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		a[i] = x % 10
		x /= 10
	}

	return a[n]
}

func (p *Problem40) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	n := 1
	i := 1
	idx := 1
	sum := 0
	for n <= 1e6 {
		if idx >= n {
			sum += p.getNthDigit(i, idx-n)
			n *= 10
		}
		idx += p.intLength(i)
		i++
	}
	answer = fmt.Sprint(sum)

	t1 := time.Now()

	runTime = t1.Sub(t0)

	return
}
