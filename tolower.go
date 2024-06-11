package bootcamp

func ToLower(s string) string {
	var t string
	ch := 'A' - 'a'
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			t = t + string(v-ch)
		} else {
			t = t + string(v)
		}
	}

	return t
}

// func main() {
// 	fmt.Println(ToUpper("salem "))        // "SALEM "
// 	fmt.Println(ToUpper("Salem Student")) // "SALEM STUDENT"
// 	fmt.Println(ToUpper("S4LEm"))         // "S4LEM"
// }
