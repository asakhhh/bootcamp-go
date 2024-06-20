package bootcamp

func ToggleBits(n byte) byte {
	num := 0
	p := 1

	for i := 0; i < 8; i++ {
		num += (1 - int(n)%2) * p
		p *= 2
		n /= 2
	}

	return byte(num)
}

// func main() {
// 	var b byte = 255
// 	fmt.Printf("%08b\n", b) // 11111111
// 	b = ToggleBits(b)
// 	fmt.Printf("%08b\n", b) // 00000000
// }
