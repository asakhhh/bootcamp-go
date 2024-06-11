package bootcamp

func Contains(str string, substr string) bool {
	if substr == "" {
		return true
	}

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

// func main() {
// 	fmt.Println(Contains("hello world", "world"))
// 	fmt.Println(Contains("test", "best"))
// }
