package main

func CountAlpha(str string) int {
	counter := 0
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			counter++
		}
	}
	return counter
}
