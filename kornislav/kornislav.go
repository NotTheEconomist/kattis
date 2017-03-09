package main

import (
	"fmt"
	"sort"
)

func main() {
	distances := make([]int, 4)
	for i := range distances {
		var A int
		fmt.Scan(&A)
		distances[i] = A
	}
	sort.Ints(distances)

	fmt.Println(distances[0] * distances[2])
}
