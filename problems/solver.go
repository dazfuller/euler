package problems

import (
	"time"
)

type Solver interface {
	Solve() (string, time.Duration)
}
