package bootcamp

// import "fmt"

func SliceTargetIndexes(arr []int, target int) []int {
	var a []int

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			a = append(a, i)
		}
	}

	return a
}

// func main() {
// 	arr := []int{2, 3, 2, 5, 6, 7, 8, 2, 10}
// 	fmt.Println(SliceTargetIndexes(arr, 2)) // [0, 2, 7]
// 	fmt.Println(SliceTargetIndexes(arr, 1)) // []
// }
