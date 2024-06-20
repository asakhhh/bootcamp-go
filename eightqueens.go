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

func Permute(pref []int, res *[][]int) {
	n := len(pref)

	if n == 8 {
		for i := 0; i < 7; i++ {
			for d := 1; i+d < 8; d++ {
				if abs(pref[i+d]-pref[i]) == d {
					return
				}
			}
		}
		pref1 := make([]int, len(pref))
		copy(pref1, pref)
		for i := range pref {
			pref1[i]++
		}
		*res = append(*res, pref1)
		return
	}

	for i := 0; i < 8; i++ {
		if !Contains(pref, i) {
			pref = append(pref, i)
			Permute(pref, res)
			pref = pref[:n]
		}
	}
}

func EightQueens() [][]int {
	var res [][]int
	var a []int
	Permute(a, &res)

	return res
}

// func main() {
// 	solutions := EightQueens()
// 	for _, solution := range solutions {
// 		for _, pos := range solution {
// 			fmt.Print(pos)
// 		}
// 		fmt.Println()
// 	}
// }
