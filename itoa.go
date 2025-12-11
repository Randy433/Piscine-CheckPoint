package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	digits := []byte{}

	for n > 0 {
		digit := n % 10

		digits = append(digits, byte(digit)+'0')

		n /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if isNegative {
		return "-" + string(digits)
	}

	return string(digits)
}

func main() {
	// Test cases
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(-987))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(500))
	fmt.Println(Itoa(-1))
}
