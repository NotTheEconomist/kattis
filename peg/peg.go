package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// Graph describes a peg game
type Graph interface {

	// ValidMovesFrom shows how many valid moves are available that will
	// finish in the empty space defined by (x, y int)
	ValidMovesFrom(x, y int) int

	// OnGraph checks if a point x, y is on the graph
	OnGraph(x, y int) bool

	// Iterator provides a <-chan GraphIter to range over
	Iterator() <-chan GraphIter
}

type GraphIter struct {
	X, Y int
	R    rune
}

type graph [][]rune

// transformFactory returns a function with signature F(x, y, k int) (int int)
// dX and dY are both expected to be the runes '+', '-', or '0'
// which indicate which direction x or y will be transformed in F.
// F returns the transformation across k units.
//
// Example:
//     NewFunc := transformFactory('-', '')
//     x, y := 1, 0
//     newX, newY := NewFunc(x, y, 1)
//     fmt.Println(newX, newY)  // 0 0
func transformFactory(dX, dY rune) func(x, y, k int) (int, int) {
	var xMod, yMod int
	switch dX {
	case '+':
		xMod = 1
	case '-':
		xMod = -1
	case '0':
		xMod = 0
	default:
		panic("Incorrect rune set as dX")
	}
	switch dY {
	case '+':
		yMod = 1
	case '-':
		yMod = -1
	case '0':
		yMod = 0
	default:
		panic("Incorrect rune set as dY")
	}
	f := func(x, y, k int) (int, int) {
		return x + (k * xMod), y + (k * yMod)
	}
	return f
}

// TODO: fixme! Not detecting valid move, but not exiting early
func (g graph) ValidMovesFrom(x, y int) int {
	destination := g[y][x]
	fmt.Printf("The rune at %d, %d is %c\n", x, y, destination)
	if destination != EMPTY {
		return 0
	}
	moves := 0
	for _, t := range Directions {
		originX, originY := t(x, y, 2)
		intermediaryX, intermediaryY := t(x, y, 1)
		if g.OnGraph(originX, originY) && g.OnGraph(intermediaryX, intermediaryY) &&
			g[originY][originX] == PEG && g[intermediaryY][intermediaryX] == PEG {
			moves++
		}
	}
	return moves
}

func (g graph) OnGraph(x, y int) bool {
	return !(0 <= x && x <= 7 || 0 <= y && y <= 7)
}

func (g graph) Iterator() <-chan GraphIter {
	c := make(chan GraphIter)
	go func(g graph, c chan<- GraphIter) {
		defer close(c)
		for i := 0; i < len(g); i++ {
			for j := 0; j < len(g[i]); j++ {
				c <- GraphIter{X: j, Y: i, R: g[i][j]}
			}
		}
	}(g, c)
	return c
}

const (
	WALL  = ' '
	PEG   = 'o'
	EMPTY = '.'
)

var (
	Left       = transformFactory('-', '0')
	Right      = transformFactory('+', '0')
	Up         = transformFactory('0', '-')
	Down       = transformFactory('0', '+')
	Directions = []func(int, int, int) (int, int){Left, Right, Up, Down}
)

// ValidMoves returns the number of valid moves in Graph g
func ValidMoves(g Graph) int {
	total := 0
	for gi := range g.Iterator() {
		x, y, r := gi.X, gi.Y, gi.R
		if r == EMPTY {
			fmt.Println("Found an empty spot!")
			total += g.ValidMovesFrom(x, y)
		}
	}
	return total
}

func main() {
	g := make(graph, 7)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(g); i++ {
		scanner.Scan()
		line := bytes.Runes(scanner.Bytes())
		g[i] = line
	}
	numMoves := ValidMoves(g)
	fmt.Println(numMoves)
}
