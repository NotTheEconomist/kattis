package main

import "fmt"

type PlayerNum int

var (
	detTime int = 3*60 + 30 // 3 minutes 30 seconds
)

func PassLeft(from PlayerNum) PlayerNum {
	to := from + 1
	if to > 8 {
		to -= 8
	}
	return to
}

func main() {
	var activePlayer PlayerNum
	fmt.Scanln(&activePlayer)

	var numQuestions int
	fmt.Scanln(&numQuestions)

	time := 0

	for i := 0; i < numQuestions; i++ {
		var timePassed int
		var resultS string
		fmt.Scan(&timePassed, &resultS)

		time += timePassed
		if time > detTime {
			// the bomb blew up, so break and output activePlayer
			fmt.Println(activePlayer)
			break
		}

		result := rune(resultS[0])
		switch result {
		case 'T':
			activePlayer = PassLeft(activePlayer)
		case 'N', 'P':
			// nothing happens, other than time passes
		}
	}
}
