package bootcamp

func PaddingStart(s string, n int) string {
	m := len([]rune(s))
	res := ""
	for i := 0; i < n-m; i++ {
		res += " "
	}
	return res + s
}

// func main() {
// 	fmt.Printf("%q\n", PaddingStart("salem", 10))
// 	fmt.Printf("%q\n", PaddingStart("salem𐙚", 10))
// 	fmt.Printf("%q\n", PaddingStart("salem", 10))
// }
