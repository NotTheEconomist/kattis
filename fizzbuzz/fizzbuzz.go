package main

import (
	"fmt"
	"strconv"
	"strings"
)

var x, y, n int

func main() {
	fmt.Scanln(&x, &y, &n)
	for i := 1; i <= n; i++ {
		aRes := make([]string, 0, 2)
		if i%x == 0 {
			aRes = append(aRes, "Fizz")
		}
		if i%y == 0 {
			aRes = append(aRes, "Buzz")
		}
		if len(aRes) == 0 {
			aRes = append(aRes, strconv.Itoa(i))
		}
		fmt.Println(strings.Join(aRes, ""))
	}
}
