package bootcamp

var used [][]bool

func IslandCost(matrix [][]int, x, y int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 || x < 0 || x >= len(matrix[0]) || y < 0 || y > len(matrix) || matrix[y][x] == 0 || (len(used) != 0 && used[y][x]) {
		return 0
	}
	if len(used) == 0 {
		for _, v := range matrix {
			used = append(used, make([]bool, len(v)))
		}
	}
	used[y][x] = true
	return matrix[y][x] + IslandCost(matrix, x, y-1) + IslandCost(matrix, x, y+1) + IslandCost(matrix, x-1, y) + IslandCost(matrix, x+1, y)
}

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

// 	fmt.Println(IslandCost(matrix, 4, 2)) // 11
// 	fmt.Println(IslandCost(matrix, 0, 0)) // 5
// 	fmt.Println(IslandCost(matrix, 1, 4)) // 1
// 	fmt.Println(IslandCost(matrix, 0, 3)) // 0
// }
