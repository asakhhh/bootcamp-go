package bootcamp

import "github.com/alem-platform/ap"

func PrintBits(b byte) {
	n := int(b)
	p := 128

	for p > 0 {
		ap.PutRune(rune('0' + (n/p)%2))
		p /= 2
	}
	ap.PutRune('\n')
}

// func main() {
// 	PrintBits(5)
// 	PrintBits(255)
// 	PrintBits(0)
// }
