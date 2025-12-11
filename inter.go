package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""

	for _, c := range s1 {
		if contains(s2, c) && !contains(res, c) {
			res += string(c)
		}
	}

	fmt.Println(res)
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
