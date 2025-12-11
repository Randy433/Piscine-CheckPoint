package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hexOutput := ""
	asciiOutput := ""

	const bytesPerLine = 5

	for i, b := range arr {
		if i > 0 && i%bytesPerLine == 0 {
			fmt.Printf("%-15s %s\n", hexOutput, asciiOutput)
			hexOutput = ""
			asciiOutput = ""
		}

		hexOutput += fmt.Sprintf("%02x ", b)

		if b >= 32 && b <= 126 {
			asciiOutput += string(b)
		} else {
			asciiOutput += "."
		}
	}

	fmt.Printf("%-15s %s\n", hexOutput, asciiOutput)
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
