package bootcamp

func MapClone(m map[string]int) map[string]int {
	res := make(map[string]int)
	for i, v := range m {
		res[i] = v
	}
	return res
}
