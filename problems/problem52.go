/*
// Permuted multiples
// 
// It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.
//
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/

package problems

import (
	"fmt"
	"sort"
	"time"
)

type Problem52 struct {}

func (p *Problem52) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()
	result := 0
	for x := 123456; x < 456789; x++ {
		if checkValue(x) {
			result = x
			break
		}
	}
	t1 := time.Now()

	answer = fmt.Sprintf("The answer to problem 52 is: %d", result)
	runTime = t1.Sub(t0)

	return
}

func areEqual(a, b sort.IntSlice) bool {
	if a.Len() != b.Len() {
		return false
	}

	a.Sort()
	b.Sort()

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func intToSlice(n int) sort.IntSlice {
	result := []int{}
	for n > 0 {
		result = append(result, n%10)
		n = n / 10
	}
	return result
}

func checkValue(n int) bool {
	a := intToSlice(n * 2)

	b := intToSlice(3 * n)

	if !areEqual(a, b) {
		return false
	}

	b = intToSlice(4 * n)

	if !areEqual(a, b) {
		return false
	}

	b = intToSlice(5 * n)

	if !areEqual(a, b) {
		return false
	}

	b = intToSlice(6 * n)

	if !areEqual(a, b) {
		return false
	}

	return true
}
