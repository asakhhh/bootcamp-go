package bootcamp

func copyof(a *[][]int, n int) [][]int {
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
		for j := 0; j < n; j++ {
			b[i][j] = (*a)[i][j]
		}
	}
	return b
}

func KnightTourCount(res *[][][]int, board *[][]int, n, y, x, cnt int) {
	if cnt == n*n {
		*res = append(*res, copyof(board, n))
		return
	}
	dy := []int{2, 2, 1, 1, -1, -1, -2, -2}
	dx := []int{1, -1, 2, -2, 2, -2, 1, -1}

	for dir := 0; dir < 8; dir++ {
		y += dy[dir]
		x += dx[dir]
		if y >= 0 && y < n && x >= 0 && x < n && (*board)[y][x] == 0 {
			cnt++
			(*board)[y][x] = cnt
			KnightTourCount(res, board, n, y, x, cnt)
			cnt--
			(*board)[y][x] = 0
		}
		y -= dy[dir]
		x -= dx[dir]
	}
}

func KnightTour(n int) [][][]int {
	res := make([][][]int, 0)
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			board[y][x] = 1
			KnightTourCount(&res, &board, n, y, x, 1)
			board[y][x] = 0
		}
	}
	return res
}

// func PrintMatrix(a [][]int) {
// 	n := len(a)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			fmt.Printf("%d\t", a[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	solutions5 := KnightTour(5)
// 	PrintMatrix(solutions5[0])
// 	// Output:
// 	// 1   14  9   20  3
// 	// 24  19  2   15  10
// 	// 13  8   25  4   21
// 	// 18  23  6   11  16
// 	// 7   12  17  22  5

// 	fmt.Println()

// 	PrintMatrix(solutions5[1])
// 	// Output:
// 	// 1   14  9   20  3
// 	// 24  19  2   15  10
// 	// 13  8   23  4   21
// 	// 18  25  6   11  16
// 	// 7   12  17  22  5

// 	solutions2 := KnightTour(2)
// 	fmt.Println("KnightTour(2):", len(solutions2)) // KnightTour(2): 0

// 	solutions1 := KnightTour(1)
// 	fmt.Println("KnightTour(1):", len(solutions1)) // KnightTour(1): 1
// }
