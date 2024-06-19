package bootcamp

func Every(arr []int, fn func(int) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}
