package main

import (
	"fmt"
	// "strconv"
)

func Itoa1(n int) string {
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

func FromTo(from int, to int) string {
	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	if from < to {
		// FORWARD LOOP: Appends comma after the number, but skips the last one (i == to)
		for i := from; i <= to; i++ {
			// 1. Padding Logic (Applies if i is a single digit)
			if i <= 9 {
				result += "0"
			}
			result += Itoa1(i)

			// 2. Comma Logic (Appends if it's NOT the last number)
			if i < to {
				result += ", "
			}
		}
	} else if from > to {
		// BACKWARD LOOP: Appends comma after the number, but skips the last one (i == to)
		for i := from; i >= to; i-- {
			// 1. Padding Logic (Applies if i is a single digit)
			if i <= 9 {
				result += "0"
			}
			result += Itoa1(i)

			// 2. Comma Logic (Appends if it's NOT the last number)
			// The condition is i > to because we are counting down.
			if i > to {
				result += ", "
			}
		}
	} else {
		// Case: from == to (Single number sequence)
		if from <= 9 {
			result += "0"
		}
		result += Itoa1(from)
	}

	// Always add the trailing newline at the end of the final result string
	return result + "\n"
}

func main() {
	// Expected: 01, 02, 03, 04, 05, 06, 07, 08, 09, 10
	fmt.Print(FromTo(1, 10))

	// Expected: 10, 09, 08, 07, 06, 05, 04, 03, 02, 01
	fmt.Print(FromTo(10, 1))

	// Expected: 10
	fmt.Print(FromTo(10, 10))

	// Expected: Invalid\n
	fmt.Print(FromTo(100, 10))
}
