package main

import (
	"fmt"
	"strconv"
)

var L, D, X int

func main() {
	fmt.Scanln(&L)
	fmt.Scanln(&D)
	fmt.Scanln(&X)

	var M, N int
	// N is min L<=N<=D where digits of N sum to X
	// M is max L<=M<=D where digits of M sum to X

	for i := L; i <= D; i++ {
		sum := 0
		for _, ch := range strconv.Itoa(i) {
			sum += int(ch - '0')
		}
		if sum == X {
			N = i
			break
		}
	}

	for i := D; i >= L; i-- {
		sum := 0
		for _, ch := range strconv.Itoa(i) {
			sum += int(ch - '0')
		}
		if sum == X {
			M = i
			break
		}
	}
	fmt.Println(N)
	fmt.Println(M)
}
