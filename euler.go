package main

import (
    "fmt"
    "github.com/dazfuller/euler/problems"
    "net/http"
    "strconv"
)

var problemMap map[int64]problems.Solver

// URLRoot gets a message when the root URL is requested
func URLRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi there")
}

// ProblemHandler executes the requested problem and returns a result
func ProblemHandler(w http.ResponseWriter, r *http.Request) {
    problemIDValue := r.URL.Path[len("/problem/"):]
    
    var problemID int64
    var err error
    
    if len(problemIDValue) == 0 {
        problemID = 0
    } else {
        problemID, err = strconv.ParseInt(problemIDValue, 10, 32)
        if err != nil {
            error := fmt.Sprintf("The value '%s' is not a valid integer value", problemIDValue)
            http.Error(w, error, http.StatusBadRequest)
            return
        }
    }
    
    err = solve(problemID, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func solve(problemID int64, w http.ResponseWriter) error {
    solution, ok := problemMap[problemID]
    
    if !ok {
        return fmt.Errorf("Problem ID %d is not a problem which has yet been solved", problemID)
    }
    
    answer, duration := solution.Solve()
    fmt.Fprintf(w, "The answer to problem %d is: %s\n", problemID, answer)
    fmt.Fprintf(w, "The solution took %v to run\n", duration)
    
    return nil
}

func initializeProblemMap() {
    problemMap = make(map[int64]problems.Solver)
    
    problemMap[38] = new(problems.Problem38)
	problemMap[39] = new(problems.Problem39)
	problemMap[40] = new(problems.Problem40)
	problemMap[41] = new(problems.Problem41)
	problemMap[42] = new(problems.Problem42)
	problemMap[43] = new(problems.Problem43)
	problemMap[44] = new(problems.Problem44)
	problemMap[45] = new(problems.Problem45)
	problemMap[46] = new(problems.Problem46)
	problemMap[47] = new(problems.Problem47)
	problemMap[48] = new(problems.Problem48)
	problemMap[49] = new(problems.Problem49)
	problemMap[50] = new(problems.Problem50)
	problemMap[52] = new(problems.Problem52)
	problemMap[53] = new(problems.Problem53)
}

func main() {
    initializeProblemMap()
    
    http.HandleFunc("/", URLRoot)
    http.HandleFunc("/problem/", ProblemHandler)
    http.ListenAndServe(":8080", nil)
}

/*
package main_old

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

var problemID int

func init() {
	flag.IntVar(&problemID, "p", 0, "Which problem should a solution be run for")
}

func main() {
	p := make(map[int]problems.Solver)
	p[38] = new(problems.Problem38)
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

	if problemID == 0 {
		// Get the keys
		keys := make([]int, len(p))
		i := 0
		for k := range p {
			keys[i] = k
			i++
		}

		// Put the keys in order
		sort.Sort(sort.IntSlice(keys))

		// Solve the problems in order
		for _, k := range keys {
			solve(k, p[k])
		}

	} else if solution, ok := p[problemID]; ok {
		solve(problemID, solution)
	} else {
		fmt.Println("That problem has not been solved yet")
	}
}
*/