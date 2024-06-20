package bootcamp

func valid(a []int) bool {
	var used [9]bool
	for _, v := range a {
		if v < 1 || v > 10 {
			return false
		}
		used[v-1] = true
	}
	for i := 0; i < 9; i++ {
		if !used[i] {
			return false
		}
	}
	return true
}

func ValidSudoku(a [9][9]int) bool {
	var cols [9][]int
	var cells [3][3][]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cols[j] = append(cols[j], a[i][j])
			cells[i/3][j/3] = append(cells[i/3][j/3], a[i][j])
		}
	}

	for i := 0; i < 9; i++ {
		if !valid(cols[i]) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !valid(cells[i][j]) {
				return false
			}
		}
	}
	return true
}
