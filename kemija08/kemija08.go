package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var Pattern = regexp.MustCompile("([aeiou])p\\w")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	crypted := scanner.Text()
	plaintext := Pattern.ReplaceAllStringFunc(crypted, func(s string) string { return string(s[0]) })
	fmt.Println(plaintext)
}
