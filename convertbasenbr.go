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

func ConvertBaseNbr(s string, base string) int {
	if len(base) < 2 {
		return -1
	}
	for i := 1; i < len(base); i++ {
		if contains(base[:i], string(base[i])) {
			return -1
		}
	}

	m := make(map[rune]int)
	for i, c := range base {
		m[c] = i
	}

	num := 0
	for _, c := range s {
		if !contains(base, string(c)) {
			return -1
		}
		num = num*len(base) + m[c]
	}

	return num
}

// func main() {
// 	fmt.Println(ConvertBaseNbr("1465", "0123456789"))
// 	fmt.Println(ConvertBaseNbr("10110111001", "01"))
// 	fmt.Println(ConvertBaseNbr("2671", "01234567"))
// 	fmt.Println(ConvertBaseNbr("5B9", "0123456789ABCDEF"))
// 	fmt.Println(ConvertBaseNbr("1", "00"))
// }
