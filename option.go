package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	flags := 0
	for _, arg := range os.Args[1:] {
		if arg == "-h" {
			printHelp()
			return
		}

		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		for i := 1; i < len(arg); i++ {
			c := arg[i]
			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			flags |= 1 << (c - 'a')
		}
	}

	printBits(flags)
}

// Print the 32-bit representation, spaced every 8 bits
func printBits(flags int) {
	for i := 31; i >= 0; i-- {
		if flags&(1<<i) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func printHelp() {
	// all bits zero
	fmt.Println("00000000 00000000 00000000 00000000")
}
