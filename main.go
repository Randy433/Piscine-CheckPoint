// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var hadError bool = false

// func fatalError() {
// 	os.Exit(1)
// }

// func tailFile(filepath string, N int) {
// 	file, err := os.ReadFile(filepath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
// 		hadError = true
// 		return
// 	}

// 	startIndex := max(len(file) - N, 0)

// 	// fmt.Print(string(file[startIndex:]))
// 	os.Stdout.Write(file[startIndex:])
// }

// func main() {
// 	arguments := os.Args

// 	if len(arguments) < 4 {
// 		fatalError()
// 	}

// 	if arguments[1] != "-c" {
// 		fatalError()
// 	}

// 	numBytes, err := strconv.Atoi(arguments[2])
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error converting to integer: %v\n", err)
// 		fatalError()
// 	}

// 	if numBytes < 0 {
// 		fatalError()
// 	}

// 	files := arguments[3:]
// 	multiFile := len(files) > 1

// 	for i, file := range files {
// 		if multiFile {
// 			if i > 0 {
// 				fmt.Println()
// 			}
// 			fmt.Printf("==> %s <==\n", file)
// 		}

// 		tailFile(file, numBytes)
// 	}

// 	if hadError {
// 		os.Exit(1)
// 	}
// }

package main

import "fmt"

func ActiveBits(n int) int {
	binaryStr := fmt.Sprintf("%b", n)

	// fmt.Print(binaryStr)
	slice := []byte(binaryStr)
	counter := 0
	for _, bit := range slice {
		if bit == '1' {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(ActiveBits(-5))
	fmt.Println(FindIf("toyy7", 'y'))
	fmt.Println(FindIf("toyy", 'i'))
}
