package main

import "fmt"

func main() {
	var T int // num test cases
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		lowerBound := 0
		lastPop := 1
		curPop := 1
		for curPop != 0 {
			if curPop > lastPop*2 {
				lowerBound += curPop - lastPop*2
			}
			lastPop = curPop
			fmt.Scan(&curPop)
		}
		fmt.Println(lowerBound)
	}
}
