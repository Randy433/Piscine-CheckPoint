package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}

	s1 := args[1]
	s2 := args[2]

	if s1 == "" {
		fmt.Println("1")
		return
	}

	i := 0
	for j := 0; j < len(s2); j++ {
		if i < len(s1) && s1[i] == s2[j] {
			i++
		}
		if i == len(s1) {
			fmt.Println("1")
			return
		}
	}

	fmt.Println("0")
}
