package bootcamp

// import "fmt"

func SliceMakeN(n int) []int {
	if n < 0 {
		return make([]int, 0)
	}

	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	return a
}

// func main() {
// 	fmt.Println(SliceMakeN(5))  // [0, 1, 2, 3, 4]
// 	fmt.Println(SliceMakeN(10)) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// }
