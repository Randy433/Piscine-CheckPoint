package main

import "fmt"

func CanJump(a []uint) bool {
	n := len(a)

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	currentIndex := 0
	for {
		steps := a[currentIndex]

		nextIndex := currentIndex + int(steps)

		if nextIndex == n-1 {
			return true
		}

		if nextIndex >= n {
			return false
		}

        if steps == 0 && nextIndex < n-1 {
            return false
        }

		currentIndex = nextIndex
	}
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}