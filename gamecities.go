package bootcamp

func ToLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func LastChar(s string) byte {
	return ToLower(s[len(s)-1])
}

func Copyof(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	return res
}

func Loop(cities []string, used []bool, cur []int, cnt int) (int, []int) {
	mx := cnt
	mxchain := Copyof(cur)

	for i := range cities {
		if !used[i] && (cnt == 0 || LastChar(cities[cur[cnt-1]]) == ToLower(cities[i][0])) {
			used[i] = true
			cur = append(cur, i)
			newcnt, news := Loop(cities, used, cur, cnt+1)
			used[i] = false
			cur = cur[:cnt]
			if mx < newcnt {
				mx = newcnt
				mxchain = Copyof(news)
			}
		}
	}

	return mx, mxchain
}

func GameCities(cities []string) []string {
	used := make([]bool, len(cities))
	for i := range used {
		used[i] = false
	}
	var cur []int

	_, gt := Loop(cities, used, cur, 0)
	var ans []string
	for _, v := range gt {
		ans = append(ans, cities[v])
	}

	return ans
}

// func main() {
// 	cities := []string{
// 		"Astana",
// 		"Tokyo",
// 		"Amman",
// 		"Monaco",
// 		"Seoul",
// 		"Nassau",
// 		"Oslo",
// 		"Rabat",
// 		"Dover",
// 		"Ottawa",
// 		"Accra",
// 		"Ulaanbaatar",
// 	}
// 	chain := GameCities(cities)
// 	fmt.Println(chain) // [Monaco Oslo Ottawa Astana Accra Amman Nassau Ulaanbaatar Rabat Tokyo]
// }
