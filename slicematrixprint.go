package bootcamp

import "github.com/alem-platform/ap"

func SliceMatrixPrint(matrix [][]rune) {
	for i, v := range matrix {
		for i := 0; i < len(v); i++ {
			ap.PutRune(v[i])
			if i != len(v)-1 {
				ap.PutRune(' ')
			}
		}
		if i != len(matrix)-1 {
			ap.PutRune('\n')
		}
	}
}
