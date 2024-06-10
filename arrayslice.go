package bootcamp

// import "fmt"

func ArraySlice(arr [20]int, low int, high int) []int {
	var a []int

	if high <= 0 {
		return a
	}

	for i := low; i < high; i++ {
		a = append(a, arr[i])
	}
	return a
}

// func main() {
// 	arr := [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(ArraySlice(arr, 3, 5)) // [3, 4]
// 	fmt.Println(ArraySlice(arr, 5, 5)) // []
// 	fmt.Println(ArraySlice(arr, 5, 1)) // []
// }
