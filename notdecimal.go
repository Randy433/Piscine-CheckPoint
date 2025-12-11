package main

import (
	"fmt"
	"strconv"
)

func NotDecimal(s string) string {
	if s == "" {
		return "\n"
	}

	// Check if s has a decimal point
	dot := -1
	for i, ch := range s {
		if ch == '.' {
			dot = i
			break
		}
	}

	// No decimal point → return as is
	if dot == -1 {
		return s + "\n"
	}

	intPart := s[:dot]
	decPart := s[dot+1:]

	// If decimal part is only "0" → return integer part
	if decPart == "0" {
		return intPart + "\n"
	}

	// Must ensure it's actually a number
	// Try parsing as float
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return s + "\n"
	}

	// Remove decimal: integer = intPart + decPart
	result := intPart + decPart

	// Ensure result is numeric (avoid cases like "." , "-.")
	if _, err := strconv.Atoi(result); err != nil {
		return s + "\n"
	}

	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("12.34")) // 1234
	fmt.Print(NotDecimal("10.0"))  // 10
	fmt.Print(NotDecimal("12"))    // 12
	fmt.Print(NotDecimal(""))      // (newline)
}
