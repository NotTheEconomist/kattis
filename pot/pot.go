package main

import (
	"fmt"
	"math"
)

var N int

func main() {
	fmt.Scanln(&N)
	total := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scanln(&a)
		pow := a % 10
		base := a / 10
		total += int(math.Pow(float64(base), float64(pow)))
	}
	fmt.Println(total)
}
