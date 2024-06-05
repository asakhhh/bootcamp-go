package bootcamp

import "github.com/alem-platform/ap"

func PrintAlphabetLn() {
	for i := 0; i < 26; i++ {
		ap.PutRune(rune('a' + i))
	}
	ap.PutRune('\n')
}

// func main() {
// 	PrintAlphabetLn()
// }
