package main

import (
	"fmt"
	"github.com/dazfuller/euler/problems"
)

func solve(problem problems.Solver) {
	answer, duration := problem.Solve()

	fmt.Println(answer)
	fmt.Printf("The solution took %v to run\n", duration)
	fmt.Println()
}

func main() {
	solve(new(problems.Problem52))
	solve(new(problems.Problem53))
}