package bootcamp

func SliceDelete(arr *[]int, low, high int) bool {
	if arr == nil || low < 0 || high > len(*arr) || low >= high {
		return false
	}

	if high == len(*arr) {
		*arr = (*arr)[:low]
	} else {
		t := make([]int, len(*arr)-high)
		copy(t, (*arr)[high:])
		*arr = append((*arr)[:low], t...)
	}

	return true
}

// func main() {
// 	arr := []int{0, 1, 2, 3, 4, 5}

// 	fmt.Println(arr)                     // [0, 1, 2, 3, 4, 5]
// 	fmt.Println(SliceDelete(&arr, 1, 3)) // true
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 3, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 5, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, -10, 5)) // false
// 	fmt.Println(arr)                       // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 0, 4)) // true
// 	fmt.Println(arr)                     // []
// }
