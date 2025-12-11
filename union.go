package main

import (
	"fmt"
	"os"
)

func contian(str string, c rune) bool  {
	for _, ch := range str {
		if ch == c {
			return true
		}
	}
	return false
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""

	// add from s1
	for _, c := range s1 {
		if !contian(res, c) {
			res += string(c)
		}
	}

	// add from s2
	for _, c := range s2 {
		if !contian(res, c) {
			res += string(c)
		}
	}
	fmt.Println(res)
}