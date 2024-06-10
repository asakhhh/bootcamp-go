package bootcamp

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func SliceUnion(slices ...[]int) []int {
	var a []int

	for _, b := range slices {
		for _, v := range b {
			if !contains(a, v) {
				a = append(a, v)
			}
		}
	}

	return a
}

// func main() {
// 	result := SliceUnion([]int{1, 2, 3, 20}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 20, 4, 5, 15, 0]
// }
