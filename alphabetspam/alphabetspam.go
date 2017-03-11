package main

import "fmt"

var (
	underscores = map[rune]bool{'_': false}
	lowercases  = map[rune]bool{'a': false,
		'b': false,
		'c': false,
		'd': false,
		'e': false,
		'f': false,
		'g': false,
		'h': false,
		'i': false,
		'j': false,
		'k': false,
		'l': false,
		'm': false,
		'n': false,
		'o': false,
		'p': false,
		'q': false,
		'r': false,
		's': false,
		't': false,
		'u': false,
		'v': false,
		'w': false,
		'x': false,
		'y': false,
		'z': false,
	}
	uppercases = map[rune]bool{'A': false,
		'B': false,
		'C': false,
		'D': false,
		'E': false,
		'F': false,
		'G': false,
		'H': false,
		'I': false,
		'J': false,
		'K': false,
		'L': false,
		'M': false,
		'N': false,
		'O': false,
		'P': false,
		'Q': false,
		'R': false,
		'S': false,
		'T': false,
		'U': false,
		'V': false,
		'W': false,
		'X': false,
		'Y': false,
		'Z': false,
	}
)

func main() {
	var s string
	fmt.Scanln(&s)

	values := map[string]int{"underscores": 0, "lowercases": 0, "uppercases": 0, "symbols": 0}

	for _, ch := range s {
		r := rune(ch)
		if _, ok := underscores[r]; ok {
			values["underscores"]++
		} else if _, ok := lowercases[r]; ok {
			values["lowercases"]++
		} else if _, ok := uppercases[r]; ok {
			values["uppercases"]++
		} else {
			values["symbols"]++
		}
	}

	fmt.Printf("%.12f\n", float64(values["underscores"])/float64(len(s)))
	fmt.Printf("%.12f\n", float64(values["lowercases"])/float64(len(s)))
	fmt.Printf("%.12f\n", float64(values["uppercases"])/float64(len(s)))
	fmt.Printf("%.12f\n", float64(values["symbols"])/float64(len(s)))
}
