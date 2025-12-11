package main

import "fmt"

func RepeatAlpha(s string) string {
	result := ""

	for _, c := range s {
		var count int

		// Check for lowercase letter
		if c >= 'a' && c <= 'z' {
			count = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			count = int(c - 'A' + 1)
		} else {
			// Not an alphabet → repeat once
			count = 1
		}

		for i := 0; i < count; i++ {
			result += string(c)
		}
	}

	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha("toyyib"))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
