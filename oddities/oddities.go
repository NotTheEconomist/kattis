package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scanln(&num)
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
