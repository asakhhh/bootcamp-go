package bootcamp

func Find(arr []int, fn func(int) bool) *int {
	for i := range arr {
		if fn(arr[i]) {
			return &arr[i]
		}
	}
	return nil
}
