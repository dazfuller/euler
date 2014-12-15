/*
// The prime 41, can be written as the sum of six consecutive primes:
// 
// 41 = 2 + 3 + 5 + 7 + 11 + 13
// This is the longest sum of consecutive primes that adds to a prime below one-hundred.
// 
// The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
// 
// Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package problems

import (
    "fmt"
    "github.com/dazfuller/euler/eulermath/primes"
    "sort"
    "time"
)

type primeList []int64

func (p primeList) contains(n int64) bool {
    for i := len(p)-1; i >= 0; i-- {
        if p[i] == n {
            return true
        }
    }
    return false
}

type result struct {
    prime  int64
    length int
}

type Problem50 struct {}

func (problem *Problem50) Solve() (answer string, runTime time.Duration) {
    t0 := time.Now()

    max := int64(1e6)
    var p primeList = primes.GetPrimeNumbersBelowN(max)
    primeSum := make([]int64, len(p)+1)
    primeSum[0] = 0
    r := new(result)
    r.prime = 0
    r.length = 22

    for i, v := range p {
        primeSum[i+1] = primeSum[i] + v
    }

    for i := 0; i < len(p); i++ {
        for j := i-(r.length); j >= 0; j-- {
            sumOfPrimes := primeSum[i] - primeSum[j]
            length := i - j

            if sumOfPrimes > max {
                break
            }

            primeIndex := sort.Search(len(p), func(i int) bool { return p[i] >= sumOfPrimes})
            if length > r.length && primeIndex < len(p) && p[primeIndex] == sumOfPrimes {
                r.prime = sumOfPrimes
                r.length = length
            }
        }
    }

    t1 := time.Now()

    answer = fmt.Sprintf("%d with a chain length of %d", r.prime, r.length)
    runTime = t1.Sub(t0)

    return
}