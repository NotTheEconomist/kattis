package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	minutes int  // How far into the competition the result was submitted
	problem rune // single rune identifying the problem. 'A', 'B', 'E', etc.
	correct bool // submitted as "right" or "wrong," will be correctly parsed by
	// Result.Parse(s string)
}

func (r *Result) Parse(s string) {

	sa := strings.Split(s, " ")
	minutes, err := strconv.Atoi(sa[0])
	if err != nil {
		panic("Bad input" + sa[0])
	}
	problem := rune(sa[1][0])
	var correct bool

	if sa[2] == "right" {
		correct = true
	} else {
		correct = false
	}

	r.minutes = minutes
	r.problem = problem
	r.correct = correct
}

func New(s string) Result {
	r := new(Result) // pointer
	r.Parse(s)
	return *r
}

func main() {

	results := make([]Result, 0)
	for {
		a := make([]string, 3)
		fmt.Scan(&a[0])
		if a[0] == "-1" {
			break
		}
		fmt.Scan(&a[1])
		fmt.Scan(&a[2])
		result := New(strings.Join(a, " "))
		results = append(results, result)
	}
	// results is now full of populated Result objects

	wrongAnswers := make(map[rune]int)
	solvedProblems := make(map[rune]bool)
	score := 0

	for _, r := range results {
		if r.correct {
			if _, ok := solvedProblems[r.problem]; !ok {
				score += r.minutes
				solvedProblems[r.problem] = true
			}
		} else {
			wrongAnswers[r.problem]++
		}
	}
	for problem, count := range wrongAnswers {
		if _, ok := solvedProblems[problem]; ok {
			score += (count * 20)
		}
	}

	fmt.Printf("%d %d\n", len(solvedProblems), score)

}
