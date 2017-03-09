package main

import (
	"fmt"
	"sort"
)

func adrianGuess(n int) rune {
	return [...]rune{'A', 'B', 'C'}[n%3]
}

func brunoGuess(n int) rune {
	return [...]rune{'B', 'A', 'B', 'C'}[n%4]
}

func goranGuess(n int) rune {
	return [...]rune{'C', 'A', 'B'}[(n%6)/2]
}

type Applicant struct {
	Guess func(int) rune
	Name  string
}

var (
	Adrian = Applicant{Guess: adrianGuess, Name: "Adrian"}
	Bruno  = Applicant{Guess: brunoGuess, Name: "Bruno"}
	Goran  = Applicant{Guess: goranGuess, Name: "Goran"}
)

func main() {
	var N int
	var key []byte

	fmt.Scan(&N) // ignored
	fmt.Scan(&key)

	winCounts := make(map[string]int)
	applicants := []Applicant{Adrian, Bruno, Goran}

	for i, b := range key {
		for _, applicant := range applicants {
			guess := applicant.Guess(i)
			if guess == rune(b) {
				winCounts[applicant.Name]++
			}
		}
	}
	var maxWins int
	for _, wins := range winCounts {
		if wins > maxWins {
			maxWins = wins
		}
	}
	winners := make([]string, 0, len(applicants))
	for name, wins := range winCounts {
		if wins == maxWins {
			winners = append(winners, name)
		}
	}
	sort.Strings(winners)

	fmt.Println(maxWins)

	for _, name := range winners {
		fmt.Println(name)
	}
}
