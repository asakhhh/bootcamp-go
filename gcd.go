package bootcamp

func GCD(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	if a > b {
		return GCD(b, a%b)
	}
	return GCD(a, b%a)
}

// func main() {
// 	fmt.Println(GCD(15, 10))
// 	fmt.Println(GCD(2, 6))
// }
