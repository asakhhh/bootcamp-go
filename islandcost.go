package bootcamp

func IslandCost1(matrix [][]int, x, y int) int {
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

	res := matrix[y][x]
	matrix[y][x] = 0
	if y > 0 {
		res += IslandCost1(matrix, x, y-1)
	}
	if y+1 < len(matrix) {
		res += IslandCost1(matrix, x, y+1)
	}
	if x > 0 {
		res += IslandCost1(matrix, x-1, y)
	}
	if x+1 < len(matrix[y]) {
		res += IslandCost1(matrix, x+1, y)
	}
	return res
}

func IslandCost(matrix [][]int, x, y int) int {
	matrix1 := make([][]int, len(matrix))
	for i := range matrix {
		matrix1[i] = make([]int, len(matrix[i]))
		copy(matrix1[i], matrix[i])
	}
	return IslandCost1(matrix1, x, y)
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
