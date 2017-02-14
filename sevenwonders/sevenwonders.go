package main

import "fmt"

type Game struct {
	numTablet, numCompass, numGear int
}

func (g Game) Score() int {
	score := g.numTablet*g.numTablet +
		g.numCompass*g.numCompass +
		g.numGear*g.numGear
	leastCount := 51 // by definition, a 50 card game cannot have this many cards
	for _, count := range [...]int{g.numTablet, g.numCompass, g.numGear} {
		if count < leastCount {
			leastCount = count
		}
	}
	score += leastCount * 7
	return score
}

func main() {
	var in string

	fmt.Scanln(&in)
	g := Game{}
	for _, ch := range in {
		switch ch {
		case 'T':
			g.numTablet++
		case 'C':
			g.numCompass++
		case 'G':
			g.numGear++
		}
	}
	fmt.Println(g.Score())
}
