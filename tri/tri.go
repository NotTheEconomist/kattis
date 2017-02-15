package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	var firstOp, secondOp string

	switch {
	case a+b == c:
		firstOp = "+"
		secondOp = "="
	case a*b == c:
		firstOp = "*"
		secondOp = "="
	case a-b == c:
		firstOp = "-"
		secondOp = "="
	case a/b == c:
		firstOp = "/"
		secondOp = "="
	case a == b+c:
		firstOp = "="
		secondOp = "+"
	case a == b*c:
		firstOp = "="
		secondOp = "*"
	case a == b-c:
		firstOp = "="
		secondOp = "-"
	case a == b/c:
		firstOp = "="
		secondOp = "/"
	}
	fmt.Printf("%d%s%d%s%d", a, firstOp, b, secondOp, c)
}
