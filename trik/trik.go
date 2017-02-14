package main

import "fmt"

var where = 1

func swap(code rune) {
	var a, b int
	switch code {
	case 'A':
		a, b = 1, 2
	case 'B':
		a, b = 2, 3
	case 'C':
		a, b = 1, 3
	}
	if a == where {
		where = b
	} else if b == where {
		where = a
	}
}

func main() {
	var moves string
	fmt.Scan(&moves)

	for _, r := range moves {
		swap(r)
	}
	fmt.Println(where)
}
