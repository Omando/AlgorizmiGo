package recursion

func GCD(n, m int) int {
	// Exit condition
	if m == 0 {
		return n
	}

	if n > m {
		return GCD(m, n%m)
	} else if n < m {
		return GCD(n, m%n)
	}

	// Both n and m are equal so return either
	return n
}
