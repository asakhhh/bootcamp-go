package bootcamp

func Sqrt(x int) int {
	if x < 0 {
		return -1
	}

	p := 0
	for ; p*p < x; p++ {
	}

	if p*p != x {
		return 0
	}
	return p
}
