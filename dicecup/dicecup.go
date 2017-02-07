package main

import "fmt"

func main() {
	var M, N int
	fmt.Scanln(&M, &N)
	maxResult := M + N
	resultSlice := make([]int, maxResult+1, maxResult+1)
	for i := 1; i <= M; i++ {

		for j := 1; j <= N; j++ {
			resultSlice[i+j]++
		}
	}
	for result := range largest(resultSlice) {
		fmt.Println(result)
	}
}

type DiceResult struct {
	occurs, value int
}

func largest(a []int) <-chan int {
	result := make(chan int)
	biggest := make([]DiceResult, 1, 4)
	smallestDR := DiceResult{0, 0}
	biggest = append(biggest, smallestDR)
	for i, occurs := range a {
		if occurs > biggest[0].occurs {
			biggest = nil
			biggest = append(biggest, DiceResult{occurs, i})
		} else if occurs == biggest[0].occurs {
			biggest = append(biggest, DiceResult{occurs, i})
		}
	}
	go func(c chan<- int) {
		defer close(c)
		for _, dr := range biggest {
			if dr.occurs == 0 {
				break
			}
			c <- dr.value
		}
	}(result)
	return result
}
