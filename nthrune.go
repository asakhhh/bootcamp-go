package bootcamp

func NthRune(s string, n int) rune {
	if n >= len(s) {
		return 0
	}
	return rune(s[n])
}

// func main() {
// 	ap.PutRune(NthRune("salem", 0))
// 	ap.PutRune(NthRune("student", 2))
// 	ap.PutRune('\n')
// }
