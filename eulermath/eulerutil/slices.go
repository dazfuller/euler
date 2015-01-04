package eulerutil

import (
	"math"
	"sort"
)

func ToInteger(u sort.IntSlice) int64 {
	result := float64(0)
	n := len(u)

	for i, v := range u {
		exp := float64(n - i - 1)
		result += math.Pow(float64(10), exp) * float64(v)
	}

	return int64(result)
}
