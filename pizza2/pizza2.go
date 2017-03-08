package main

import (
	"fmt"
	"math"
)

func main() {

	var R, C float64
	fmt.Scanln(&R, &C)

	outerA := R * R * math.Pi

	innerR := R - C
	innerA := innerR * innerR * math.Pi

	fmt.Printf("%.6f\n", innerA/outerA*100)
}
