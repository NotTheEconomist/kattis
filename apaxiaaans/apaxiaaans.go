package main

import "fmt"

func main() {
	var inName string
	fmt.Scanln(&inName)

	var lastLetter rune

	for _, r := range inName {
		if r != lastLetter {
			lastLetter = r
			fmt.Printf("%c", r)
		}
	}
}
