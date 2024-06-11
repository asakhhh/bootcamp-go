package bootcamp

func Rot13(s string) string {
	var t string

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			t += string('a' + (rune(c)-'a'+13)%26)
		} else if c >= 'A' && c <= 'Z' {
			t += string('A' + (rune(c)-'A'+13)%26)
		} else {
			t += string(c)
		}
	}

	return t
}
