package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64 // radius of circles
	fmt.Scanln(&R)

	euclideanArea := R * R * math.Pi
	taxicabArea := R * R * 2

	fmt.Printf("%.6f\n", euclideanArea)
	fmt.Printf("%.6f\n", taxicabArea)
}
