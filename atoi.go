package bootcamp

func Atoi(s string) int {
	minus := false
	first := 0

	if s[0] == '-' {
		minus = true
	}
	if s[0] == '-' || s[0] == '+' {
		first++
	}

	var n int
	for i := first; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(rune(s[i])-'0')
	}

	if minus {
		n = -n
	}
	return n
}

// func main() {
// 	fmt.Println(Atoi("123"))
// 	fmt.Println(Atoi("+123"))
// 	fmt.Println(Atoi("-123"))
// 	fmt.Println(Atoi("-123!"))
// 	fmt.Println(Atoi("abc"))
// }
