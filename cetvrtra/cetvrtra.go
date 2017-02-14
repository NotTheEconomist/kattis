package main

import "fmt"

func main() {
	var X, Y, X2, Y2, X3, Y3 int
	fmt.Scanln(&X, &Y)
	fmt.Scanln(&X2, &Y2)
	fmt.Scanln(&X3, &Y3)

	var resultX, resultY int
	switch {
	case X == X2:
		resultX = X3
	case X == X3:
		resultX = X2
	case X2 == X3:
		resultX = X
	}
	switch {
	case Y == Y2:
		resultY = Y3
	case Y == Y3:
		resultY = Y2
	case Y2 == Y3:
		resultY = Y
	}
	fmt.Printf("%d %d", resultX, resultY)
}
