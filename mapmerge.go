package bootcamp

func MapMerge(m1, m2 map[string]int) map[string]int {
	res := make(map[string]int)

	for i, v := range m1 {
		res[i] = v
	}
	for i, v := range m2 {
		res[i] = v
	}

	return res
}
