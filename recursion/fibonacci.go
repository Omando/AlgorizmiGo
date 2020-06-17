package recursion

/* Calculate Fibonacci value iteratively
n:   0 1 2 3 4 5 6 7
fib: 0 1 1 2 3 5 8 13
*/
func FibonacciIterative(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	n_2 := 0
	n_1 := 1
	var result int
	for i := 2; i <= n; i++ {
		result = n_2 + n_1
		n_2 = n_1
		n_1 = result
	}

	return result
}

func FibonacciRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	_, n1 := fibonacciRecursiveImpl(n) // last result stored in n1. No need to do n2+n1 to get result
	return n1
}

func fibonacciRecursiveImpl(n int) (n2 int, n1 int) {
	if n == 1 {
		return 0, 1
	}

	n--
	n2, n1 = fibonacciRecursiveImpl(n)
	result := n2 + n1
	n2 = n1
	n1 = result
	return n2, n1 // last result stored in n1. No need to do n2+n1 to get result
}
