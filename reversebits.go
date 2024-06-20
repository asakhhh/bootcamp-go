package bootcamp

func ReverseBits(b byte) byte {
	n := int(b)
	a := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 8; i++ {
		a[i] = n % 2
		n /= 2
	}
	// for i := 7; i-(8-ln) >= 0; i-- {
	// 	a[i] = a[i-(8-ln)]
	// }
	ans := 0
	for i := 0; i < 8; i++ {
		ans = ans*2 + a[i]
	}
	return byte(ans)
}

// func main() {
// 	fmt.Printf("%08b %08b\n", 5, ReverseBits(5))
// 	fmt.Printf("%08b %08b\n", 255, ReverseBits(255))
// 	fmt.Printf("%08b %08b\n", 1, ReverseBits(1))
// }
