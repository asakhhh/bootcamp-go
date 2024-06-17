package bootcamp

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
