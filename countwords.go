package bootcamp

func ToLower(s string) string {
	var t string
	ch := 'A' - 'a'
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			t = t + string(v-ch)
		} else {
			t = t + string(v)
		}
	}

	return t
}

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
		if (c >= 'a' && c <= 'z') || c == ' ' {
			t += string(c)
		}
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

func CountSubstr(s []string, substr string) int {
	cnt := 0
	for _, v := range s {
		if substr == v {
			cnt++
		}
	}
	return cnt
}

func CountWords(s string) map[string]int {
	ss := ToLower(s)
	t := Split(ss, " ")
	res := make(map[string]int)

	for _, v := range t {
		res[v] = CountSubstr(t, v)
	}

	return res
}

// func main() {
// 	s := "The soup was stirred and stirred until thickened."
// 	wordCounts := CountWords(s)
// 	fmt.Println(wordCounts) // map[the:1 soup:1 was:1 and:1 stirred:2 until:1 thickened:1]
// }
