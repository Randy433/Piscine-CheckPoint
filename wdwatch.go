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

	i := 0 // index for s1

	for _, c := range s2 {
		if i < len(s1) && c == rune(s1[i]) {
			i++
		}
		if i == len(s1) {
			fmt.Println(s1)
			return
		}
	}
}
