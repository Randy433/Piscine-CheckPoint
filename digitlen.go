package main

import "fmt"

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	counter := 0
	// var divisor int
	for n != 0 {
		n = n / base
		counter++
	}
	return counter
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
