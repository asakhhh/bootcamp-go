package bootcamp

func MergeSortedArrays(arr1, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	pt1, pt2 := 0, 0
	var res []int
	for pt1 < n || pt2 < m {
		if pt1 == n {
			res = append(res, arr2[pt2])
			pt2++
		} else if pt2 == m {
			res = append(res, arr1[pt1])
			pt1++
		} else if arr1[pt1] < arr2[pt2] {
			res = append(res, arr1[pt1])
			pt1++
		} else {
			res = append(res, arr2[pt2])
			pt1++
		}
	}
	return res
}
