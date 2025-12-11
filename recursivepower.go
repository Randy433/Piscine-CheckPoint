package main

func recursivepower(nb int, power int) int {
	if power == 0 {
		return 1
	}

	if power > 0 {
		return nb * recursivepower(nb, power-1)
	}
	return 0
}
