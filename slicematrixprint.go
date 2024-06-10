package bootcamp

import "github.com/alem-platform/ap"

func SliceMatrixPrint(matrix [][]rune) {
	for _, v := range matrix {
		for i := 0; i < len(v); i++ {
			ap.PutRune(v[i])
			if i != len(v)-1 {
				ap.PutRune(' ')
			}
		}
		ap.PutRune('\n')
	}
}

// func main() {
// 	matrix := [][]rune{
// 		{'a', 'b', 'c'},
// 		{'d', 'e', 'f'},
// 		{'g', 'h', 'i'},
// 	}
// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// a b c
// 	// d e f
// 	// g h i
// }
