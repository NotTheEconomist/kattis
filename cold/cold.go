package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var a []int
	fmt.Scanln(&a)
	fmt.Println(n, a)

	result := 0

	// for _, temp := range t {
	// 	if temp < 0 {
	// 		result++
	// 	}
	// }

	fmt.Println(result)
}
