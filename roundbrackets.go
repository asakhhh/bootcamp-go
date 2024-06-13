package bootcamp

func RoundBrackets(s string) bool {
	c := 0

	for _, v := range s {
		if v == '(' {
			c++
		} else if v == ')' {
			c--
			if c < 0 {
				return false
			}
		}
	}
	return c == 0 && len(s) != 0
}
