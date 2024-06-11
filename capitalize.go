package bootcamp

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	new_word := true
	var t string
	for _, v := range s {
		if new_word {
			if v >= 'a' && v <= 'z' {
				t += string(rune(v) + 'A' - 'a')
			} else {
				t += string(v)
			}
			new_word = false
		} else {
			t += string(v)
			if v == ' ' {
				new_word = true
			}
		}
	}

	return t
}

// func main() {
// 	fmt.Println(Capitalize("salem student"))
// 	fmt.Println(Capitalize("SALEM-5TUDENT, Q41aisyn?"))
// }
