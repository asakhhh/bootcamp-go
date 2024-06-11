package bootcamp

func Concat(s ...string) string {
	var t string
	for _, v := range s {
		t += v
	}
	return t
}
