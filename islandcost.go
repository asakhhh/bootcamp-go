package bootcamp

var used [][]bool

func IslandCost(matrix [][]int, x, y int) int {
	if y < 0 || x < 0 {
		return 0
	}
	if len(matrix) <= y {
		return 0
	}
	if len(matrix[y]) <= x {
		return 0
	}
	if matrix[y][x] == 0 {
		return 0
	}
	if len(used) <= y {
		used = used[:0]
		for _, v := range matrix {
			t := make([]bool, len(v))
			for i := range t {
				t[i] = false
			}
			used = append(used, t)
		}
	}
	used[y][x] = true
	res := matrix[y][x]
	if y > 0 && !used[y-1][x] && matrix[y-1][x] > 0 {
		res += IslandCost(matrix, x, y-1)
	}
	if y+1 < len(matrix) && !used[y+1][x] && matrix[y-1][x] > 0 {
		res += IslandCost(matrix, x, y+1)
	}
	if x > 0 && !used[y][x-1] && matrix[y-1][x] > 0 {
		res += IslandCost(matrix, x-1, y)
	}
	if x+1 < len(matrix[y]) && !used[y][x+1] && matrix[y-1][x] > 0 {
		res += IslandCost(matrix, x+1, y)
	}
	return res
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
