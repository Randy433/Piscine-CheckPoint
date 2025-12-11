package main

func Fibonacci(index int) int {
	// 1. Handle Negative Index (Requirement)
	if index < 0 {
		return -1
	}
	// 2. Handle Index 0 (Base Case 1)
	if index == 0 {
		return 0
	}
	// 3. Handle Index 1 (Base Case 2)
	if index == 1 {
		return 1
	}

	return Fibonacci(index-1) + Fibonacci(index-2)

}
