package bootcamp

func isOpen(b rune) bool {
	return b == '(' || b == '{' || b == '['
}

func isClose(b rune) bool {
	return b == ')' || b == '}' || b == ']'
}

func isMatch(a, b rune) bool {
	if a == '{' && b == '}' {
		return true
	}
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	return false
}

func Brackets(s string) bool {
	a := make([]rune, 0)

	for _, c := range s {
		if isOpen(c) {
			a = append(a, c)
		} else if isClose(c) {
			if len(a) == 0 || !isMatch(a[len(a)-1], c) {
				return false
			} else {
				a = a[:len(a)-1]
			}
		}
	}
	return len(a) == 0
}
