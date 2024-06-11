package bootcamp

func Itoa(n int) string {
	var t string

	if n < 0 {
		t += "-"
		n = -n
	}

	p, l := 1, 1
	for ; p*10 <= n; p *= 10 {
		l++
	}

	for i := 0; i < l; i++ {
		t += string(rune('0' + (n/p)%10))
		p /= 10
	}
	return t
}

// func main() {
// 	fmt.Println(Itoa(123))
// 	fmt.Println(Itoa(-456))
// 	fmt.Println(Itoa(0))
// }
