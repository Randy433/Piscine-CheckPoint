package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Validation: must be a proper camelCase
	n := len(s)
	if !(s[0] >= 'A' && s[0] <= 'Z' || s[0] >= 'a' && s[0] <= 'z') {
		return s
	}
	// must not end with an uppercase letter
	if s[n-1] >= 'A' && s[n-1] <= 'Z' {
		return s
	}

	prevUpper := false
	for i := 0; i < n; i++ {
		ch := s[i]
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			return s
		}
		if ch >= 'A' && ch <= 'Z' {
			if prevUpper {
				return s
			}
			prevUpper = true
		} else {
			prevUpper = false
		}
	}

	// Build result with underscores
	result := ""
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' && i > 0 {
			result += "_"
		}
		result += string(ch)
	}

	return result
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}