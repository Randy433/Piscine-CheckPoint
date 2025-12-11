package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	runes := []rune{}

	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	for i := 2; i <= len(str)-1; i += 3 {
		runes = append(runes, rune(str[i]))
	}

	return string(runes) + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
