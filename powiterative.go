package bootcamp

func PowIterative(x, power int) int {
	if power < 0 {
		return -1
	}
	res := 1

	for ; power > 0; power-- {
		res *= x
	}

	return res
}
