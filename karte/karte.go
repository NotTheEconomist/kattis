package main

import (
	"fmt"
	"strconv"
)

var (
	suits = map[rune]int{'P': 13, 'K': 13, 'H': 13, 'T': 13}
	seen  = map[Card]bool{}
)

type Card struct {
	rank int
	suit rune
}

func main() {
	var hand string
	fmt.Scanln(&hand)

	var error bool

	for i := 0; i < len(hand); i += 3 {
		suit := rune(hand[i])
		rankStr := hand[i+1 : i+3] // "01" - "13"
		rank, ok := strconv.Atoi(rankStr)
		if ok != nil {
			error = true
			break
		}
		c := Card{rank: rank, suit: suit}
		if _, ok := seen[c]; ok {
			// duplicate card
			error = true
			break
		} else {
			seen[c] = true
		}
		suits[suit]--
	}
	if !error {
		fmt.Println(suits['P'], suits['K'], suits['H'], suits['T'])
	} else {
		fmt.Println("GRESKA")
	}
}
