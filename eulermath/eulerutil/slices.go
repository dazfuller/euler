package eulerutil

import (
	"math"
)

func ToInteger(slice []int) int64 {
	k := int64(0)
	for i := 0; i < len(slice); i++ {
		k = int64(10)*k + int64(slice[i])
	}
	return k
}

func ToSlice(n int64) []int {
	count := int(math.Floor(math.Log10(float64(n)) + 1.0))
	ar := make([]int, count)
	count--
	for n != 0 {
		ar[count] = int(n % 10)
		n /= 10
		count--
	}
	return ar
}

type compare func(a, b int) bool

func Permutation(slice []int, cmp compare) bool {
	var k int
	for k = len(slice) - 2; k >= 0; k-- {
		if cmp(slice[k], slice[k+1]) {
			break
		}
	}

	if k == -1 {
		return false
	}

	var l int
	for l = len(slice) - 1; l >= 0; l-- {
		if cmp(slice[k], slice[l]) {
			break
		}
	}

	slice[k], slice[l] = slice[l], slice[k]

	for i, j := k+1, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return true
}

func NextPermutation(slice []int) bool {
	return Permutation(slice, func(a, b int) bool { return a < b })
}

func PrevPermutation(slice []int) bool {
	return Permutation(slice, func(a, b int) bool { return a > b })
}
