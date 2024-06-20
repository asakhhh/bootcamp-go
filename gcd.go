package bootcamp

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	if a > b {
		return GCD(b, a)
	}
	return GCD(b%a, a)
}
