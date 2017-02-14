package main

import "fmt"

var cost float64
var N int
var w, l float64

func main() {
	fmt.Scanln(&cost)
	fmt.Scanln(&N)

	var total float64
	for i := 0; i < N; i++ {
		fmt.Scanln(&w, &l)
		area := w * l
		total += area * cost
	}

	fmt.Printf("%.7f\n", total)
}
