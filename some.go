package bootcamp

func Some(arr []int, fn func(int) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}
