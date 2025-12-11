package main

func recursivefactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb > 0 {
		nb = nb * recursivefactorial(nb-1)
		return nb
	}
	return 0
}
