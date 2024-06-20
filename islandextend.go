package bootcamp

func Add(matrix [][]int, used [][]bool, x, y int) int {
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) || used[y][x] {
		return 0
	}
	if matrix[y][x] == 0 {
		matrix[y][x] = 1
		used[y][x] = true
		return 1
	}
	used[y][x] = true

	return Add(matrix, used, x-1, y) + Add(matrix, used, x+1, y) + Add(matrix, used, x, y-1) + Add(matrix, used, x, y+1)
}

func IslandExtend(matrix [][]int, x, y int) bool {
	if y < 0 || y >= len(matrix) || x < 0 || x >= len(matrix[y]) {
		return false
	}

	if matrix[y][x] == 0 {
		return false
	}

	used := make([][]bool, len(matrix))
	for i := range used {
		used[i] = make([]bool, len(matrix[i]))
		for j := range used[i] {
			used[i][j] = false
		}
	}
	return Add(matrix, used, x, y) > 0
}

// func PrintMatrix(matrix [][]int) {
// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			fmt.Print(matrix[i][j])
// 			ap.PutRune(' ')
// 		}
// 		ap.PutRune('\n')
// 	}
// }

// func main() {
// 	matrix := [][]int{
// 		{1, 1, 1, 0, 0, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 1, 0, 0, 0, 0},
// 		{0, 0, 0, 1, 2, 3, 1, 0, 0},
// 		{0, 0, 0, 1, 1, 1, 0, 0, 1},
// 		{0, 1, 0, 0, 0, 0, 0, 1, 2},
// 		{0, 0, 0, 1, 0, 0, 0, 0, 1},
// 		{0, 0, 1, 1, 2, 2, 0, 0, 0},
// 	}

// 	fmt.Println(IslandExtend(matrix, 4, 2)) // true
// 	PrintMatrix(matrix)
// 	// 1 1 1 0 1 0 0 0 0
// 	// 1 1 0 1 1 1 1 0 0
// 	// 0 0 1 1 2 3 1 1 0
// 	// 0 0 1 1 1 1 1 0 1
// 	// 0 1 0 1 1 1 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 0, 0)) // true
// 	PrintMatrix(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 0 1 1 1 1 1 0 1
// 	// 0 1 0 1 1 1 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 1, 4)) // true
// 	PrintMatrix(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 1 1 1 1 1 1 0 1
// 	// 1 1 1 1 1 1 0 1 2
// 	// 0 1 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	fmt.Println(IslandExtend(matrix, 0, 3)) // false
// 	PrintMatrix(matrix)
// 	// 1 1 1 1 1 0 0 0 0
// 	// 1 1 1 1 1 1 1 0 0
// 	// 1 1 1 1 2 3 1 1 0
// 	// 0 1 1 1 1 1 1 0 1
// 	// 1 1 1 1 1 1 0 1 2
// 	// 0 1 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0
// }
