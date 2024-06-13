package bootcamp

func PowRecursion(x, power int) int {
	if power < 0 {
		return -1
	}
	if power == 0 {
		return 1
	}
	if power%2 == 1 {
		return x * PowRecursion(x, power-1)
	}

	xp := PowRecursion(x, power/2)
	return xp * xp
}
