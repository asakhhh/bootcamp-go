package bootcamp

func SliceBatch(slice []int, size int) [][]int {
	var a [][]int

	for i := 0; i+size < len(slice); i += size {
		t := make([]int, size)
		copy(t, slice[i:i+size])
		a = append(a, t)
	}
	if len(slice)%size != 0 {
		t := make([]int, len(slice)%size)
		copy(t, slice[len(slice)/size*size:])
		a = append(a, t)
	}

	return a
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	batch := SliceBatch(arr, 2)
// 	for _, v := range batch {
// 		fmt.Println(v)
// 	}
// 	// Output:
// 	// [1, 2]
// 	// [3, 4]
// 	// [5, 6]
// 	// [7]
// }
