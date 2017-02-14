package main

import "fmt"

var H, M int

func main() {
	fmt.Scanln(&H, &M)
	M -= 45 // subtract 45 minutes
	if M < 0 {
		M += 60 // handle minute wrap
		H -= 1  // subtract an hour in that case
		if H < 0 {
			H += 24 // handle hour wrap
		}
	}
	fmt.Printf("%d %d", H, M)
}
