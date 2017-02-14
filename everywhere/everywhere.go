package main

import "fmt"

var numCases int

func main() {
	fmt.Scanln(&numCases)
	for i := 0; i < numCases; i++ {
		cities := make(map[string]bool) // set
		var numCities int
		fmt.Scanln(&numCities)
		for j := 0; j < numCities; j++ {
			var cityName string
			fmt.Scanln(&cityName)
			cities[cityName] = true
		}
		fmt.Println(len(cities))
	}
}
