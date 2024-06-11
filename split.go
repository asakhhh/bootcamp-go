package bootcamp

func Split(s string, sep string) []string {
	on_sep := 0
	var t string
	var res []string

	if sep == "" {
		for _, c := range s {
			res = append(res, string(c))
		}
		return res
	}

	for _, c := range s {
		t += string(c)
		if rune(c) == rune(sep[on_sep]) {
			on_sep++
		}
		if on_sep == len(sep) {
			res = append(res, t[:len(t)-len(sep)])
			t = ""
			on_sep = 0
		}
	}
	if len(t) > 0 {
		res = append(res, t)
	}
	if len(res) == 0 {
		res = append(res, "")
	}

	return res
}

// func main() {
// 	fmt.Println(Split("a,b,c", ","))
// 	fmt.Println(Split("salem-student", "-"))
// 	fmt.Println(Split("salem", ""))
// }
