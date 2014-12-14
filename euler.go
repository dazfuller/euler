package main

import (
	"fmt"
	"github.com/dazfuller/euler/problems"
)

func solve(number int, problem problems.Solver) {
	answer, duration := problem.Solve()

	fmt.Printf("The answer to problem %d is: %s\n", number, answer)
	fmt.Printf("The solution took %v to run\n", duration)
	fmt.Println()
}

func main() {
	solve(50, new(problems.Problem50))
	solve(52, new(problems.Problem52))
	solve(53, new(problems.Problem53))
}