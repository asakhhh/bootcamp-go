package bootcamp

func MapKeys(m map[string]int) []string {
	var res []string
	for i := range m {
		res = append(res, i)
	}
	return res
}
