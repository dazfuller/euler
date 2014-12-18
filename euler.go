package main

import (
	"flag"
	"fmt"
	"github.com/dazfuller/euler/problems"
)

func solve(number int, problem problems.Solver) {
	answer, duration := problem.Solve()

	fmt.Printf("The answer to problem %d is: %s\n", number, answer)
	fmt.Printf("The solution took %v to run\n", duration)
	fmt.Println()
}

var problemId int

func init() {
	flag.IntVar(&problemId, "p", 0, "Which problem should a solution be run for")
}

func main() {
	p := make(map[int]problems.Solver)
	p[49] = new(problems.Problem49)
	p[50] = new(problems.Problem50)
	p[52] = new(problems.Problem52)
	p[53] = new(problems.Problem53)

	flag.Parse()

	if problemId == 0 {
		for k, v := range p {
			solve(k, v)
		}
	} else if solution, ok := p[problemId]; ok {
		solve(problemId, solution)
	} else {
		fmt.Println("That problem has not been solved yet")
	}
}
