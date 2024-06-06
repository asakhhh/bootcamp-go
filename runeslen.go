package bootcamp

// import "fmt"

func RunesLen(arr [20]rune) int {
	c := 0

	for c < 20 && arr[c] != 0 {
		c++
	}

	return c
}

// func main() {
// 	arr := [20]rune{'a', 'b', 'c', '\000', 'e'}

// 	fmt.Printf("%d", RunesLen(arr))
// }
