package bootcamp

var used [][]bool

func IslandCost(matrix [][]int, x, y int) int {
	if y < 0 || x < 0 || len(matrix) <= y || len(matrix[y]) <= x || matrix[y][x] == 0 || (len(used) != 0 && used[y][x]) {
		return 0
	}
	if len(used) == 0 {
		for _, v := range matrix {
			t := make([]bool, len(v))
			for i := range t {
				t[i] = false
			}
			used = append(used, t)
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
