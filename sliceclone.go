package bootcamp

func SliceClone(src, dst *[]int) {
	if src == nil {
		dst = nil
		return
	}

	*dst = make([]int, cap(*src))

	for i := 0; i < len(*src); i++ {
		(*dst)[i] = (*src)[i]
	}
}

// func main() {
// 	arr := make([]int, 10)
// 	for i, v := range []int{10, 20, 13, 5, 12, 31} {
// 		arr[i] = v
// 	}

// 	clone := []int{}

// 	fmt.Println("arr:", arr, len(arr), cap(arr))       // arr: [10, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(arr)) // clone: [] 0 0

// 	SliceClone(&arr, &clone)

// 	fmt.Println("arr:", arr, len(arr), cap(arr))       // arr: [10, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(arr)) // clone: [10, 20, 13, 5, 12, 31] 6 10
// }
