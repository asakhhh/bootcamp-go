package bootcamp

func MapLen(m map[string]int) int {
	cnt := 0
	for range m {
		cnt++
	}
	return cnt
}

func MapContains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

func MapEqual(m1, m2 map[string]int) bool {
	if MapLen(m1) != MapLen(m2) {
		return false
	}

	for i, v := range m1 {
		if !MapContains(m2, i) || m2[i] != v {
			return false
		}
	}

	return true
}
