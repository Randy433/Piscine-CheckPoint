package main

func sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	for i := 1; ; i++ {
		square := i * i
		if square == nb {
			return i
		}
		if square > nb {
			return 0
		}
	}
}
