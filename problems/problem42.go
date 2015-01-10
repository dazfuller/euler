/*
// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle
// numbers are:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// By converting each letter in a word to a number corresponding to its alphabetical position and adding
// these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If
// the word value is a triangle number then we shall call the word a triangle word.
*/

package problems

import (
	"fmt"
	"github.com/dazfuller/euler/eulermath"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Problem42 struct{}

func sumOfWord(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) - 64
	}
	return sum
}

func (p *Problem42) Solve() (answer string, runTime time.Duration) {
	t0 := time.Now()

	byteContent, err := ioutil.ReadFile("problems/data/p042_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(byteContent[1:len(byteContent)-1]), "\",\"")
	count := 0

	for _, word := range words {
		if eulermath.IsTriangular(sumOfWord(word)) {
			count++
		}
	}

	t1 := time.Now()

	answer = fmt.Sprint(count)
	runTime = t1.Sub(t0)

	return
}
