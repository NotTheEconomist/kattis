package main

import (
	"fmt"
	"strconv"
)

var N int

func main() {
	fmt.Scanln(&N)
	bRep := fmt.Sprintf("%b", N)
	bLen := len(bRep)
	backwards := make([]rune, bLen)
	for i, r := range bRep {
		backwards[bLen-i-1] = r
	}
	result, _ := strconv.ParseInt(string(backwards), 2, 0)
	fmt.Println(result)
}
