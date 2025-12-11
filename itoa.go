package main

import "fmt"

// Itoa simulates the behavior of strconv.Itoa, converting an integer
// to its string representation while correctly handling the sign.
func Itoa(n int) string {
	// 1. Handle the zero case immediately
	if n == 0 {
		return "0"
	}

	// 2. Determine the sign and work with the absolute value
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	// 3. Extract digits (they will be collected in reverse order)
	digits := []byte{}

	for n > 0 {
		// Get the last digit (remainder)
		digit := n % 10

		// Convert the digit (0-9) to its ASCII character ('0' is ASCII 48)
		digits = append(digits, byte(digit)+'0')

		// Remove the last digit
		n /= 10
	}

	// 4. Reverse the collected digits
	// We use the two-pointer swap method for in-place reversal.
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// 5. Prepend the sign if necessary
	if isNegative {
		return "-" + string(digits)
	}

	// Convert the byte slice of digits to a final string
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
