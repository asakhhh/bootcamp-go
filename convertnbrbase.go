package bootcamp

func Contains(str string, substr string) bool {
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
		if Contains(base[:i], string(base[i])) {
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

// func main() {
// 	result := ConvertNbrBase(1465, "0123456789")
// 	fmt.Println(result) // 1465

// 	result = ConvertNbrBase(1465, "01")
// 	fmt.Println(result) // 10110111001

// 	result = ConvertNbrBase(1465, "01234567")
// 	fmt.Println(result) // 2671

// 	result = ConvertNbrBase(1465, "0123456789ABCDEF")
// 	fmt.Println(result) // 5B9

// 	result = ConvertNbrBase(1465, "00")
// 	fmt.Println(result) //
// }
