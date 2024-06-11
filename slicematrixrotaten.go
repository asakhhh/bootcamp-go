package bootcamp

func SliceMatrixRotateClockwise(matrix [][]rune) {
	n := len(matrix)

	if n == 0 || n != len(matrix[0]) {
		return
	}

	a := make([][]rune, n)
	for i := 0; i < n; i++ {
		a[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			a[i][j] = matrix[n-1-j][i]
		}
	}

	for i := 0; i < n; i++ {
		copy(matrix[i], a[i])
	}
}

func SliceMatrixRotateN(matrix [][]rune, turns int) {
	turns = (turns%4 + 4) % 4

	for i := 0; i < turns; i++ {
		SliceMatrixRotateClockwise(matrix)
	}

	// n := len(matrix)
	// a := make([][]rune, n)
	// for i := 0; i < n; i++ {
	// 	a[i] = make([]rune, n)
	// 	for j := 0; j < n; j++ {
	// 		a[i][j] = matrix[n-1-j][i]
	// 	}
	// }

	// for i := 0; i < n; i++ {
	// 	copy(matrix[i], a[i])
	// }

	// for i := 0; i < n; i++ {
	// 	a[i] = make([]rune, n)
	// 	for j := 0; j < n; j++ {
	// 		a[i][j] = matrix[n-1-j][i]
	// 	}
	// }

	// for i := 0; i < n; i++ {
	// 	copy(matrix[i], a[i])
	// }

	// for i := 0; i < n; i++ {
	// 	a[i] = make([]rune, n)
	// 	for j := 0; j < n; j++ {
	// 		a[i][j] = matrix[n-1-j][i]
	// 	}
	// }

	// for i := 0; i < n; i++ {
	// 	copy(matrix[i], a[i])
	// }
}

// func main() {
// 	matrix := [][]rune{
// 		{'a', 'b', 'c'},
// 		{'d', 'e', 'f'},
// 		{'g', 'h', 'i'},
// 	}

// 	fmt.Println(matrix)

// 	SliceMatrixRotateCounterClockwise(matrix)

// 	fmt.Println(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
// }
