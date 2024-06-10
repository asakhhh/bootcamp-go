package bootcamp

func SlicePop(arr *[]int) int {
	if arr == nil || len(*arr) == 0 {
		return 0
	}

	t := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]

	return t
}

// func main() {
// 	arr := []int{5, 10, 20}
// 	size := len(arr)

// 	for i := 0; i < size; i++ {
// 		deleted := SlicePop(&arr)
// 		fmt.Println(deleted)
// 	}
// 	// Output:
// 	// 20
// 	// 10
// 	// 5

// 	deleted := SlicePop(&arr)
// 	fmt.Println(deleted) // 0
// }
