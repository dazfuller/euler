package main

import (
	"flag"
	"fmt"
	"github.com/dazfuller/euler/problems"
	"sort"
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
	p[39] = new(problems.Problem39)
	p[40] = new(problems.Problem40)
	p[41] = new(problems.Problem41)
	p[42] = new(problems.Problem42)
	p[43] = new(problems.Problem43)
	p[44] = new(problems.Problem44)
	p[45] = new(problems.Problem45)
	p[46] = new(problems.Problem46)
	p[47] = new(problems.Problem47)
	p[48] = new(problems.Problem48)
	p[49] = new(problems.Problem49)
	p[50] = new(problems.Problem50)
	p[52] = new(problems.Problem52)
	p[53] = new(problems.Problem53)

	flag.Parse()

	if problemId == 0 {
		// Get the keys
		keys := make([]int, len(p))
		i := 0
		for k, _ := range p {
			keys[i] = k
			i++
		}

		// Put the keys in order
		sort.Sort(sort.IntSlice(keys))

		// Solve the problems in order
		for _, k := range keys {
			solve(k, p[k])
		}

	} else if solution, ok := p[problemId]; ok {
		solve(problemId, solution)
	} else {
		fmt.Println("That problem has not been solved yet")
	}
}
