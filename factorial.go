package main

func factorial(a int) int {
	b := 1
	if a == 0 {
		return 1
	}
	if a > 0 {
		for i := a; i > 0; i-- {
			b *= i
		}
		return b
	}
	return 0
}
