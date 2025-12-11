package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]
	cells := [2048]byte{}
	ptr := 0

	// Precompute matching brackets for fast jumps
	bracketMap := make(map[int]int)
	stack := []int{}

	for i, c := range code {
		if c == '[' {
			stack = append(stack, i)
		} else if c == ']' {
			if len(stack) == 0 {
				// invalid code, but problem says code is always valid
				return
			}
			openIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bracketMap[openIdx] = i
			bracketMap[i] = openIdx
		}
	}

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
			if ptr >= len(cells) {
				ptr = 0 // wrap around (optional)
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = len(cells) - 1 // wrap around (optional)
			}
		case '+':
			cells[ptr]++
		case '-':
			cells[ptr]--
		case '.':
			fmt.Print(string(cells[ptr]))
		case '[':
			if cells[ptr] == 0 {
				i = bracketMap[i] // jump forward
			}
		case ']':
			if cells[ptr] != 0 {
				i = bracketMap[i] // jump back
			}
		// any other character is ignored
		}
	}
}
