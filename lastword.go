package main

import "fmt"

func LastWord(s string) string {
	runes := []rune{}
	wordFound := false

	// check for length 0 for the string
	if len(s) == 0 {
		return "" + "\n"
	}

	// Start from the end and move backward
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if c != ' ' {
			// Build the word in reverse order
			runes = append(runes, c)
			wordFound = true
		} else if wordFound {
			// We've already collected a word, and now found a space → stop
			break
		}
	}

	// Now reverse the collected runes to get the correct order
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) + "\n"
}

func main() {
	fmt.Print(LastWord("hello world"))
	fmt.Print(LastWord("   go   "))
	fmt.Print(LastWord("singleword"))
	fmt.Print(LastWord(" multiple words in here "))
}
