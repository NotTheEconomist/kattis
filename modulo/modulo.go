package main

import "fmt"

func main() {
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scanln(&numbers[i])
	}
	uniqueSet := make(map[int]bool)
	for _, n := range numbers {
		uniqueSet[n%42] = true
	}
	fmt.Println(len(uniqueSet))
}
