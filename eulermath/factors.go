package eulermath

import (
	"math/big"
)

var c = big.NewInt(1)

func Factorial(n *big.Int) (result *big.Int) {
	if n.Cmp(c) == 0 || n.Cmp(c) == -1 {
		return c
	}

	result = new(big.Int)
	result.Set(n)
	result.Mul(result, Factorial(n.Sub(n, c)))

	return result
}
