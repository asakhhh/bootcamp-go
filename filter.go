package bootcamp

func Filter(arr []int, fn func(int) bool) []int {
	var res []int

	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
