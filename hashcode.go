package main

import (
	"fmt"
	// "honnef.co/go/tools/analysis/code"
	// "strconv"
)

func HashCode(dec string) string {
	if len(dec) == 0 {
		return ""
	}

	hashed := ""

	for _, c := range dec {
		code := (int(c) + len(dec)) % 127
		if code < 33 {
			code += 33
		}
		hashed += string(code)
	}
	return hashed
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
