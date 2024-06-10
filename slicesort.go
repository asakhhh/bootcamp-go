package bootcamp

func SliceSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				t := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = t
			}
		}
	}
}

// func main() {
// 	arr := []int{10, 90, 20, 5, 12, 3, 0}
// 	SliceSort(arr)
// 	fmt.Println(arr) // [0, 3, 5, 10, 12, 20, 90]
// }
