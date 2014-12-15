package eulermath

import (
	"errors"
	"fmt"
	"math/big"
)

func NumberOfCombinations(n, r int64) (result uint64, err error) {
	result = 0

	if r > n {
		err = errors.New(fmt.Sprintf("r cannot be greater than n for values (n=%d, r=%d)", n, r))
		return
	}

	a := Factorial(big.NewInt(n))
	b := Factorial(big.NewInt(r))
	c := Factorial(big.NewInt(n - r))

	divisor := new(big.Int)
	divisor.Mul(b, c)

	divResult := new(big.Int)
	divResult.Div(a, divisor)

	result = divResult.Uint64()

	return
}
