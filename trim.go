package bootcamp

func Trim(s string) string {
	if len(s) == 0 {
		return ""
	}

	st, fn, n := 0, len(s), len(s)

	for ; st < n && s[st] == ' '; st++ {
	}
	for ; fn > 0 && s[fn-1] == ' '; fn-- {
	}

	var t string
	for i := st; i < fn; i++ {
		t += string(s[i])
	}

	return t
}

// func main() {
// 	fmt.Println(Trim("   Salem student!   "))
// 	fmt.Println(Trim("   Trim spaces   "))
// }
