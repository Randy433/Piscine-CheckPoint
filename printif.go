package main

func FindIf(str string, ch rune) bool {
	for _, c := range str {
		if c == ch {
			return true
		}
	}
	return false
}
