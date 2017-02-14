package main

import (
	"fmt"
	"strconv"
	"strings"
)

var expected = []int{1, 1, 2, 2, 2, 8}

func main() {
	got := make([]int, 6)
	for i := range got {
		fmt.Scan(&got[i])
	}

	want := make([]string, 6)
	for i, v := range got {
		need := expected[i] - v
		want[i] = strconv.Itoa(need)
	}
	fmt.Println(strings.Join(want, " "))
}
