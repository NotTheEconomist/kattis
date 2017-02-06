package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)
	if N%2 == 1 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
