package bootcamp

// import "fmt"

func ArraySetOne(arr *[20]int, idx int, val int) bool {
	if idx >= 20 || arr == nil {
		return false
	}
	(*arr)[idx] = val
	return true
}

// func main() {
// 	arr := [20]int{1, 2, 3}
// 	ok := ArraySetOne(&arr, 1, 5)

// 	fmt.Println("ok:", ok)   // ok: true
// 	fmt.Println("arr:", arr) // arr: [1, 5, 3, ...]
// }
