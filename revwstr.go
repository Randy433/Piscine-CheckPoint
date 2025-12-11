package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	words := splitWords(s)

	// reverse words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// print result
	for i := 0; i < len(words); i++ {
		fmt.Print(words[i])
		if i != len(words)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func splitWords(s string) []string {
	var result []string
	start := 0

	for i := 0; i <= len(s); i++ {
		// split on space or end of string
		if i == len(s) || s[i] == ' ' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}

	return result
}
