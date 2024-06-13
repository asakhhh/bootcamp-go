package bootcamp

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			return false
		}
	}
	return true
}
