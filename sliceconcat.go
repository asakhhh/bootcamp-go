package bootcamp

func SliceConcat(slices ...[]int) []int {
	var a []int

	for _, i := range slices {
		a = append(a, i...)
	}

	return a
}

// func main() {
// 	result := SliceConcat([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 4, 5, 1, 2, 3, 4, 15, 0, 2]
// }
