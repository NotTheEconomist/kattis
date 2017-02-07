package main

import "fmt"

func main() {
	var count int
	var n int
	fmt.Scanln(&n)

	temps := make([]int, n)
	for i := range temps {
		_, err := fmt.Scan(&temps[i])
		if err != nil {
			panic("Failed to read")
		}
	}

	for _, temp := range temps {
		if temp < 0 {
			count++
		}
	}
	fmt.Println(count)
}
