package main

import "fmt"

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FindPrevPrime returns the first prime number ≤ n
func FindPrevPrime(n int) int {
	for i := n; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(11)) // 11
	fmt.Println(FindPrevPrime(10)) // 7
	fmt.Println(FindPrevPrime(1))  // 0
	fmt.Println(FindPrevPrime(20)) // 19
	fmt.Println(FindPrevPrime(3))  // 3
}
