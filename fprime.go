package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsPrime2(nbr int) bool {
	if nbr < 2 {
		return false
	}
	for i := 2; i*i <= nbr; i++ {
		if nbr%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}

	val, err := strconv.Atoi(args[1])
	if err != nil || val <= 1 {
		return
	}

	res := ""

	for i := 2; i <= val; i++ {
		if IsPrime2(i) {
			for val%i == 0 {
				res += strconv.Itoa(i)
				val = val / i
				if val > 1 {
					res += "*"
				}
			}
		}
	}

	fmt.Println(res)
}
