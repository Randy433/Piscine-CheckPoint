package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	len2 := len(slice2)

	for i := 0; i < len2; i++ {
		if i < len2 {
			slice1 = append(slice1, slice2[i])
		}
	}
	return slice1
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}