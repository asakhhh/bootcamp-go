package bootcamp

// import "fmt"

func ArrayEqual(arr1, arr2 *[20]int) bool {
	if arr1 == nil || arr2 == nil {
		return arr1 == nil && arr2 == nil
	}

	for i := 0; i < 20; i++ {
		if (*arr1)[i] != (*arr2)[i] {
			return false
		}
	}

	return true
}

// func main() {
// 	arr1 := [20]int{77, 69, 76, 65}
// 	arr2 := [20]int{77, 69, 76, 65}
// 	fmt.Println(ArrayEqual(&arr1, &arr2))            // true
// 	fmt.Println(ArrayEqual(&arr1, &[20]int{77, 78})) // false
// }
