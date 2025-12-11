package main

import "fmt"

func CamelToSnakeCase(s string) string {
	output := ""
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			return s
		}
		if i+1 < len(s) && s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
		if !(s[i] >= 'A' && s[i] <= 'Z') && !(s[i] >= 'a' && s[i] <= 'z') {
			return s
		}
		if s[i] >= 'A' && s[i] <= 'Z' && i == 0 {
			output += string(s[i])
			continue
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			output += "_" + string(s[i])
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			output += string(s[i])
			continue
		}
	}
	return output
}
