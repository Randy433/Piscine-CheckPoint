package main

import (
	"fmt"
)

func FifthAndSkip(s string) string {
	if len(s) == 0 {
		return "\n"
	}

	// Count non-space characters
	count := 0
	for _, c := range s {
		if c != ' ' {
			count++
		}
	}

	if count < 5 {
		return "Invalid Input\n"
	}

	result := ""
	block := ""     // collects up to 5 characters
	nonSpaceCount := 0 // counts characters ignoring spaces

	for _, c := range s {
		if c == ' ' {
			continue
		}

		nonSpaceCount++

		// The 6th character must be skipped
		if nonSpaceCount%6 == 0 {
			continue
		}

		block += string(c)

		// Every time we collect 5 characters, flush them
		if len(block) == 5 {
			result += block + " "
			block = ""
		}
	}

	// Add remaining characters if any
	if block != "" {
		result += block
	}

	return result
}

func main() {
	fmt.Print(FifthAndSkip("abc de fghijklm"))
}
