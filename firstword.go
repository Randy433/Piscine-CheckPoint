package main

import "fmt"

func FirstWord(s string) string {
	runes := []rune{}
	wordFound := false

	if len(s) == 0 {
		return "" + "\n"
	}
	for _, c := range s {
		if c != ' ' {
			runes = append(runes, c)
			wordFound = true
		} else if wordFound {
			// A word was found but immediately after a space
			return string(runes) + "\n"
		}
	}
	return string(runes) + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
