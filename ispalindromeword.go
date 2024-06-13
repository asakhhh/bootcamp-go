package bootcamp

func IsPalindromeWord(s string) bool {
	n := len(s)

	if n%2 == 1 && !(s[n/2] >= 'a' && s[n/2] <= 'z') && !(s[n/2] >= 'A' && s[n/2] <= 'Z') {
		return false
	}
	if n == 0 {
		return false
	}

	for i := 0; i < n/2; i++ {
		c1 := s[i]
		c2 := s[n-1-i]

		if !(c1 >= 'a' && c1 <= 'z') && !(c1 >= 'A' && c1 <= 'Z') {
			return false
		} else if !(c2 >= 'a' && c2 <= 'z') && !(c2 >= 'A' && c2 <= 'Z') {
			return false
		}
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 'a' - 'A'
		}
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 'a' - 'A'
		}
		if c1 != c2 {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(IsPalindromeWord("raceCAr")) // true
// 	fmt.Println(IsPalindromeWord("level"))   // true
// 	fmt.Println(IsPalindromeWord("sALem"))   // false
// }
