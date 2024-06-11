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

// func main() {
// 	matrix := [][]rune{
// 		{'a', 'b', 'c'},
// 		{'d', 'e', 'f'},
// 		{'g', 'h', 'i'},
// 	}

// 	fmt.Println(matrix)

// 	SliceMatrixRotateClockwise(matrix)

// 	fmt.Println(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
// }
