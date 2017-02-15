package main

import "fmt"

func sumDigits(num int) (sum int) {
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return
}

func main() {
	var N int
	for {
		fmt.Scanln(&N)
		if N == 0 {
			return
		}
		target := sumDigits(N)
		for p := 11; true; p++ {
			if sumDigits(N*p) == target {
				fmt.Println(p)
				break
			}
		}
	}
}
