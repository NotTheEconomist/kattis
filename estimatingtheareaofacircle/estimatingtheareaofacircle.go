package main

import (
	"fmt"
	"math"
)

func main() {
	var R, marked, hit float64
	for {
		fmt.Scanln(&R, &marked, &hit)
		if R == 0 && marked == 0 && hit == 0 {
			break
		}
		calcArea := R * R * math.Pi
		estimateArea := R * R * 4 * (hit / marked)
		fmt.Printf("%f %f\n", calcArea, estimateArea)
	}
}
