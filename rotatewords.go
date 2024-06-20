package bootcamp

// import "fmt"

// func isspace(s string) bool {
// 	return len(s) > 0 && s[0] == ' '
// }

// func Split(s string) []string {
// 	var res []string

// 	for _, c := range s {
// 		if c == ' ' {
// 			if len(res) > 0 && isspace(res[len(res)-1]) {
// 				res[len(res)-1] += " "
// 			} else {
// 				if len(res) == 0 {
// 					res = append(res, "")
// 				}
// 				res = append(res, " ")
// 			}
// 		} else {
// 			if len(res) == 0 || isspace(res[len(res)-1]) {
// 				res = append(res, string(c))
// 			} else {
// 				res[len(res)-1] += string(c)
// 			}
// 		}
// 	}
// 	if len(res) > 0 && isspace(res[len(res)-1]) {
// 		res = append(res, "")
// 	}

// 	return res
// }

// func RotateWords(s string) string {
// 	get := Split(s)
// 	var t []string
// 	var odd int
// 	if len(get) <= 1 {
// 		return ""
// 	} else if !isspace(get[0]) {
// 		t = make([]string, 0, (len(get))+1/2)
// 		for i := 0; i < len(get); i += 2 {
// 			t = append(t, get[i])
// 		}
// 		odd = 0
// 	} else {
// 		t = make([]string, 0, len(get)/2)
// 		for i := 1; i < len(get); i += 2 {
// 			t = append(t, get[i])
// 		}
// 		odd = 1
// 	}

// 	tt := []string{t[len(t)-1]}
// 	t = t[:len(t)-1]
// 	tt = append(tt, t...)
// 	fmt.Println(tt)

// 	for i := odd; i < len(get); i += 2 {
// 		get[i] = tt[i/2]
// 	}

// 	var ans string
// 	for _, v := range get {
// 		ans += v
// 	}
// 	return ans
// }

// // func main() {
// // 	fmt.Printf("%q\n", RotateWords("salem alem"))
// // 	fmt.Printf("%q\n", RotateWords("salem!    alem"))
// // 	fmt.Printf("%q\n", RotateWords("one two three four"))
// // 	fmt.Printf("%q\n", RotateWords("a b c d e f"))
// // }
