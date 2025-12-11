package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	arr, ok := parseArray(arrStr)
	if !ok {
		return
	}

	target, ok := parseInt(targetStr)
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	findPairs(arr, target)
}

func findPairs(arr []int, target int) {
	found := false

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Printf("Pair found: [%d, %d]\n", i, j)
				found = true
			}
		}
	}

	if !found {
		fmt.Println("No pairs found.")
	}
}

func parseArray(s string) ([]int, bool) {
	if len(s) < 2 || s[0] != '[' || s[len(s)-1] != ']' {
		fmt.Println("Invalid input.")
		return nil, false
	}

	// Remove brackets
	content := s[1 : len(s)-1]
	if len(content) == 0 {
		return []int{}, true
	}

	var nums []int
	current := ""
	for i := 0; i < len(content); i++ {
		c := content[i]

		if c == '-' || (c >= '0' && c <= '9') {
			current += string(c)
		} else if c == ',' || c == ' ' {
			if current != "" {
				val, ok := parseInt(current)
				if !ok {
					fmt.Println("Invalid number:", current)
					return nil, false
				}
				nums = append(nums, val)
				current = ""
			}
		} else {
			fmt.Println("Invalid input.")
			return nil, false
		}
	}

	// Last number
	if current != "" {
		val, ok := parseInt(current)
		if !ok {
			fmt.Println("Invalid number:", current)
			return nil, false
		}
		nums = append(nums, val)
	}

	return nums, true
}

func parseInt(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	}

	num := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int(s[i]-'0')
	}

	return num * sign, true
}
