package bootcamp

func RotN(s string, n int) string {
	var t string

	n = (n%26 + 26) % 26

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			t += string('a' + (rune(c)-'a'+rune(n))%26)
		} else if c >= 'A' && c <= 'Z' {
			t += string('A' + (rune(c)-'A'+rune(n))%26)
		} else {
			t += string(c)
		}
	}

	return t
}
