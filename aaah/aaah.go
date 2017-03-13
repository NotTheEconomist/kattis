package main

import "fmt"

func main() {
	var got, want string
	fmt.Scanln(&got)
	fmt.Scanln(&want)

	if len(got) >= len(want) {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}
}
