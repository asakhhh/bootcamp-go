package bootcamp

func MapValues(m map[string]int) []int {
	var res []int
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
