package main

import "fmt"

var N int

func main() {
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var A, B string
		fmt.Scanln(&A)
		fmt.Scanln(&B)
		resultSlice := make([]rune, len(A))
		for j := 0; j < len(A); j++ {
			if A[j] == B[j] {
				resultSlice[j] = '.'
			} else {
				resultSlice[j] = '*'
			}
		}
		fmt.Printf("%s\n%s\n%s\n", A, B, string(resultSlice))
		if i != N-1 {
			fmt.Println("") // blank line between test cases
		}
	}
}
