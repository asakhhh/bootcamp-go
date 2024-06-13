package bootcamp

func Contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func UniqComb2(s string) []string {
	if len(s) < 2 {
		return make([]string, 0)
	}
	for i, v := range s {
		if Contains(s[:i], v) {
			return make([]string, 0)
		}
	}

	var res []string
	for i, v := range s {
		for i1, v1 := range s {
			if i1 != i {
				res = append(res, string(v)+string(v1))
			}
		}
	}

	return res
}
