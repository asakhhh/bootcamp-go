package bootcamp

func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	t := *a
	*a = *b
	*b = t
}
