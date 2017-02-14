package main

import "fmt"

var scoreTable = map[byte]int{
	'A': 11, 'K': 4, 'Q': 3, 'J': 2,
	'T': 10, '9': 0, '8': 0, '7': 0}

func main() {
	var numHands int
	var tmpDomSuit string

	fmt.Scanln(&numHands, &tmpDomSuit)
	domSuit := tmpDomSuit[0]
	totalScore := 0
	for i := 0; i < 4*numHands; i++ {
		var card string
		fmt.Scanln(&card)
		totalScore += scoreTable[card[0]]
		if card[1] == domSuit {
			switch card[0] {
			case '9':
				totalScore += 14
			case 'J':
				totalScore += 18
			}
		}
	}
	fmt.Println(totalScore)
}
