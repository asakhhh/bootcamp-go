package bootcamp

func Sort(arr []int, fn func(int, int) bool) {
	if len(arr) <= 1 {
		return
	}
	n := len(arr)
	Sort(arr[:n/2], fn)
	Sort(arr[n/2:], fn)

	res := make([]int, 0)
	pt1, pt2 := 0, n/2
	for pt1 < n/2 || pt2 < n {
		if pt1 == n/2 {
			res = append(res, arr[pt2])
			pt2++
		} else if pt2 == n {
			res = append(res, arr[pt1])
			pt1++
		} else {
			if fn(arr[pt1], arr[pt2]) {
				res = append(res, arr[pt1])
				pt1++
			} else {
				res = append(res, arr[pt2])
				pt2++
			}
		}
	}
	copy(arr, res)
}

// func main() {
// 	ascending := func(a, b int) bool {
// 		return a < b
// 	}

// 	descending := func(a, b int) bool {
// 		return a > b
// 	}

// 	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// 	Sort(arr, ascending)
// 	fmt.Println(arr) // [1 1 2 3 4 5 5 6 9]

// 	arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// 	Sort(arr, descending)
// 	fmt.Println(arr) // [9 6 5 5 4 3 2 1 1]
// }
