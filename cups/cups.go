package main

import (
	"fmt"
	"sort"
	"strconv"
)

var N int // how many cups
var Cups = make([]Cup, 0, N)

type Cup struct {
	radius int
	color  string
}

func (c *Cup) String() string {
	return c.color
}

type ByRadius []Cup

func (c ByRadius) Len() int           { return len(c) }
func (c ByRadius) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByRadius) Less(i, j int) bool { return c[i].radius < c[j].radius }

func main() {
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var a, b string
		fmt.Scanln(&a, &b)
		var r int
		var color string
		var err error
		if r, err = strconv.Atoi(a); err != nil {
			r, _ = strconv.Atoi(b)
			color = a
		} else {
			r /= 2
			color = b
		}
		Cups = append(Cups, Cup{radius: r, color: color})
	}
	sort.Sort(ByRadius(Cups))
	for _, c := range Cups {
		fmt.Println(c.String())
	}
}
