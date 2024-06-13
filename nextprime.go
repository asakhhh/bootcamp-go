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

func NextPrime(n int) int {
	if n <= 0 {
		return NextPrime(1)
	}

	if !IsPrime(n + 1) {
		return NextPrime(n + 1)
	}
	return n + 1
}
