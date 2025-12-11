package main

func Slice(a []string, nbrs... int) []string{
	n := len(a)

	if len(nbrs) == 0 {
		// no indexes given → return empty
		return []string{}
	}

	// Convert start index
	start := nbrs[0]
	if start < 0 {
		start = n + start
	}
	if start < 0 {
		start = 0
	}
	if start > n {
		start = n
	}

	// If only one index
	if len(nbrs) == 1 {
		return a[start:]
	}

	// Convert end index
	end := nbrs[1]
	if end < 0 {
		end = n + end
	}
	if end < 0 {
		end = 0
	}
	if end > n {
		end = n
	}

	// Ensure correct slicing order
	if start > end {
		return []string{}
	}

	return a[start:end]
}
