package bootcamp

func MapDeleteIf(m map[string]int, fn func(int) bool) {
	for k, v := range m {
		if fn(v) {
			delete(m, k)
		}
	}
}
