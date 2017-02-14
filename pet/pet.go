package main

import "fmt"

func MaxIntSliceIdx(s []int) (indexOfMax int) {
	biggest := s[0]
	indexOfMax = 0
	for i, v := range s {
		if v > biggest {
			indexOfMax = i
			biggest = v
		}
	}
	return
}

func main() {
	game := make([]int, 5)
	for i := range game {
		var a, b, c, d int
		fmt.Scanln(&a, &b, &c, &d)
		game[i] = a + b + c + d
	}
	resultIdx := MaxIntSliceIdx(game)
	fmt.Println(resultIdx+1, game[resultIdx])
}
