package main

import "fmt"

func SliceMatrixRotateCounterClockwise(matrix [][]rune) {
	n := len(matrix)
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

	for i := 0; i < n; i++ {
		a[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			a[i][j] = matrix[n-1-j][i]
		}
	}

	for i := 0; i < n; i++ {
		copy(matrix[i], a[i])
	}

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

func main() {
	matrix := [][]rune{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}

	fmt.Println(matrix)

	SliceMatrixRotateCounterClockwise(matrix)

	fmt.Println(matrix)
	// Output:
	// g d a
	// h e b
	// i f c
}
