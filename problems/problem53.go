package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath"
	"log"
	"time"
)

type Problem53 struct {}

func (p *Problem53) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()
	count := 0

	for n := int64(23); n <= 100; n++ {
		for r := int64(1); r <= n; r++ {
			combinations, err := eulermath.NumberOfCombinations(n, r)
			if err != nil {
				log.Fatal(err)
			}
			if combinations > 1e6 {
				count++
			}
		}
	}

	t1 := time.Now()

	answer = fmt.Sprintf("The answer to problem 53 is: %d", count)
	runTime = t1.Sub(t0)

	return
}