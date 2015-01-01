package eulermath

import (
	"math"
)

func GetTriangular(n int) int {
	return (n * (n + 1)) / 2
}

func IsTriangular(value int) bool {
	x := float64(value)
	n := (math.Sqrt(8*x+1) - 1.0) / 2.0

	return n == math.Floor(n)
}

func GetPentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func IsPentagonal(value int) bool {
	x := float64(value)
	n := (math.Sqrt(24*x+1) + 1.0) / 6.0

	return n == math.Floor(n)
}
