package main

import "fmt"

var scoreTable = map[rune]int{
	'A': 11, 'K': 4, 'Q': 3, 'J': 2,
	'T': 10, '9': 0, '8': 0, '7': 0}

type Card struct {
	rank, suit rune
}

func (c Card) Score(domSuit rune) (score int) {
	score = scoreTable[c.rank]
	if c.suit == domSuit {
		switch c.rank {
		case 'J':
			score = 20
		case '9':
			score = 14
		}
	}
	return
}

type Hand []Card

func (h Hand) Score(domSuit rune) (score int) {
	for _, c := range h {
		score += c.Score(domSuit)
	}
	return
}

func main() {
	var numHands int
	var t string

	fmt.Scanln(&numHands, &t)
	domSuit := rune(t[0])
	totalScore := 0
	for i := 0; i < numHands; i++ {
		hand := make(Hand, 4)
		for j := 0; j < 4; j++ {
			var handAsLine string
			fmt.Scanln(&handAsLine)
			hand[j] = Card{rank: rune(handAsLine[0]), suit: rune(handAsLine[1])}
		}
		totalScore += hand.Score(domSuit)
	}
	fmt.Println(totalScore)
}
