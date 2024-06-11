package bootcamp

func CountRunes(s string) int {
	cnt := 0
	for _, v := range s {
		if v == v {
			cnt++
		}
	}
	return cnt
}

// func main() {
// 	fmt.Println(CountRunes("salem❗"))
// }
