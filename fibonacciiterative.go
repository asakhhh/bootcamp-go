package bootcamp

func FibonacciIterative(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}

	prev, cur := 0, 1
	for i := 1; i < n; i++ {
		prev, cur = cur, prev+cur
	}

	return cur
}
