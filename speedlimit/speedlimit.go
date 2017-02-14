package main

import "fmt"

func main() {
	var lines int
	for {
		fmt.Scanln(&lines)
		if lines == -1 {
			break
		}
		distance := 0
		var speed, elapsed, lastElapsed int
		for i := 0; i < lines; i++ {
			fmt.Scanln(&speed, &elapsed)
			distance += speed * (elapsed - lastElapsed)
			lastElapsed = elapsed
		}
		fmt.Printf("%d miles\n", distance)
	}
}
