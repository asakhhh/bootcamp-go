package bootcamp

func Eraser(s string) string {
	var res []rune
	for _, c := range s {
		if c == '<' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, c)
		}
	}
	return string(res)
}

// func main() {
// 	fmt.Printf("%q\n", Eraser("s<alem"))
// 	fmt.Printf("%q\n", Eraser("<<<\n"))
// }
