package bootcamp

import "github.com/alem-platform/ap"

func PrintSquare(w int) {
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			ap.PutRune('0')
			ap.PutRune(' ')
		}
		ap.PutRune('\n')
	}
}

// func main() {
// 	PrintSquare(3)
// }
