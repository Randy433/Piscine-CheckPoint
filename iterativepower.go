package main

func iterativepower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power > 0 {
		for i := power; i > 1; i-- {
			nb *= nb
		}
		return nb
	}
	return 0

}
