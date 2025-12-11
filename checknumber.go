package main

func CheckNumber(st string) bool {
	for _, ch := range st {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}
