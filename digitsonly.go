package bootcamp

func DigitsOnly(s string) string {
	var t string

	for _, v := range s {
		if v >= '0' && v <= '9' {
			t = t + string(v)
		}
	}

	return t
}

// func main() {
// 	fmt.Println(DigitsOnly("abc123"))
// 	fmt.Println(DigitsOnly("Sa1em student! 23"))
// 	fmt.Println(DigitsOnly("no digits here!"))
// }
