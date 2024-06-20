package bootcamp

func contains(str string, substr string) bool {
	found := 0
	for _, v := range str {
		if rune(v) == rune(substr[found]) {
			found++
		} else {
			found = 0
		}
		if found == len(substr) {
			return true
		}
	}
	return false
}

func ConvertNbrBase(n int, base string) string {
	if len(base) < 2 || n < 0 {
		return ""
	}
	for i := 1; i < len(base); i++ {
		if contains(base[:i], string(base[i])) {
			return ""
		}
	}

	b := len(base)
	var ans string

	for n > 0 {
		ans = string(base[n%b]) + ans
		n /= b
	}

	return ans
}
