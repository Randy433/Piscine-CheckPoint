package main

import "fmt"

func CountCharacter(str string) int {
	counter := 0
	for _, c := range str {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(CountCharacter("Toyyib"))
}
