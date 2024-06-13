package bootcamp

func Contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func UniqCombN(s string, n int) []string {
	if len(s) < n {
		return make([]string, 0)
	}
	for i, v := range s {
		if Contains(s[:i], v) {
			return make([]string, 0)
		}
	}
	if n == 0 {
		return append(make([]string, 0), "")
	}

	var res []string
	for i, v := range s {
		prev := UniqCombN(s[:i]+s[i+1:], n-1)
		for _, v1 := range prev {
			res = append(res, string(v)+v1)
		}
	}

	return res
}

// func main() {
// 	fmt.Println(UniqCombN("abc", 1)) // ["a", "b", "c"]
// 	fmt.Println(UniqCombN("abc", 2)) // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqCombN("ab", 2))  // ["ab", "ba"]
// 	fmt.Println(UniqCombN("a", 1))   // ["a"]
// 	fmt.Println(UniqCombN("ab", 3))  // []
// 	fmt.Println(UniqCombN("aa", 1))  // []
// }
