package main

import (
	"fmt"
	"strconv"
	"os"
)

func IsPrime1(nbr int) bool {
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

func AddPrimeSum(nbr int) int {
	if nbr < 2 {
		return 0
	}

	result := 0

	for i := nbr; i >= 2; i-- {
		if IsPrime1(i) {
			result += i
		}
	}
	return result
}

func main() {
	arg := os.Args
	
	if len(arg) != 2 {
		fmt.Println(0)
		return 
	}

	val, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(0)
		return
	}

	fmt.Println(AddPrimeSum(val))
}