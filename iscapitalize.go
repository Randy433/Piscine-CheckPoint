package main

import "fmt"

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	startOfWord := true // To track when a new word starts

	for _, c := range s {
		if startOfWord {
			// If word starts with a lowercase letter → false
			if c >= 'a' && c <= 'z' {
				return false
			}
			startOfWord = false
		}

		// Word separators: spaces or tabs
		if c == ' ' || c == '\t' {
			startOfWord = true
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
