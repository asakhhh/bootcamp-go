package bootcamp

import "github.com/alem-platform/ap"

func PutNumber(n int) {
	pw := 1

	if n < 0 {
		ap.PutRune('-')
		n = -n
	} else if n == 0 {
		ap.PutRune('0')
		return
	}

	for pw*10 < n {
		pw *= 10
	}

	for pw >= 1 {
		ap.PutRune(rune('0' + n/pw%10))
		pw /= 10
	}
}

// func main() {
// 	PutNumber(-1394083)
// }
