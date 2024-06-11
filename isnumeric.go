package bootcamp

func IsNumeric(s string) bool {
	if len(s) == 0 || !(s[0] == '+' || s[0] == '-' || (s[0] >= '0' && s[0] <= '9')) {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(IsNumeric("123"))     // true
// 	fmt.Println(IsNumeric("-123"))    // true
// 	fmt.Println(IsNumeric("+123"))    // true
// 	fmt.Println(IsNumeric("+-123"))   // false
// 	fmt.Println(IsNumeric(" 123"))    // false
// 	fmt.Println(IsNumeric("123 abc")) // false
// 	fmt.Println(IsNumeric("123abc"))  // false
// 	fmt.Println(IsNumeric("ab"))      // false
// }
