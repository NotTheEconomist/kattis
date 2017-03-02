package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var R, C, Zr, Zc int
	fmt.Scan(&R, &C, &Zr, &Zc)

	data := make([][]byte, 0, R*Zr)
	for r := 0; r < R; r++ {
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	for i := 0; i < R; i++ {
		s.Scan()
		rowSlice := make([]byte, C*Zc)
		for j, b := range s.Text() {
			for offset := 0; offset < Zc; offset++ {
				rowSlice[j*Zc+offset] = byte(b)
			}
		}
		for ii := 0; ii < Zr; ii++ {
			data = append(data, rowSlice)
		}
	}

	for _, row := range data {
		fmt.Println(string(row))
	}
}
