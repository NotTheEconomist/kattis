package main

import (
	"fmt"
	"math"
)

const g float64 = 9.81

func main() {
	var N int
	fmt.Scanln(&N)
	var v0, theta, x, h1, h2 float64
	for i := 0; i < N; i++ {
		fmt.Scanln(&v0, &theta, &x, &h1, &h2)
		theta = theta * math.Pi / 180 // angle -> radian
		// x(t) = v0tcosθ
		t := x / (v0 * math.Cos(theta)) // time t where x(t) == h1
		// y(t) = v0tsinθ−12gt2
		actualH := v0*t*math.Sin(theta) - (0.5)*g*t*t
		if actualH <= h2-1 && actualH >= h1+1 {
			fmt.Println("Safe")
		} else {
			fmt.Println("Not Safe")
		}
	}
}
