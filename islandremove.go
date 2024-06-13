package bootcamp

func IslandRemove(matrix [][]int, x, y int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[y][x] == 0 {
		return
	}
	matrix[y][x] = 0
	if y > 0 {
		IslandRemove(matrix, x, y-1)
	}
	if y+1 < len(matrix) {
		IslandRemove(matrix, x, y+1)
	}
	if x > 0 {
		IslandRemove(matrix, x-1, y)
	}
	if x+1 < len(matrix[0]) {
		IslandRemove(matrix, x+1, y)
	}
}

// func PrintMatrix(matrix [][]int) {
// 	for _, v := range matrix {
// 		for _, vv := range v {
// 			fmt.Printf("%d ", vv)
// 		}
// 		fmt.Println()
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
// 	PrintMatrix(matrix)
// 	// Output:
// 	// 1 1 1 0 0 0 0 0 0
// 	// 1 1 0 0 1 0 0 0 0
// 	// 0 0 0 1 2 3 1 0 0
// 	// 0 0 0 1 1 1 0 0 1
// 	// 0 1 0 0 0 0 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0

// 	IslandRemove(matrix, 4, 2)
// 	fmt.Println()

// 	PrintMatrix(matrix)
// 	// Output:
// 	// 1 1 1 0 0 0 0 0 0
// 	// 1 1 0 0 0 0 0 0 0
// 	// 0 0 0 0 0 0 0 0 0
// 	// 0 0 0 0 0 0 0 0 1
// 	// 0 1 0 0 0 0 0 1 2
// 	// 0 0 0 1 0 0 0 0 1
// 	// 0 0 1 1 2 2 0 0 0
// }
