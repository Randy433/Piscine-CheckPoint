package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	finalList := []int{}

	len1 := len(slice1)
	len2 := len(slice2)

	// Determine which slice is larger
	startWithFirst := false
	if len1 > len2 {
		startWithFirst = true
	}

	// If equal, also start with the first slice (as per instructions)
	if len1 == len2 {
		startWithFirst = true
	}

	// Find the max length between both slices
	maxLen := len1
	if len2 > maxLen {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if startWithFirst {
			// Append from first slice if index is valid
			if i < len1 {
				finalList = append(finalList, slice1[i])
			}
			// Append from second slice if index is valid
			if i < len2 {
				finalList = append(finalList, slice2[i])
			}
		} else {
			// Append from second slice first (since it's larger)
			if i < len2 {
				finalList = append(finalList, slice2[i])
			}
			// Append from first slice next (if available)
			if i < len1 {
				finalList = append(finalList, slice1[i])
			}
		}
	}

	return finalList
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
