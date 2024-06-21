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

// func main() {
// 	validSudoku := [9][9]int{
// 		{5, 3, 4, 6, 7, 8, 9, 1, 2},
// 		{6, 7, 2, 1, 9, 5, 3, 4, 8},
// 		{1, 9, 8, 3, 4, 2, 5, 6, 7},
// 		{8, 5, 9, 7, 6, 1, 4, 2, 3},
// 		{4, 2, 6, 8, 5, 3, 7, 9, 1},
// 		{7, 1, 3, 9, 2, 4, 8, 5, 6},
// 		{9, 6, 1, 5, 3, 7, 2, 8, 4},
// 		{2, 8, 7, 4, 1, 9, 6, 3, 5},
// 		{3, 4, 5, 2, 8, 6, 1, 7, 9},
// 	}

// 	invalidSudoku := [9][9]int{
// 		{5, 3, 4, 6, 7, 8, 9, 1, 2},
// 		{6, 7, 2, 1, 9, 5, 3, 4, 8},
// 		{1, 9, 8, 3, 4, 2, 5, 6, 7},
// 		{8, 5, 9, 7, 6, 1, 4, 2, 3},
// 		{4, 2, 6, 8, 5, 3, 7, 9, 1},
// 		{7, 1, 3, 9, 2, 4, 8, 5, 6},
// 		{9, 6, 1, 5, 3, 7, 2, 8, 4},
// 		{2, 8, 7, 4, 1, 9, 6, 3, 5},
// 		{3, 4, 5, 2, 8, 6, 1, 7, 8}, // invalid row
// 	}

// 	fmt.Println(ValidSudoku(validSudoku))   // true
// 	fmt.Println(ValidSudoku(invalidSudoku)) // false
// }
