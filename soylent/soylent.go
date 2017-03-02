package main

import "fmt"

const CAL_PER_BOTTLE int = 400

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var dailyCal int
		fmt.Scan(&dailyCal)
		numBottles := dailyCal / CAL_PER_BOTTLE
		if numBottles*CAL_PER_BOTTLE < dailyCal {
			numBottles++
		}
		fmt.Println(numBottles)
	}
}
