package bootcamp

func Join(elements []string, sep string) string {
	var t string
	for i, v := range elements {
		t += v
		if i != len(elements)-1 {
			t += sep
		}
	}
	return t
}
