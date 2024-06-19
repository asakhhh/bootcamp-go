package bootcamp

func ForEach(arr []int, fn func(*int)) {
	for i := range arr {
		fn(&arr[i])
	}
}
