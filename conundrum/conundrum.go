package main

import "fmt"

func main() {
	var ciphertext string
	fmt.Scanln(&ciphertext)

	discount := 0
	needleString := "PER"

	for offset := 0; offset < 3; offset++ {
		needle := needleString[offset]
		for i := offset; i < len(ciphertext); i += 3 {
			if ciphertext[i] == needle {
				discount++
			}
		}
	}
	fmt.Println(len(ciphertext) - discount)
}
