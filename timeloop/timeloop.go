package main

import "fmt"

var N int

func main() {
	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("%d Abracadabra\n", i)
	}
}
