package main

import "fmt"

func main() {

	var N int

	for count := 1; true; count++ {
		fmt.Scanln(&N)
		if N == 0 {
			break
		}
		names := make([]string, N)
		resortedNames := make([]string, 0, N)
		for i := 0; i < N; i++ {
			fmt.Scanln(&names[i])
		}

		var start int
		if len(names)%2 == 0 {
			start = N - 1
		} else {
			start = N - 2
		}

		for i := 0; i < N; i += 2 {
			resortedNames = append(resortedNames, names[i])
		}
		for i := start; i > 0; i -= 2 {
			resortedNames = append(resortedNames, names[i])
		}

		fmt.Printf("SET %d\n", count)
		for _, name := range resortedNames {
			fmt.Println(name)
		}
	}
}
