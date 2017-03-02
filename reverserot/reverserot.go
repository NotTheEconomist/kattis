package main

import "fmt"

func rot(N int, ch byte) byte {
	switch ch {
	case '_':
		ch = 'Z' + 1
	case '.':
		ch = 'Z' + 2
	}
	newCh := (ch - 'A' + byte(N)) % 28
	switch newCh {
	case 26:
		return byte('_')
	case 27:
		return byte('.')
	}
	return newCh + 'A'
}

func rotBytes(N int, ba []byte) []byte {
	result := make([]byte, len(ba))
	for i := 0; i < len(ba); i++ {
		result[i] = rot(N, ba[i])
	}
	return result
}

func reverseBytes(ba []byte) []byte {
	length := len(ba)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = ba[length-i-1]
	}
	return result
}

func main() {
	for {
		var N int
		fmt.Scan(&N)
		if N == 0 {
			break
		}
		var plaintext []byte
		fmt.Scan(&plaintext)
		fmt.Println(string(reverseBytes(rotBytes(N, plaintext))))
	}
}
