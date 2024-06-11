package bootcamp

func Repeat(s string, count int) string {
	var t string

	for i := 0; i < count; i++ {
		t += s
	}

	return t
}
