package recursion

import "fmt"

func StackUnwind(n int) {
	// Exit condition
	if n < 0 {
		return
	}

	fmt.Printf("Before recursion - %v\n", n)
	StackUnwind(n - 1)
	fmt.Printf("After recursion - %v\n", n)
}

func RecursionWithDataCaches(n int) {
	// TODO
}
