package main

import (
	"fmt"
	"math"
)

var h, v int

func main() {
	// math.Sin(vAsRad) == h / ladderLength
	fmt.Scanln(&h, &v)
	vAsRad := float64(v) * math.Pi / 180.0
	ladderLength := int(math.Ceil(float64(h) / math.Sin(vAsRad)))
	fmt.Println(ladderLength)
}
