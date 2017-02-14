package main

import "fmt"

var P int

func main() {
	fmt.Scanln(&P, &P) // ignore first number

	fmt.Println(P)
}
