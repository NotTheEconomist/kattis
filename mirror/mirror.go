package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		swap := len(runes) - i - 1
		runes[i], runes[swap] = runes[swap], runes[i]
	}
	return string(runes)
}

func main() {
	var T int
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		fmt.Printf("Test %d\n", i+1)
		var R, C int
		fmt.Scanln(&R, &C)
		input := make([]string, R)
		for r := R - 1; r >= 0; r-- {
			var line string
			fmt.Scanln(&line)
			input[r] = Reverse(line)
		}
		fmt.Println(strings.Join(input, "\n"))
	}
}
