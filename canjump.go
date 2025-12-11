package main

import "fmt"

func CanJump(a []uint) bool {
	// Get the length of the list
	n := len(a)

	// case 1: if n is 0 return false
	if n == 0 {
		return false
	}
	// if n is 1 return true, it means we are at the last index
	if n == 1 {
		return true
	}

	// initialize the current index to 0
	currentIndex := 0
	for {
		// Get the step (value of the current) index
		steps := a[currentIndex]

		// set the next index to be the current index + step
		nextIndex := currentIndex + int(steps)

		// if the next index is the last element, return true
		if nextIndex == n-1 {
			return true
		}

		// index out of range, return false
		if nextIndex >= n {
			return false
		}

		// if step is 0 and we haven't reached the last index, can't move again, returen false
        if steps == 0 && nextIndex < n-1 {
            return false
        }

		// reinitialize the current index
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