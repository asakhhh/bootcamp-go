package bootcamp

import "math/rand"

func TestNthRune(fn func(s string, n int) rune) bool {
	var s string
	n := int(rand.Int31n(1000) + 1)

	for i := 0; i < n; i++ {
		s += string(byte(rand.Int31n(256)))
	}

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if fn(s, i) != r[i] {
			return false
		}
	}
	if fn(s, -1) != -1 || fn(s, len(s)+1) != -1 {
		return false
	}
	return true
}
