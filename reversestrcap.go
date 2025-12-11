package main

import (
	"fmt"
	"os"
	// "unicode"
)

func ToLower(s rune) rune  {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return s
}

func ToUpper(s rune) rune  {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		runes := []rune(arg)
		for i := 0; i < len(runes); i++ {
			// if the current element is a space
			if runes[i] == ' '{
				continue
			}

			if i+1 == len(runes) || runes[i+1] == ' ' {
				runes[i] = ToUpper(runes[i])
			} else {
				runes[i] = ToLower(runes[i])
			}
		}
		// for i := 0; i < len(runes); i++ {
		// 	if i+1 < len(runes) && runes[i+1] != ' ' && runes[i] != ' ' {
		// 		runes[i] = unicode.ToLower(runes[i])
		// 	} else if runes[i] != ' ' {
		// 		runes[i] = unicode.ToUpper(runes[i])
		// 	}
		// }

		fmt.Println(string(runes))
	}
}
