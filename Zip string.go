package main

import "fmt"

func ZipString(s string) string {
	result := ""
	for i:=0; i < len(s); i++{
		count :=1
		for i+1 < len(s) && s[i] == s[i+1]{
			count++
			i++
		}
		result += string(count + '0') + string(s[i])
		//result += string.Itoa(count) + string(s[i])
	}
	return result
}