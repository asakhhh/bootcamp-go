package bootcamp

func MapLen(m map[string]int) int {
	cnt := 0
	for range m {
		cnt++
	}
	return cnt
}
