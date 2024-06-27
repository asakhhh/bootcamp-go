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
			(*board)[y][x] = cnt + 1
			KnightTourCount(res, board, n, y, x, cnt+1)
			(*board)[y][x] = 0
		}
		y -= dy[dir]
		x -= dx[dir]
	}
}

func KnightTour(n int) [][][]int {
	res := make([][][]int, 0)

	if n <= 0 {
		return res
	}

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	for y := 0; y < 1; y++ {
		for x := 0; x < 1; x++ {
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
// 	fmt.Println(len(KnightTour(5)))
// }
