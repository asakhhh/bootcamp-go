package bootcamp

func Contains(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Permute(pref []int, res [][]int) {
	n := len(pref)

	if n == 8 {
		for i := 0; i < 7; i++ {
			for d := 1; i+d < 8; d++ {
				if abs(pref[i+d]-pref[i]) == d {
					return
				}
			}
		}
		res = append(res, pref)
		return
	}

	for i := 1; i <= 8; i++ {
		if !Contains(pref, i) {
			pref = append(pref, i)
			Permute(pref, res)
			pref = pref[:n-1]
		}
	}
}

func EightQueens() [][]int {
	var res [][]int

	Permute([]int{}, res)

	return res
}
