package bootcamp

func TestNthRune(fn func(s string, n int) rune) bool {
	s := "abacabac11"

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if fn(s, i) != r[i] {
			return false
		}
	}
	if fn(s, -1) >= 0 || fn(s, len(s)+1) >= 0 {
		return false
	}
	return true
}
