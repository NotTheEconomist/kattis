package main

import "fmt"

func main() {
	var P int
	fmt.Scan(&P)

	for K := 1; K < P+1; K++ {
		var unused, N int
		fmt.Scan(&unused)
		fmt.Scan(&N)
		var sumFirstN, sumOddN, sumEvenN int
		for n := 1; n <= N; n++ {
			sumFirstN += n
			sumOddN += (n-1)*2 + 1
			sumEvenN += n * 2
		}
		fmt.Printf("%d %d %d %d\n", K, sumFirstN, sumOddN, sumEvenN)
	}
}
