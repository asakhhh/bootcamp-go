package bootcamp

func MapClone(m map[string]int) map[string]int {
	var res map[string]int
	for i, v := range m {
		res[i] = v
	}
	return res
}
