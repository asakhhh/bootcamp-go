package bootcamp

// func MergeSortedArrays(arr1, arr2 []int) []int {
// 	n, m := len(arr1), len(arr2)
// 	pt1, pt2 := 0, 0
// 	var res []int
// 	for pt1 < n || pt2 < m {
// 		if pt1 == n {
// 			res = append(res, arr2[pt2])
// 			pt2++
// 		} else if pt2 == m {
// 			res = append(res, arr1[pt1])
// 			pt1++
// 		} else if arr1[pt1] < arr2[pt2] {
// 			res = append(res, arr1[pt1])
// 			pt1++
// 		} else {
// 			res = append(res, arr2[pt2])
// 			pt2++
// 		}
// 	}
// 	return res
// }

// func MergeSort(arr []int) {
// 	if len(arr) <= 1 {
// 		return
// 	}
// 	n := len(arr)
// 	MergeSort(arr[:n/2])
// 	MergeSort(arr[n/2:])
// 	res := MergeSortedArrays(arr[:n/2], arr[n/2:])
// 	copy(arr, res)
// }

// func MissingNumber(arr []int) int {
// 	mex := 1
// 	res := make([]int, len(arr))
// 	copy(res, arr)
// 	MergeSort(res)
// 	for _, v := range res {
// 		if v == mex {
// 			mex++
// 		}
// 	}
// 	return mex
// }

// func main() {
// 	fmt.Println(MissingNumber([]int{1, 2, 3, 4, 5}))
// }
